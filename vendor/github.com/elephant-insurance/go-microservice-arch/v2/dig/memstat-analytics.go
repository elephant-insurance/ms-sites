package dig

import (
	"time"

	enum "github.com/elephant-insurance/enumerations/v2"
	"github.com/elephant-insurance/go-microservice-arch/v2/uf"
)

// Analytics are functions that examine diagnostic stats to determine whether to display alerts for bad conditions

// MemStatAnalytic is a function that looks at memory stats and returns alerts for any bad conditions it detects
type MemStatAnalytic func([]*memStat, ReadOnlyMemPeakJumpBuffer, SystemInfo) []*uf.Event

// memStatAnalytics is the singleton container for our memory analytics
var memStatAnalytics []MemStatAnalytic

// AddMemStatAnalytic is the only way to add a custom analytic from outside the package
// The count argument specifies how many of the most recent memStats the function expects to receive as its first argument
func AddMemStatAnalytic(count int, fn MemStatAnalytic) {
	if fn == nil {
		return
	}

	cfn := func(ms []*memStat, pjb ReadOnlyMemPeakJumpBuffer, si SystemInfo) []*uf.Event {
		memslice := memStats.Latest(count)
		sysi := getCurrentSystemInfo(false, nil)
		return fn(memslice, peakJumpBuffer, sysi)
	}

	memStatAnalytics = append(memStatAnalytics, cfn)
}

func runMemStatAnalytics() []*uf.Event {
	var rtn []*uf.Event
	if len(memStatAnalytics) < 1 {
		return rtn
	}

	rtn = []*uf.Event{}

	sysi := getCurrentSystemInfo(false, nil)
	for _, v := range memStatAnalytics {
		res := v(memStats.Earliest(0), peakJumpBuffer, sysi)
		rtn = append(rtn, res...)
	}

	if len(rtn) > 0 {
		return rtn
	}

	return nil
}

// MemStatAnalyticNewPeaksLast24Hours emits alerts if it detects new peaks in memory usage over the past 24 hours
var MemStatAnalyticNewPeaksLast24Hours = func(msa []*memStat, pjb ReadOnlyMemPeakJumpBuffer, sysi SystemInfo) []*uf.Event {
	// debug("MemStatAnalyticNewPeaksLast24Hours")

	var rtn []*uf.Event

	// Don't return anything until we've established a baseline, otherwise we might get a lot of meaningless warnings while the system warms up
	if memStats.GetBaseline() == nil {
		return rtn
	}

	aDayAgo := time.Now().Add(time.Hour * -24)
	sinceTime := aDayAgo
	bl := memStats.GetBaseline()
	if !bl.Timestamp.Before(aDayAgo) {
		// we haven't been running for a full day since we took our baseline, so only look at records saved since then
		sinceTime = bl.Timestamp.Add(time.Millisecond * -1)
	}

	// get the list of all new peaks recorded in the last day, or since the baseline, whichever is later
	peaks := peakJumpBuffer.Since(sinceTime, nil)

	// count each type of peak we find
	// Sys peaks are more urgent than others, and any stat that peaks more than once is a bigger concern
	allocPeaks, heapObjPeaks, sysPeaks := 0, 0, 0
	pkCnt := len(peaks)
	for i := 0; i < pkCnt; i++ {
		thisPeak := peaks[i]
		if thisPeak == nil {
			continue
		}

		switch thisPeak.StatPeaked {
		case MemStatFieldAlloc:
			allocPeaks++
		case MemStatFieldHeapObjects:
			heapObjPeaks++
		case MemStatFieldSys:
			sysPeaks++
		default:
			break
		}
	}

	// add alerts according to the type and number of new peaks found
	rtn = []*uf.Event{}
	switch allocPeaks {
	case 0:
		// nothing to do
		break
	case 1:
		// One new alloc peak in the last day is a small problem
		rtn = append(rtn, &uf.Event{
			Description: alertMsgMemPeakAllocOneLastDay,
			Severity:    &enum.EventSeverity.Cautionary.ID,
		})
	default:
		// Multiple alloc peak in the last day is a problem
		rtn = append(rtn, &uf.Event{
			Description: alertMsgMemPeakAllocMultiLastDay,
			Severity:    &enum.EventSeverity.Cautionary.ID,
		})
	}

	switch heapObjPeaks {
	case 0:
		// nothing to do
		break
	case 1:
		// One new heap objects peak in the last day is a small problem
		rtn = append(rtn, &uf.Event{
			Description: alertMsgMemPeakHeapObjOneLastDay,
			Severity:    &enum.EventSeverity.Cautionary.ID,
		})
	default:
		// Multiple new heap objects peaks in the last day is a problem
		rtn = append(rtn, &uf.Event{
			Description: alertMsgMemPeakHeapObjMultiLastDay,
			Severity:    &enum.EventSeverity.Cautionary.ID,
		})
	}

	switch sysPeaks {
	case 0:
		// nothing to do
		break
	case 1:
		// One new sys peak in the last day is a problem
		rtn = append(rtn, &uf.Event{
			Description: alertMsgMemPeakSysOneLastDay,
			Severity:    &enum.EventSeverity.Cautionary.ID,
		})
	default:
		// Multiple new peaks in the last day is a big problem
		rtn = append(rtn, &uf.Event{
			Description: alertMsgMemPeakSysMultiLastDay,
			Severity:    &enum.EventSeverity.Cautionary.ID,
		})
	}

	if len(rtn) > 0 {
		return rtn
	}

	return nil
}

// MemStatAnalyticGrowthSinceBaseline emits alerts if it finds a lot of RAM growth since the baseline was captured
var MemStatAnalyticGrowthSinceBaseline = func([]*memStat, ReadOnlyMemPeakJumpBuffer, SystemInfo) []*uf.Event {
	// debug("MemStatAnalyticGrowthSinceBaseline")

	var rtn []*uf.Event
	if memStats.GetBaseline() == nil {
		return rtn
	}

	rtn = []*uf.Event{}

	currentMem := getMemStats()
	bl := memStats.GetBaseline()

	// Get all the numbers we need to compare
	bAlloc, bHashObjects, bSys := bl.Alloc, bl.HeapObjects, bl.Sys

	// these are the 50% and 100% growth thresholds
	bAlloc50, bAlloc100 := bAlloc+bAlloc/2, bAlloc+bAlloc
	bHashObjects50, bHashObjects100 := bHashObjects+bHashObjects/2, bHashObjects+bHashObjects
	bSys50, bSys100 := bSys+bSys/2, bSys+bSys

	if currentMem.Alloc > bAlloc100 {
		rtn = append(rtn, &uf.Event{
			Description: alertMsgMemGrowthAlloc50PercentOverBaseline,
			Severity:    &enum.EventSeverity.Cautionary.ID,
		})
	} else if currentMem.Alloc > bAlloc50 {
		rtn = append(rtn, &uf.Event{
			Description: alertMsgMemGrowthAlloc100PercentOverBaseline,
			Severity:    &enum.EventSeverity.Cautionary.ID,
		})
	}

	if currentMem.HeapObjects > bHashObjects100 {
		rtn = append(rtn, &uf.Event{
			Description: alertMsgMemGrowthHashObj50PercentOverBaseline,
			Severity:    &enum.EventSeverity.Cautionary.ID,
		})
	} else if currentMem.HeapObjects > bHashObjects50 {
		rtn = append(rtn, &uf.Event{
			Description: alertMsgMemGrowthHashObj100PercentOverBaseline,
			Severity:    &enum.EventSeverity.Cautionary.ID,
		})
	}

	if currentMem.Sys > bSys100 {
		rtn = append(rtn, &uf.Event{
			Description: alertMsgMemGrowthSys100PercentOverBaseline,
			Severity:    &enum.EventSeverity.Cautionary.ID,
		})
	} else if currentMem.Sys > bSys50 {
		rtn = append(rtn, &uf.Event{
			Description: alertMsgMemGrowthSys50PercentOverBaseline,
			Severity:    &enum.EventSeverity.Cautionary.ID,
		})
	} else if currentMem.Sys > bSys {
		rtn = append(rtn, &uf.Event{
			Description: alertMsgMemGrowthSysAnyOverBaseline,
			Severity:    &enum.EventSeverity.Cautionary.ID,
		})
	}

	if len(rtn) > 0 {
		return rtn
	}

	return nil
}

// MemStatAnalyticMonotonicGrowthLast24Hours emits alerts if it finds positive RAM growth but no RAM reduction in the past 24 hours
// Note that Sys is ALWAYS monotonic (in other words, it never shrinks) so there is no point alerting on it
var MemStatAnalyticMonotonicGrowthLast24Hours = func([]*memStat, ReadOnlyMemPeakJumpBuffer, SystemInfo) []*uf.Event {
	// debug("MemStatAnalyticMonotonicGrowthLast24Hours")

	var rtn []*uf.Event
	if memStats.GetBaseline() == nil {
		return rtn
	}

	// how soon after startup are we interested in looking at this?
	const minLengthForMonotonicGrowth = 120

	rtn = []*uf.Event{}

	aDayAgo := time.Now().Add(time.Hour * -24)
	sinceTime := aDayAgo
	bl := memStats.GetBaseline()
	if !bl.Timestamp.Before(aDayAgo) {
		// we haven't been running for a full day since we took our baseline, so only look at records saved since then
		sinceTime = bl.Timestamp.Add(time.Millisecond * -1)
	}

	// get the list of all new peaks recorded in the last day, or since the baseline, whichever is later
	stats := memStats.Since(sinceTime)
	if len(stats) < minLengthForMonotonicGrowth {
		// either we haven't gotten to the baseline, or we did very recently
		// no point worrying about growth over such a short period
		return nil
	}

	current := getMemStats()

	// Have we grown at all in the past 24 hours?
	earliest := stats[len(stats)-1]
	if !(current.Alloc > earliest.Alloc || current.HeapObjects > earliest.HeapObjects) {
		// no net growth means no worries
		return nil
	}

	maxAlloc, maxHeap := current.Alloc, current.HeapObjects
	allEliminated, allocEliminated, heapEliminated := false, false, false
	for i := 0; i < len(stats) && !allEliminated; i++ {
		thisMS := stats[i]
		if thisMS == nil {
			continue
		}

		// going backwards, any increase in the max we've seen represents a RAM reduction
		// so we iterate backwards until we find an increase or we get to the end of the 24 hour array
		if thisMS.Alloc > maxAlloc {
			allocEliminated = true
		} else {
			maxAlloc = thisMS.Alloc
		}

		if thisMS.HeapObjects > maxHeap {
			heapEliminated = true
		} else {
			maxHeap = thisMS.HeapObjects
		}

		allEliminated = allocEliminated && heapEliminated
	}

	if !allocEliminated {
		rtn = append(rtn, &uf.Event{
			Description: alertMsgMemGrowthAllocMonotonicLastDay,
			Severity:    &enum.EventSeverity.Cautionary.ID,
		})
	}

	if !heapEliminated {
		rtn = append(rtn, &uf.Event{
			Description: alertMsgMemGrowthHashObjMonotonicLastDay,
			Severity:    &enum.EventSeverity.Cautionary.ID,
		})
	}

	if len(rtn) > 0 {
		return rtn
	}

	return nil
}

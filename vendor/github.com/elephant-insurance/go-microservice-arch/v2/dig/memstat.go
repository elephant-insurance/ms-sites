package dig

import (
	"fmt"
	"runtime"
	"time"
)

// memStat is just a much smaller version of GoLang runtime.MemStats, edited for easier reading.
type memStat struct {
	Alloc          uint64    `json:"BytesAllocated"`
	HeapObjects    uint64    `json:"HeapObjectCount"`
	SequenceNumber uint64    `json:"SequenceNumber,omitempty"`
	Sys            uint64    `json:"BytesReserved"`
	Timestamp      time.Time `json:"Timestamp,omitempty"`
}

// memStatCSVHeader is returned as the first row when memstats are requested in CSV format
const memStatCSVHeader = "SequenceNumber,BytesAllocated,BytesReserved,HeapObjectCount,Timestamp"

func abbreviate(src *runtime.MemStats) *memStat {
	if src == nil {
		return nil
	}

	rtn := memStat{
		Alloc:       src.Alloc,
		Sys:         src.Sys,
		HeapObjects: src.HeapObjects,
		Timestamp:   time.Now(),
	}

	return &rtn
}

// getMemStats creates a snapshot of our current memory usage
func getMemStats() *memStat {
	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)
	return abbreviate(&stats)
}

// memStats is the global singleton container for all memStats
var memStats memStatBuffer

// ToCSV renders the memstat record as a CSV row
func (m *memStat) ToCSV() string {
	if m == nil {
		return ""
	}
	return fmt.Sprintf("%v,%v,%v,%v,%v", m.SequenceNumber, m.Alloc, m.Sys, m.HeapObjects, m.Timestamp.Format(time.RFC3339Nano))
}

// Copy makes a new copy of the memStat
func (m *memStat) Copy() *memStat {
	if m == nil {
		return nil
	}

	return &memStat{
		Alloc:          m.Alloc,
		HeapObjects:    m.HeapObjects,
		SequenceNumber: m.SequenceNumber,
		Sys:            m.Sys,
		Timestamp:      m.Timestamp,
	}
}

// RecordMemStats fires a goroutine to take memory usage snapshots at intervals
func RecordMemStats() {
	interval := MemStatIntervalSeconds
	go func() {
		ticker := time.Tick(time.Duration(interval) * time.Second)
		for range ticker {
			AddMemStat()
		}
	}()
}

// makeMemStat creates a memstat from params, principally for testing
func makeMemStat(alloc, hobjs, sys uint64, timestamp *time.Time) *memStat {
	ts := timestamp
	if ts == nil {
		rightNow := time.Now()
		ts = &rightNow
	}

	return &memStat{
		Alloc:       alloc,
		HeapObjects: hobjs,
		Sys:         sys,
		Timestamp:   *ts,
	}
}

func newDummyMemStat() interface{} {
	return &memStat{
		Timestamp: time.Now(),
	}
}

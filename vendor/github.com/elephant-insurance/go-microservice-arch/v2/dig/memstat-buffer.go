package dig

import (
	"context"
	"sync"
	"time"

	"github.com/elephant-insurance/enumerations/v2"
	"github.com/elephant-insurance/go-microservice-arch/v2/clicker"
	"github.com/elephant-insurance/go-microservice-arch/v2/log"
	"github.com/elephant-insurance/go-microservice-arch/v2/rbuf"
	"github.com/elephant-insurance/go-microservice-arch/v2/uf"
)

type memStatBuffer interface {
	// Add places a new item at the end of the buffer
	Add(item *memStat)
	// Available is the actual number of entries currently in the buffer.
	Available() int
	// Capacity is the total number of items the buffer can hold. It is set at construction and does not change.
	Capacity() int
	// Earliest returns the earliest count items non-destructively, earliest first, without removing them
	Earliest(count int) []*memStat
	// First returns the single oldest entry available without removing it
	First() *memStat
	// Last returns the single latest entry available without removing it
	Last() *memStat
	// Latest returns the latest count items non-destructively, latest first, without removing them
	Latest(count int) []*memStat
	// Next removes and returns the single oldest entry, blocking until an entry is available, if necessary
	Next() *memStat
	// PullFromEarliest removes the count earliest entries from the buffer and returns them, earliest first
	PullFromEarliest(count int) []*memStat
	// PullFromLatest removes the count latest entries from the buffer and returns them, latest first
	PullFromLatest(count int) []*memStat
	// SelectCountFromEarliest non-destructively returns the count earliest entries from the buffer that satisfy the ItemSelector, earliest first
	SelectCountFromEarliest(selector rbuf.ItemSelector, count int) []*memStat
	// SelectCountFromLatest non-destructively returns the count latest entries from the buffer that satisfy the ItemSelector, latest first
	SelectCountFromLatest(selector rbuf.ItemSelector, count int) []*memStat
	// SelectRangeFromEarliest returns the continuous range of earliest entries from the buffer that satisfy the ItemSelector, earliest first
	// If the earliest element does not satisfy the ItemSelector, returns nil
	SelectRangeFromEarliest(selector rbuf.ItemSelector) []*memStat
	// SelectRangeFromLatest returns the continuous range of latest entries from the buffer that satisfy the ItemSelector, latest first
	// If the latest element does not satisfy the ItemSelector, returns nil
	SelectRangeFromLatest(selector rbuf.ItemSelector) []*memStat
	// TotalAdded returns the total number of items added to the buffer since it was created
	TotalAdded() int

	// After non-destructively returns the list of all memstats with sequence numbers greater than index
	After(index uint64) []*memStat

	// Since returns the latesst items with timestamp after sinceTime, earliest first
	Since(sinceTime time.Time) []*memStat

	// GetBaseline returns the baseline memstat recorded for this app, if any
	GetBaseline() *memStat
}

// msBuffer is an implementation of a typed ring-buffer for memStats
type msBuffer struct {
	baseline           *memStat
	myBuffer           rbuf.RingBuffer
	lastSequenceNumber *clicker.Clicker
	sync.Mutex
}

// NewMemStatBuffer returns a thread-safe (Concurrent) buffer for memStats.
func NewMemStatBuffer(capacity int) (memStatBuffer, error) {
	myBuf, err := rbuf.NewConcurrentRingBuffer(capacity, newDummyMemStat)
	if err != nil {
		return nil, err
	}

	return &msBuffer{
		baseline:           nil,
		myBuffer:           myBuf,
		lastSequenceNumber: &clicker.Clicker{},
	}, nil
}

// AddMemStat takes a reading fo the current memory usage and adds it to the global table of memStats
func AddMemStat() {
	stats := getMemStats()
	memStats.Add(stats)
	msg := `RAM usage statistics`
	evt := uf.EventFactory.New(&enumerations.Event.ApplicationDiagnostic.ID, nil, msg)
	log.ForFunc(context.Background()).WithEvent(evt).Debug(msg)
}

// RingBuffer Interface

// Add places a new item at the end of the buffer
func (msb *msBuffer) Add(item *memStat) {
	// overriding Add makes sure nothing but the type we want finds its way into the buffer
	// you may want to allow adding nil entries, depending on requirements
	if msb == nil || item == nil {
		return
	}
	msb.Lock()
	defer msb.Unlock()
	msb.lastSequenceNumber.Click(1)
	item.SequenceNumber = uint64(msb.lastSequenceNumber.Clicks)

	msb.myBuffer.Add(item)

	if item.Timestamp.Sub(startTime) > time.Minute*minutesToWaitForMemBaseline {
		// we've had enough time to initialize and get settled
		if msb.baseline == nil {
			// we haven't recorded our memStatBaseline yet, so do that now
			msb.baseline = item
		}

		peakJumpBuffer.CheckForPeak(item)
	}
}

// Available is the actual number of entries currently in the buffer.
func (msb *msBuffer) Available() int {
	if msb == nil {
		return 0
	}

	return msb.myBuffer.Available()
}

// Capacity is the total number of items the buffer can hold. It is set at construction and does not change.
func (msb *msBuffer) Capacity() int {
	if msb == nil {
		return 0
	}

	return msb.myBuffer.Capacity()
}

// Earliest returns the earliest count items non-destructively, earliest first, without removing them
func (msb *msBuffer) Earliest(count int) []*memStat {
	if msb == nil {
		return nil
	}

	msb.Lock()
	defer msb.Unlock()

	ir := msb.myBuffer.Earliest(count)

	return convertArray(ir, false)
}

// First returns the single oldest entry available without removing it
func (msb *msBuffer) First() *memStat {
	if msb == nil {
		return nil
	}

	w := msb.myBuffer.First()
	if w != nil {
		return w.(*memStat)
	}

	return nil
}

func (msb *msBuffer) GetBaseline() *memStat {
	return msb.baseline
}

// Last returns the single latest entry available without removing it
func (msb *msBuffer) Last() *memStat {
	if msb == nil {
		return nil
	}

	w := msb.myBuffer.Last()
	if w != nil {
		return w.(*memStat)
	}

	return nil
}

// Latest returns the latest count items non-destructively, latest first, without removing them
func (msb *msBuffer) Latest(count int) []*memStat {
	if msb == nil {
		return nil
	}

	msb.Lock()
	defer msb.Unlock()

	ir := msb.myBuffer.Latest(count)

	return convertArray(ir, false)
}

// Next removes and returns the single oldest entry, blocking until an entry is available, if necessary
func (msb *msBuffer) Next() *memStat {
	if msb == nil {
		return nil
	}

	w := msb.myBuffer.Next()
	if w != nil {
		return w.(*memStat)
	}

	return nil
}

// PullFromEarliest removes the count earliest entries from the buffer and returns them, earliest first
func (msb *msBuffer) PullFromEarliest(count int) []*memStat {
	if msb == nil {
		return nil
	}

	msb.Lock()
	defer msb.Unlock()

	ir := msb.myBuffer.PullFromEarliest(count)

	return convertArray(ir, false)
}

// PullFromLatest removes the count latest entries from the buffer and returns them, latest first
func (msb *msBuffer) PullFromLatest(count int) []*memStat {
	if msb == nil {
		return nil
	}

	msb.Lock()
	defer msb.Unlock()

	ir := msb.myBuffer.PullFromLatest(count)

	return convertArray(ir, false)
}

// SelectCountFromEarliest non-destructively returns the count earliest entries from the buffer that satisfy the ItemSelector, earliest first
func (msb *msBuffer) SelectCountFromEarliest(selector rbuf.ItemSelector, count int) []*memStat {
	if msb == nil {
		return nil
	}

	ir := msb.myBuffer.SelectCountFromEarliest(selector, count)

	return convertArray(ir, false)
}

// SelectCountFromLatest non-destructively returns the count latest entries from the buffer that satisfy the ItemSelector, latest first
func (msb *msBuffer) SelectCountFromLatest(selector rbuf.ItemSelector, count int) []*memStat {
	if msb == nil {
		return nil
	}

	ir := msb.myBuffer.SelectCountFromLatest(selector, count)

	return convertArray(ir, false)
}

// SelectRangeFromEarliest returns the continuous range of earliest entries from the buffer that satisfy the ItemSelector, earliest first
// If the earliest element does not satisfy the ItemSelector, returns nil
func (msb *msBuffer) SelectRangeFromEarliest(selector rbuf.ItemSelector) []*memStat {
	if msb == nil {
		return nil
	}

	ir := msb.myBuffer.SelectRangeFromEarliest(selector)

	return convertArray(ir, false)
}

// SelectRangeFromLatest returns the continuous range of latest entries from the buffer that satisfy the ItemSelector, latest first
// If the latest element does not satisfy the ItemSelector, returns nil
func (msb *msBuffer) SelectRangeFromLatest(selector rbuf.ItemSelector) []*memStat {
	if msb == nil {
		return nil
	}

	ir := msb.myBuffer.SelectRangeFromLatest(selector)

	return convertArray(ir, false)
}

// TotalAdded returns the total number of items added to the buffer since it was created
func (msb *msBuffer) TotalAdded() int {
	if msb == nil {
		return 0
	}

	return msb.myBuffer.TotalAdded()
}

// convertArray converts an array of interface{} to []*memStat
func convertArray(ary []interface{}, reverse bool) []*memStat {
	var rtn []*memStat
	if ary == nil {
		return rtn
	}

	count := len(ary)
	rtn = make([]*memStat, count)
	indexer := func(i int) int { return i }
	if reverse {
		indexer = func(i int) int { return count - (i + 1) }
	}

	for i := 0; i < count; i++ {
		var wp *memStat
		obj := ary[i]
		if obj != nil {
			wp = obj.(*memStat)
		}

		rtn[indexer(i)] = wp
	}

	return rtn
}

// After non-destructively returns the list of all memstats with sequence numbers greater than index
func (msb *msBuffer) After(index uint64) []*memStat {
	if msb == nil {
		return nil
	}

	sinceSelector := func(obj interface{}) bool {
		w, ok := obj.(*memStat)
		if !ok {
			return false
		}

		return index < w.SequenceNumber
	}

	ir := msb.myBuffer.SelectRangeFromLatest(sinceSelector)

	return convertArray(ir, true)
}

// Since non-destructively returns the list of all memstats with timestamps after sinceTime
func (msb *msBuffer) Since(sinceTime time.Time) []*memStat {
	if msb == nil {
		return nil
	}

	sinceSelector := func(obj interface{}) bool {
		w, ok := obj.(*memStat)
		if !ok {
			return false
		}

		return sinceTime.Before(w.Timestamp)
	}

	ir := msb.myBuffer.SelectRangeFromLatest(sinceSelector)

	return convertArray(ir, true)
}

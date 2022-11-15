package timing

import (
	"sync"
	"time"

	"github.com/elephant-insurance/go-microservice-arch/v2/clicker"
	"github.com/elephant-insurance/go-microservice-arch/v2/rbuf"
)

type TimingBuffer interface {
	// Add places a new item at the end of the buffer
	Add(item *Timing)
	// Available is the actual number of entries currently in the buffer.
	Available() int
	// Capacity is the total number of items the buffer can hold. It is set at construction and does not change.
	Capacity() int
	// Earliest returns the earliest count items non-destructively, earliest first, without removing them
	Earliest(count int) []*Timing
	// First returns the single oldest entry available without removing it
	First() *Timing
	// Last returns the single latest entry available without removing it
	Last() *Timing
	// Latest returns the latest count items non-destructively, latest first, without removing them
	Latest(count int) []*Timing
	// Next removes and returns the single oldest entry, blocking until an entry is available, if necessary
	Next() *Timing
	// PullFromEarliest removes the count earliest entries from the buffer and returns them, earliest first
	PullFromEarliest(count int) []*Timing
	// PullFromLatest removes the count latest entries from the buffer and returns them, latest first
	PullFromLatest(count int) []*Timing
	// SelectCountFromEarliest non-destructively returns the count earliest entries from the buffer that satisfy the ItemSelector, earliest first
	SelectCountFromEarliest(selector rbuf.ItemSelector, count int) []*Timing
	// SelectCountFromLatest non-destructively returns the count latest entries from the buffer that satisfy the ItemSelector, latest first
	SelectCountFromLatest(selector rbuf.ItemSelector, count int) []*Timing
	// SelectRangeFromEarliest returns the continuous range of earliest entries from the buffer that satisfy the ItemSelector, earliest first
	// If the earliest element does not satisfy the ItemSelector, returns nil
	SelectRangeFromEarliest(selector rbuf.ItemSelector) []*Timing
	// SelectRangeFromLatest returns the continuous range of latest entries from the buffer that satisfy the ItemSelector, latest first
	// If the latest element does not satisfy the ItemSelector, returns nil
	SelectRangeFromLatest(selector rbuf.ItemSelector) []*Timing
	// TotalAdded returns the total number of items added to the buffer since it was created
	TotalAdded() int

	// After non-destructively returns the list of all timings with sequence numbers greater than index
	After(index uint64) []*Timing

	// Since returns the latesst items with timestamp after sinceTime, earliest first
	Since(sinceTime time.Time) []*Timing
}

// ApplicationTimings => this tiny initialized value will be overwritten by Initialize
var ApplicationTimings, _ = NewTimingBuffer(10)

// NewTimingBuffer returns a thread-safe (Concurrent) buffer for timings.
func NewTimingBuffer(capacity int) (TimingBuffer, error) {
	myBuf, err := rbuf.NewConcurrentRingBuffer(capacity, newDummyTiming)
	if err != nil {
		return nil, err
	}

	return &tmBuffer{
		myBuffer:           myBuf,
		lastSequenceNumber: &clicker.Clicker{},
	}, nil
}

type tmBuffer struct {
	myBuffer           rbuf.RingBuffer
	lastSequenceNumber *clicker.Clicker
	sync.Mutex
}

// Add places a new item at the end of the buffer
func (tb *tmBuffer) Add(item *Timing) {
	// overriding Add makes sure nothing but the type we want finds its way into the buffer
	// you may want to allow adding nil entries, depending on requirements
	if tb == nil || item == nil {
		return
	}
	tb.Lock()
	defer tb.Unlock()
	tb.lastSequenceNumber.Click(1)
	item.SequenceNumber = uint64(tb.lastSequenceNumber.Clicks)

	tb.myBuffer.Add(item)
}

// Available is the actual number of entries currently in the buffer.
func (tb *tmBuffer) Available() int {
	if tb == nil {
		return 0
	}

	return tb.myBuffer.Available()
}

// Capacity is the total number of items the buffer can hold. It is set at construction and does not change.
func (tb *tmBuffer) Capacity() int {
	if tb == nil {
		return 0
	}

	return tb.myBuffer.Capacity()
}

// Earliest returns the earliest count items non-destructively, earliest first, without removing them
func (tb *tmBuffer) Earliest(count int) []*Timing {
	if tb == nil {
		return nil
	}

	tb.Lock()
	defer tb.Unlock()

	ir := tb.myBuffer.Earliest(count)

	return convertTimingArray(ir, false)
}

// First returns the single oldest entry available without removing it
func (tb *tmBuffer) First() *Timing {
	if tb == nil {
		return nil
	}

	w := tb.myBuffer.First()
	if w != nil {
		return w.(*Timing)
	}

	return nil
}

// Last returns the single latest entry available without removing it
func (tb *tmBuffer) Last() *Timing {
	if tb == nil {
		return nil
	}

	w := tb.myBuffer.Last()
	if w != nil {
		return w.(*Timing)
	}

	return nil
}

// Latest returns the latest count items non-destructively, latest first, without removing them
func (tb *tmBuffer) Latest(count int) []*Timing {
	if tb == nil {
		return nil
	}

	tb.Lock()
	defer tb.Unlock()

	ir := tb.myBuffer.Latest(count)

	return convertTimingArray(ir, false)
}

// Next removes and returns the single oldest entry, blocking until an entry is available, if necessary
func (tb *tmBuffer) Next() *Timing {
	if tb == nil {
		return nil
	}

	w := tb.myBuffer.Next()
	if w != nil {
		return w.(*Timing)
	}

	return nil
}

// PullFromEarliest removes the count earliest entries from the buffer and returns them, earliest first
func (tb *tmBuffer) PullFromEarliest(count int) []*Timing {
	if tb == nil {
		return nil
	}

	tb.Lock()
	defer tb.Unlock()

	ir := tb.myBuffer.PullFromEarliest(count)

	return convertTimingArray(ir, false)
}

// PullFromLatest removes the count latest entries from the buffer and returns them, latest first
func (tb *tmBuffer) PullFromLatest(count int) []*Timing {
	if tb == nil {
		return nil
	}

	tb.Lock()
	defer tb.Unlock()

	ir := tb.myBuffer.PullFromLatest(count)

	return convertTimingArray(ir, false)
}

// SelectCountFromEarliest non-destructively returns the count earliest entries from the buffer that satisfy the ItemSelector, earliest first
func (tb *tmBuffer) SelectCountFromEarliest(selector rbuf.ItemSelector, count int) []*Timing {
	if tb == nil {
		return nil
	}

	ir := tb.myBuffer.SelectCountFromEarliest(selector, count)

	return convertTimingArray(ir, false)
}

// SelectCountFromLatest non-destructively returns the count latest entries from the buffer that satisfy the ItemSelector, latest first
func (tb *tmBuffer) SelectCountFromLatest(selector rbuf.ItemSelector, count int) []*Timing {
	if tb == nil {
		return nil
	}

	ir := tb.myBuffer.SelectCountFromLatest(selector, count)

	return convertTimingArray(ir, false)
}

// SelectRangeFromEarliest returns the continuous range of earliest entries from the buffer that satisfy the ItemSelector, earliest first
// If the earliest element does not satisfy the ItemSelector, returns nil
func (tb *tmBuffer) SelectRangeFromEarliest(selector rbuf.ItemSelector) []*Timing {
	if tb == nil {
		return nil
	}

	ir := tb.myBuffer.SelectRangeFromEarliest(selector)

	return convertTimingArray(ir, false)
}

// SelectRangeFromLatest returns the continuous range of latest entries from the buffer that satisfy the ItemSelector, latest first
// If the latest element does not satisfy the ItemSelector, returns nil
func (tb *tmBuffer) SelectRangeFromLatest(selector rbuf.ItemSelector) []*Timing {
	if tb == nil {
		return nil
	}

	ir := tb.myBuffer.SelectRangeFromLatest(selector)

	return convertTimingArray(ir, false)
}

// TotalAdded returns the total number of items added to the buffer since it was created
func (tb *tmBuffer) TotalAdded() int {
	if tb == nil {
		return 0
	}

	return tb.myBuffer.TotalAdded()
}

// convertTimingArray converts an array of interface{} to []*Timing
func convertTimingArray(ary []interface{}, reverse bool) []*Timing {
	var rtn []*Timing
	if ary == nil {
		return rtn
	}

	count := len(ary)
	rtn = make([]*Timing, count)
	indexer := func(i int) int { return i }
	if reverse {
		indexer = func(i int) int { return count - (i + 1) }
	}

	for i := 0; i < count; i++ {
		var wp *Timing
		obj := ary[i]
		if obj != nil {
			wp = obj.(*Timing)
		}

		rtn[indexer(i)] = wp
	}

	return rtn
}

// After non-destructively returns the list of all memstats with sequence numbers greater than index
func (tb *tmBuffer) After(index uint64) []*Timing {
	if tb == nil {
		return nil
	}

	sinceSelector := func(obj interface{}) bool {
		w, ok := obj.(*Timing)
		if !ok {
			return false
		}

		return index < w.SequenceNumber
	}

	ir := tb.myBuffer.SelectRangeFromLatest(sinceSelector)

	return convertTimingArray(ir, true)
}

// Since non-destructively returns the list of all memstats with timestamps after sinceTime
func (tb *tmBuffer) Since(sinceTime time.Time) []*Timing {
	if tb == nil {
		return nil
	}

	sinceSelector := func(obj interface{}) bool {
		w, ok := obj.(*Timing)
		if !ok {
			return false
		}

		return sinceTime.Before(w.TimeInitiated)
	}

	ir := tb.myBuffer.SelectRangeFromLatest(sinceSelector)

	return convertTimingArray(ir, true)
}

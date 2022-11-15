package dig

import (
	"time"

	rb "github.com/elephant-insurance/go-microservice-arch/v2/rbuf"
)

// peakJumpBuffer is a package-level singleton collection of all memory maxes
var peakJumpBuffer MemPeakJumpBuffer

// MemPeakJumpBuffer is a type-specific version of the RingBuffer interface
type MemPeakJumpBuffer interface {
	// Add places a new item at the end of the buffer
	Add(item *MemPeakJump)
	// Available is the actual number of entries currently in the buffer.
	Available() int
	// Capacity is the total number of items the buffer can hold. It is set at construction and does not change.
	Capacity() int
	// Earliest returns the earliest count items non-destructively, earliest first, without removing them
	Earliest(count int) []*MemPeakJump
	// First returns the single oldest entry available without removing it
	First() *MemPeakJump
	// Last returns the single latest entry available without removing it
	Last() *MemPeakJump
	// Latest returns the latest count items non-destructively, latest first, without removing them
	Latest(count int) []*MemPeakJump
	// Next removes and returns the single oldest entry, blocking until an entry is available, if necessary
	Next() *MemPeakJump
	// PullFromEarliest removes the count earliest entries from the buffer and returns them, earliest first
	PullFromEarliest(count int) []*MemPeakJump
	// PullFromLatest removes the count latest entries from the buffer and returns them, latest first
	PullFromLatest(count int) []*MemPeakJump
	// TotalAdded returns the total number of items added to the buffer since it was created
	TotalAdded() int
	// CheckForPeak checks to see whether this memStat represents a new max for any memStat field
	// If this is a new max, it is added to the newPeak buffer and the latest maxes are updated
	CheckForPeak(m *memStat)
	// Since returns all peaks recorded since the time submitted, optionally filtered by Type
	Since(t time.Time, f *MemStatField) []*MemPeakJump
}

// ReadOnlyMemPeakJumpBuffer is a read-only subset of the MemPeakJumpBuffer interface
type ReadOnlyMemPeakJumpBuffer interface {
	// Available is the actual number of entries currently in the buffer.
	Available() int
	// Capacity is the total number of items the buffer can hold. It is set at construction and does not change.
	Capacity() int
	// Earliest returns the earliest count items non-destructively, earliest first, without removing them
	Earliest(count int) []*MemPeakJump
	// First returns the single oldest entry available without removing it
	First() *MemPeakJump
	// Last returns the single latest entry available without removing it
	Last() *MemPeakJump
	// Latest returns the latest count items non-destructively, latest first, without removing them
	Latest(count int) []*MemPeakJump
	// TotalAdded returns the total number of items added to the buffer since it was created
	TotalAdded() int
	// Since returns all peaks recorded since the time submitted, optionally filtered by Type
	Since(t time.Time, f *MemStatField) []*MemPeakJump
}

// nmBuffer is an example implementation of a typed ring-buffer
// The safest implementation is a private type (nmBuffer) implementing a public interface (MemPeakJumpBuffer)
type nmBuffer struct {
	// currentPeaks is a computed map of field id's to their latest max values
	// This just saves us having to recompute the maxes every time we want to check for a new one
	currentPeaks map[MemStatField]uint64
	myBuffer     rb.RingBuffer
}

func NewMemPeakJumpBuffer(capacity int) (MemPeakJumpBuffer, error) {
	myBuf, err := rb.NewRingBuffer(capacity, NewDummyMemPeakJump)
	if err != nil {
		return nil, err
	}

	return &nmBuffer{
		currentPeaks: newPeakMap(),
		myBuffer:     myBuf,
	}, nil
}

func NewConcurrentMemPeakJumpBuffer(capacity int) (MemPeakJumpBuffer, error) {
	myBuf, err := rb.NewConcurrentRingBuffer(capacity, NewDummyMemPeakJump)
	if err != nil {
		return nil, err
	}

	return &nmBuffer{
		currentPeaks: newPeakMap(),
		myBuffer:     myBuf,
	}, nil
}

func newPeakMap() map[MemStatField]uint64 {
	return map[MemStatField]uint64{
		MemStatFieldUnknown:     0,
		MemStatFieldAlloc:       0,
		MemStatFieldHeapObjects: 0,
		MemStatFieldSys:         0,
	}
}

// GetMemPeaks returns copies of all memory peaks in the buffer
func GetMemPeaks() []*MemPeakJump {
	return peakJumpBuffer.Latest(0)
}

// RingBuffer Interface
// Add places a new item at the end of the buffer
func (nmb *nmBuffer) Add(item *MemPeakJump) {
	// overriding Add makes sure nothing but the type we want finds its way into the buffer
	// you may want to allow adding nil entries, depending on requirements
	if nmb == nil || item == nil {
		return
	}

	nmb.myBuffer.Add(item)
}

// Available is the actual number of entries currently in the buffer.
func (nmb *nmBuffer) Available() int {
	if nmb == nil {
		return 0
	}

	return nmb.myBuffer.Available()
}

// Capacity is the total number of items the buffer can hold. It is set at construction and does not change.
func (nmb *nmBuffer) Capacity() int {
	if nmb == nil {
		return 0
	}

	return nmb.myBuffer.Capacity()
}

// Earliest returns the earliest count items non-destructively, earliest first, without removing them
func (nmb *nmBuffer) Earliest(count int) []*MemPeakJump {
	if nmb == nil {
		return nil
	}

	ir := nmb.myBuffer.Earliest(count)
	rtn := make([]*MemPeakJump, len(ir))

	for i := 0; i < len(ir); i++ {
		var wp *MemPeakJump
		if ir[i] != nil {
			wp = ir[i].(*MemPeakJump)
		}

		rtn[i] = wp
	}

	return rtn
}

// First returns the single oldest entry available without removing it
func (nmb *nmBuffer) First() *MemPeakJump {
	if nmb == nil {
		return nil
	}

	w := nmb.myBuffer.First()
	if w != nil {
		return w.(*MemPeakJump)
	}

	return nil
}

// Last returns the single latest entry available without removing it
func (nmb *nmBuffer) Last() *MemPeakJump {
	if nmb == nil {
		return nil
	}

	w := nmb.myBuffer.Last()
	if w != nil {
		return w.(*MemPeakJump)
	}

	return nil
}

// Latest returns the latest count items non-destructively, latest first, without removing them
func (nmb *nmBuffer) Latest(count int) []*MemPeakJump {
	if nmb == nil {
		return nil
	}

	ir := nmb.myBuffer.Latest(count)
	rtn := make([]*MemPeakJump, len(ir))

	for i := 0; i < len(ir); i++ {
		var wp *MemPeakJump
		if ir[i] != nil {
			wp = ir[i].(*MemPeakJump)
		}

		rtn[i] = wp
	}

	return rtn
}

// Next removes and returns the single oldest entry, blocking until an entry is available, if necessary
func (nmb *nmBuffer) Next() *MemPeakJump {
	if nmb == nil {
		return nil
	}

	w := nmb.myBuffer.Next()
	if w != nil {
		return w.(*MemPeakJump)
	}

	return nil
}

// PullFromEarliest removes the count earliest entries from the buffer and returns them, earliest first
func (nmb *nmBuffer) PullFromEarliest(count int) []*MemPeakJump {
	if nmb == nil {
		return nil
	}

	ir := nmb.myBuffer.PullFromEarliest(count)
	rtn := make([]*MemPeakJump, len(ir))

	for i := 0; i < len(ir); i++ {
		var wp *MemPeakJump
		if ir[i] != nil {
			wp = ir[i].(*MemPeakJump)
		}

		rtn[i] = wp
	}

	return rtn
}

// PullFromLatest removes the count latest entries from the buffer and returns them, latest first
func (nmb *nmBuffer) PullFromLatest(count int) []*MemPeakJump {
	if nmb == nil {
		return nil
	}

	ir := nmb.myBuffer.PullFromLatest(count)
	rtn := make([]*MemPeakJump, len(ir))

	for i := 0; i < len(ir); i++ {
		var wp *MemPeakJump
		if ir[i] != nil {
			wp = ir[i].(*MemPeakJump)
		}

		rtn[i] = wp
	}

	return rtn
}

// TotalAdded returns the total number of items added to the buffer since it was created
func (nmb *nmBuffer) TotalAdded() int {
	if nmb == nil {
		return 0
	}

	return nmb.myBuffer.TotalAdded()
}

// CheckForPeak checks to see whether this memStat represents a new max for any memStat field
// If this is a new max, it is added to the newPeak buffer and the latest maxes are updated
func (nmb *nmBuffer) CheckForPeak(m *memStat) {
	if nmb == nil || m == nil {
		return
	}

	if m.Alloc > nmb.currentPeaks[MemStatFieldAlloc] {
		nmb.currentPeaks[MemStatFieldAlloc] = m.Alloc
		nmb.Add(&MemPeakJump{
			Record:     m,
			StatPeaked: MemStatFieldAlloc,
		})
	}

	if m.HeapObjects > nmb.currentPeaks[MemStatFieldHeapObjects] {
		nmb.currentPeaks[MemStatFieldHeapObjects] = m.HeapObjects
		nmb.Add(&MemPeakJump{
			Record:     m,
			StatPeaked: MemStatFieldHeapObjects,
		})
	}

	if m.Sys > nmb.currentPeaks[MemStatFieldSys] {
		nmb.currentPeaks[MemStatFieldSys] = m.Sys
		nmb.Add(&MemPeakJump{
			Record:     m,
			StatPeaked: MemStatFieldSys,
		})
	}
}

// Since returns all peaks recorded since the time submitted, optionally filtered by Type
func (nmb *nmBuffer) Since(t time.Time, f *MemStatField) []*MemPeakJump {
	var rtn []*MemPeakJump

	latest := nmb.Latest(0)
	if len(latest) < 1 {
		return rtn
	}

	rtn = []*MemPeakJump{}

	found := false
	for i := 0; i < len(latest) && !found; i++ {
		thisMS := latest[i]
		if thisMS != nil {
			if thisMS.Record.Timestamp.Before(t) {
				// we've gone far enough back in the records to pass the "since" time, so we are done
				found = true
			}

			if f == nil || thisMS.StatPeaked == *f {
				rtn = append(rtn, thisMS)
			}
		}
	}

	if len(rtn) > 0 {
		return rtn
	}

	return nil
}

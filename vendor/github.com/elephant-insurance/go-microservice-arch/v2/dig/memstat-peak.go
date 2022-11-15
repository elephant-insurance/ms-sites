package dig

import "time"

// vars and types for keeping track of changes in memstats over time for on-page alerting

// MemPeakJump is a record of a new maximum memstat, i.e. a jump in the peak of a statistic
// New maximums are a useful tool for diagnosing memory leaks
type MemPeakJump struct {
	// Record is a pointer to the memory snapshot containing the new max
	Record *memStat
	// StatPeaked tracks which stat was actually maxed
	StatPeaked MemStatField
}

// Copy makes a new copy of the MemPeakJump
func (mpj *MemPeakJump) Copy() *MemPeakJump {
	if mpj == nil {
		return nil
	}

	ms := mpj.Record.Copy()

	return &MemPeakJump{Record: ms, StatPeaked: mpj.StatPeaked}
}

// NewDummyMemPeakJump creates a dummy mem peak jump for pre-allocating memory
func NewDummyMemPeakJump() interface{} {
	ms := &memStat{
		Alloc:          dummyPeakAlloc,
		HeapObjects:    dummyPeakHeapObjects,
		SequenceNumber: dummyPeakSequenceNumber,
		Sys:            dummyPeakSys,
		Timestamp:      time.Now(),
	}

	return &MemPeakJump{
		Record:     ms,
		StatPeaked: MemStatFieldUnknown,
	}
}

type MemStatField int

const (
	MemStatFieldUnknown MemStatField = iota
	MemStatFieldAlloc
	MemStatFieldHeapObjects
	MemStatFieldSys
)

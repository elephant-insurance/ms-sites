package rbuf

// RingBuffer is a general-purpose buffer structure for storing sequences of data objects
// It has a fixed size, set at startup, and cannot grow or shrink
// It will optionally pre-fill itself with dummy objects on construction to limit memory usage changes
// The buffer is circular in structure, so it can never be overrun
// Additional entries written to a full buffer overwrite the earliest entries
// It can be used as a FIFO buffer or a LIFO stack
// Entries can be removed as they are processed, or just allowed to roll off as they're overwritten
// Retrieved entities will require type assertion before they can be used
//  (e.g., var myEntity *myType; myEntity, ok := (buf.Last()).(*myType))

type RingBuffer interface {
	// Add places a new item at the end of the buffer
	Add(item interface{})
	// Available is the actual number of entries currently in the buffer.
	Available() int
	// Capacity is the total number of items the buffer can hold. It is set at construction and does not change.
	Capacity() int
	// Earliest returns the earliest count items non-destructively, earliest first, without removing them
	Earliest(count int) []interface{}
	// First returns the single oldest entry available without removing it
	First() interface{}
	// Last returns the single latest entry available without removing it
	Last() interface{}
	// Latest returns the latest count items non-destructively, latest first, without removing them
	Latest(count int) []interface{}
	// Next removes and returns the single oldest entry, blocking until an entry is available, if necessary
	Next() interface{}
	// PullFromEarliest removes the count earliest entries from the buffer and returns them, earliest first
	PullFromEarliest(count int) []interface{}
	// PullFromLatest removes the count latest entries from the buffer and returns them, latest first
	PullFromLatest(count int) []interface{}
	// SelectCountFromEarliest non-destructively returns the count earliest entries from the buffer that satisfy the ItemSelector, earliest first
	SelectCountFromEarliest(selector ItemSelector, count int) []interface{}
	// SelectCountFromLatest non-destructively returns the count latest entries from the buffer that satisfy the ItemSelector, latest first
	SelectCountFromLatest(selector ItemSelector, count int) []interface{}
	// SelectRangeFromEarliest returns the continuous range of earliest entries from the buffer that satisfy the ItemSelector, earliest first
	// If the earliest element does not satisfy the ItemSelector, returns nil
	SelectRangeFromEarliest(selector ItemSelector) []interface{}
	// SelectRangeFromLatest returns the continuous range of latest entries from the buffer that satisfy the ItemSelector, latest first
	// If the latest element does not satisfy the ItemSelector, returns nil
	SelectRangeFromLatest(selector ItemSelector) []interface{}
	// TotalAdded returns the total number of items added to the buffer since it was created
	TotalAdded() int
}

// DummyFactory is a simple function for generating dummy entries at startup for pre-filling the buffer
// Use a function that creates full-sized entries to prevent memory fluctuations as the buffer grows
// Use nil (or a func that returns nil) to skip the pre-filling step
type DummyFactory func() interface{}

// ItemSelector is a custom selector function that can be passed in for querys against the buffer
// The implementing func should be aware of the underlying type of the elements
// and have full access to its "real" properties
type ItemSelector func(interface{}) bool

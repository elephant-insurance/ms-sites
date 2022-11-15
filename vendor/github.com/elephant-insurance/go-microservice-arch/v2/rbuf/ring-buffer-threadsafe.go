package rbuf

import (
	"errors"
	"sync"
)

// tsrBuffer (synched struct buffer) is a simple thread-safe implementation of the RingBuffer interface
// All write operations are mutually-exclusive so that the buffer may be used
// from multiple threads without causing race conditions
// This buffer will perform slightly worse than the single-threaded version (rBuffer)
type tsrBuffer struct {
	sync.Mutex
	myBuffer *rBuffer
}

// NewConcurrentRingBuffer creates a thread-safe circular buffer of the specified size.
// You can use a nil second parameter to avoid prefilling the buffer with full entities.
func NewConcurrentRingBuffer(size int, prefiller DummyFactory) (RingBuffer, error) {
	sBuf, err := NewRingBuffer(size, prefiller)
	if err != nil {
		return nil, err
	}

	myBuf, ok := sBuf.(*rBuffer)
	if !ok {
		return nil, errors.New("failed to cast interface to *rBuffer")
	}

	return &tsrBuffer{
		myBuffer: myBuf,
	}, nil
}

// Add places a new item at the end of the buffer
// This alters the buffer and so must be thread-locked.
func (tsrb *tsrBuffer) Add(item interface{}) {
	if tsrb == nil {
		return
	}

	tsrb.Lock()
	defer tsrb.Unlock()
	tsrb.myBuffer.Add(item)
}

// Available is the actual number of entries currently in the buffer.
// We don't bother to thread-lock this because it's just a straight read
func (tsrb *tsrBuffer) Available() int {
	if tsrb == nil {
		return 0
	}

	return tsrb.myBuffer.Available()
}

// Capacity is the total number of items the buffer can hold. It is set at construction and does not change.
// We don't bother to thread-lock this because it's just a straight read
func (tsrb *tsrBuffer) Capacity() int {
	if tsrb == nil {
		return 0
	}

	return tsrb.myBuffer.Capacity()
}

// Earliest returns the earliest count items non-destructively, earliest first, without removing them
// This is a straight read operation, but requires multiple operations to build the return value
// So we thread-lock it
func (tsrb *tsrBuffer) Earliest(count int) []interface{} {
	if tsrb == nil {
		return nil
	}

	tsrb.Lock()
	defer tsrb.Unlock()
	return tsrb.myBuffer.Earliest(count)
}

// First returns the single oldest entry available without removing it
// We don't bother to thread-lock this because it's just a straight read
func (tsrb *tsrBuffer) First() interface{} {
	if tsrb == nil {
		return nil
	}

	return tsrb.myBuffer.First()
}

// Last returns the single latest entry available without removing it
// We don't bother to thread-lock this because it's just a straight read
func (tsrb *tsrBuffer) Last() interface{} {
	if tsrb == nil {
		return nil
	}

	return tsrb.myBuffer.Last()
}

// Latest returns the latest count items non-destructively, latest first, without removing them
// This is a straight read operation, but requires multiple operations to build the return value
// So we thread-lock it
func (tsrb *tsrBuffer) Latest(count int) []interface{} {
	if tsrb == nil {
		return nil
	}

	tsrb.Lock()
	defer tsrb.Unlock()
	return tsrb.myBuffer.Latest(count)
}

func (tsrb *tsrBuffer) Next() interface{} {
	// unfortunately, we can't just wrap this in a mutex.Lock block because
	// we risk halting the thread with a double-lock

	// can't block with a nil buffer
	if tsrb == nil {
		return nil
	}

	// first, try to just return the first entry
	e := tsrb.PullFromEarliest(1)
	if len(e) > 0 {
		return e[0]
	}

	// we weren't able to pull one, so we have to wait
	// create a channel we can listen to, type doesn't matter
	// make it a one-entry buffered channel so that we can write to it without blocking
	c := make(chan bool, 1)
	tsrb.addChannel(c)
	for range c {
		e := tsrb.PullFromEarliest(1)
		if len(e) > 0 {
			return e[0]
		}
	}

	// this should only happen if multiple threads are waiting to read from the buffer
	// we didn't get an entry and our channel has been closed, so we have to return nil
	// because another thread that was waiting got to the entry before this one did
	// The consuming loop should check for a nil return and continue when it gets one back
	return nil
}

// PullFromEarliest removes the count earliest entries from the buffer and returns them, earliest first
// This alters the buffer and so must be thread-locked.
func (tsrb *tsrBuffer) PullFromEarliest(count int) []interface{} {
	if tsrb == nil {
		return nil
	}

	tsrb.Lock()
	defer tsrb.Unlock()
	return tsrb.myBuffer.PullFromEarliest(count)
}

// PullFromLatest removes the count latest entries from the buffer and returns them, latest first
// This alters the buffer and so must be thread-locked.
func (tsrb *tsrBuffer) PullFromLatest(count int) []interface{} {
	if tsrb == nil {
		return nil
	}

	tsrb.Lock()
	defer tsrb.Unlock()
	return tsrb.myBuffer.PullFromLatest(count)
}

// SelectCountFromEarliest non-destructively returns the count earliest entries from the buffer that satisfy the ItemSelector, earliest first
func (tsrb *tsrBuffer) SelectCountFromEarliest(selector ItemSelector, count int) []interface{} {
	return tsrb.myBuffer.SelectCountFromEarliest(selector, count)
}

// SelectCountFromLatest non-destructively returns the count latest entries from the buffer that satisfy the ItemSelector, latest first
func (tsrb *tsrBuffer) SelectCountFromLatest(selector ItemSelector, count int) []interface{} {
	return tsrb.myBuffer.SelectCountFromLatest(selector, count)
}

// SelectRangeFromEarliest returns the continuous range of earliest entries from the buffer that satisfy the ItemSelector, earliest first
// If the earliest element does not satisfy the ItemSelector, returns nil
func (tsrb *tsrBuffer) SelectRangeFromEarliest(selector ItemSelector) []interface{} {
	return tsrb.myBuffer.SelectRangeFromEarliest(selector)
}

// SelectRangeFromLatest returns the continuous range of latest entries from the buffer that satisfy the ItemSelector, latest first
// If the latest element does not satisfy the ItemSelector, returns nil
func (tsrb *tsrBuffer) SelectRangeFromLatest(selector ItemSelector) []interface{} {
	return tsrb.myBuffer.SelectRangeFromLatest(selector)
}

// TotalAdded returns the total number of items added to the buffer since it was created
// We don't bother to thread-lock this because it's just a straight read
func (tsrb *tsrBuffer) TotalAdded() int {
	if tsrb == nil {
		return 0
	}

	return tsrb.myBuffer.TotalAdded()
}

func (tsrb *tsrBuffer) addChannel(c chan bool) {
	if c != nil {
		tsrb.Lock()
		defer tsrb.Unlock()
		tsrb.myBuffer.addChannel(c)
	}
}

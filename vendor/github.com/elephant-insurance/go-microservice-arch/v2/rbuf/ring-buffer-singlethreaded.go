package rbuf

import (
	"fmt"
)

// rBuffer is a simple caching buffer that stores a constant number of entries at all times

// rBuffer is a simple implementation of the RingBuffer interface
// Please note that rBuffer IS NOT THREAD-SAFE
// This buffer will perform slightly better than the concurrent version (tsrBuffer)
type rBuffer struct {
	// available tracks how many "real" entries are in the buffer
	available int
	// blockingChannels is a slice of channels representing threads listening for items to be added
	blockingChannels []chan bool
	// entries is our private slice of entries
	entries []interface{}
	// first is the position of the earliest available entry
	first int
	// next is the position for the next entry to be inserted
	next int
	// totalCount tracks how many entries we've added since construction
	totalCount int
}

// MinRingBufferSize sets the minimum size for a ring buffer
const MinRingBufferSize = 3

// NewRingBuffer creates a circular buffer of the specified size
// You can use a nil second parameter to avoid prefilling the buffer with full entities
func NewRingBuffer(size int, prefiller DummyFactory) (RingBuffer, error) {
	if size < MinRingBufferSize {
		return nil, fmt.Errorf("new RingBuffer has a minimum size of %v", MinRingBufferSize)
	}

	e := make([]interface{}, size)

	if prefiller != nil {
		for i := 0; i < size; i++ {
			e[i] = prefiller()
		}
	}

	bc := make([]chan bool, 0)

	rtn := rBuffer{
		available:        0,
		blockingChannels: bc,
		entries:          e,
		first:            0,
		next:             0,
		totalCount:       0,
	}

	return &rtn, nil
}

// Add adds a new entry to the end of the list,
// overwriting the earliest entry to keep the length constant
func (sb *rBuffer) Add(item interface{}) {
	if sb == nil {
		return
	}

	sb.entries[sb.next] = item
	sb.totalCount++

	// do we have any dummy entries?
	noDummies := false
	if sb.next == sb.first && sb.available > 0 {
		// we don't have any dead/dummy entries
		noDummies = true
	}

	if sb.next+1 < len(sb.entries) {
		sb.next++
	} else {
		sb.next = 0
	}

	if noDummies {
		// we didn't have any empty space, so we had to overwrite the previous last
		sb.first = sb.next
	} else {
		// we overwrote a dead/dummy entry, so we have one more entry available now
		sb.available++
	}

	// release any blocking threads
	if len(sb.blockingChannels) > 0 {
		for _, c := range sb.blockingChannels {
			c <- true
			close(c)
		}

		sb.blockingChannels = []chan bool{}
	}
}

func (sb *rBuffer) Available() int {
	if sb != nil {
		return sb.available
	}

	return 0
}

func (sb *rBuffer) Capacity() int {
	if sb != nil {
		return len(sb.entries)
	}

	return 0
}

// Earliest returns the first count entries added (0 = all)
// skipping any dummies, earliest first
func (sb *rBuffer) Earliest(count int) []interface{} {
	return sb.getRange(count, false, true, false, nil)
}

func (sb *rBuffer) First() interface{} {
	if sb == nil || sb.available < 1 {
		return nil
	}

	return sb.entries[sb.first]
}

func (sb *rBuffer) Last() interface{} {
	if sb == nil || sb.available < 1 {
		return nil
	}

	// if the buffer has just written (or written over) the last element of its entries array,
	// then available > 0 and next == 0, so we have to return the last element
	if sb.next > 0 {
		return sb.entries[sb.next-1]
	}

	return sb.entries[len(sb.entries)-1]
}

// Latest returns the last count entries added (0 = all)
// skipping any dummies, latest first
func (sb *rBuffer) Latest(count int) []interface{} {
	return sb.getRange(count, false, false, false, nil)
}

// Next will remove and return the earliest entry if one is available
// If no entries are available, it will block until it can return one.
func (sb *rBuffer) Next() interface{} {
	// can't block with a nil buffer
	if sb == nil {
		return nil
	}

	// first, try to just return the first entry
	e := sb.getRange(1, true, true, false, nil)
	if len(e) > 0 {
		return e[0]
	}

	// we weren't able to pull one, so we have to wait
	// create a channel we can listen to, type doesn't matter
	// make it a one-entry buffered channel so that we can write to it without blocking
	c := make(chan bool, 1)
	sb.addChannel(c)
	for range c {
		e := sb.getRange(1, true, true, false, nil)
		if len(e) > 0 {
			return e[0]
		}
	}

	// this should only happen if multiple threads are waiting to read from the buffer
	// we didn't get an entry and our channel has been closed, so we have to return nil
	// because another thread that was waiting got to the entry before this one did
	// The consuming loop should check for a nil return and continue when it gets one back
	// rBuffer is NOT the right implementation to use in a multi-threaded environment
	return nil
}

// PullFromEarliest removes and returns the first count entries added (0 = all)
// skipping any dummies, earliest first
func (sb *rBuffer) PullFromEarliest(count int) []interface{} {
	return sb.getRange(count, true, true, false, nil)
}

// PullFromLatest removes and returns the last count entries added (0 = all)
// skipping any dummies, latest first
func (sb *rBuffer) PullFromLatest(count int) []interface{} {
	return sb.getRange(count, true, false, false, nil)
}

func (sb *rBuffer) TotalAdded() int {
	if sb != nil {
		return sb.totalCount
	}

	return 0
}

// SelectCountFromEarliest returns up to count earliest entries satisfying selector
// The entries are not removed, and they may be non-contiguous
func (sb *rBuffer) SelectCountFromEarliest(selector ItemSelector, count int) []interface{} {
	return sb.getRange(count, false, true, false, selector)
}

// SelectCountFromLatest non-destructively returns the count latest entries from the buffer that satisfy the ItemSelector, latest first
func (sb *rBuffer) SelectCountFromLatest(selector ItemSelector, count int) []interface{} {
	return sb.getRange(count, false, false, false, selector)
}

func (sb *rBuffer) SelectRangeFromEarliest(selector ItemSelector) []interface{} {
	return sb.getRange(0, false, true, true, selector)
}

func (sb *rBuffer) SelectRangeFromLatest(selector ItemSelector) []interface{} {
	return sb.getRange(0, false, false, true, selector)
}

// getRange is the private do-all func for getting entries
// Various interface methods just wrap this with certain options pre-selected to make the interface friendlier
func (sb *rBuffer) getRange(count int, remove, earliest, contiguous bool, selector ItemSelector) []interface{} {
	// For now, can't remove if we're using a selector, though it may be possible for contiguous ranges
	if sb == nil || count < 0 || count > len(sb.entries) || (remove && selector != nil) {
		return nil
	}

	sel := selector
	if sel == nil {
		sel = func(interface{}) bool { return true }
	}

	cbcap := len(sb.entries)
	ct := count
	if ct == 0 {
		ct = cbcap
	}

	// ct is the effective number of records requested, 0 < ct <= cbcap

	var rtn []interface{}

	// sb.available is the total number of real, active records currently in the cb
	// we might want more than sb.available back, but we can't have it
	if ct > sb.available {
		// size our return by what we have
		rtn = make([]interface{}, sb.available)
	} else {
		// cb has enough entries to satisfy the request fully
		rtn = make([]interface{}, ct)
	}

	// keep track of the max entries we'll be returning
	rc := len(rtn)

	// get our looping funcs
	// this enables us to use the same code for earliest and latest requests by altering the behavior of the loops appropriately
	cfn1, cfn2, ifn1, ifn2, rfn := getLoopingFuncs(earliest, sb.first, sb.next, sb.available, rc, cbcap)

	// starting with the first or last entry, copy entries into the return value
	// until we either run out of available entries or run out of array space and have to circle around
	// copied keeps track of how many entries we've actually written to the return slice
	// considered keeps track of how many we have looked at
	copied, considered := 0, 0
	// rangeTripped keeps track of whether we've reached the end of a contiguous range, if we're interested in that
	rangeTripped := false
	// loop conditions: cfn1 makes sure we don't run past the end (for earliest queries) or beginning (for latest) entry
	// copied < rc makes sure we don't write more entries than we want to the return value
	// !(contiguous && rangeTripped) makes sure we stop writing ass soon as our selector fails IF we only want contiguous items
	for i := 0; cfn1(i) && copied < rc && !(contiguous && rangeTripped); i++ {
		thisEntry := sb.entries[ifn1(i)]
		if sel(thisEntry) {
			rtn[copied] = thisEntry
			copied++
		} else {
			rangeTripped = true
		}
		considered++
	}

	// no point in continuing if we've already considered every entry
	if considered < sb.available {
		// we have not considered enough entries yet so we know copied must also be less than requested
		// we must have reached the end of the array, beginning or end, so circle back from the other end
		// loop conditions: cfn1 makes sure we don't run past the end (for earliest queries) or beginning (for latest) entry
		// copied < rc makes sure we don't write more entries than we want to the return value
		// !(contiguous && rangeTripped) makes sure we stop writing ass soon as our selector fails IF we only want contiguous items
		for i := 0; cfn2(i) && copied < rc && !(contiguous && rangeTripped); i++ {
			thisEntry := sb.entries[ifn2(i)]
			if sel(thisEntry) {
				rtn[copied] = thisEntry
				copied++
			}
		}
	}

	// if we're removing entries, we need to move the pointers to the ends of the available entries
	// and update the count of available entries
	// The calculation depends on whether we're removing from latest or earliest, so we use a parametric func to figure it
	if remove {
		sb.first, sb.next, sb.available = rfn()
	}

	if copied > 0 {
		return rtn[0:copied]
	}

	return nil
}

// getLoopingFuncs generates parametric functions for controlling loops over ranges of entries
// Since the only difference between "earliest" and "latest" funcs is where they start and the direction they go in,
// generating these on the fly allows us to use the same loops for both "earliest" and "latest" queries
// THIS THING IS PRETTY HARD TO FIGURE OUT!
// If you need to figure it out, I suggest breaking down getRange into two copies, eliminating the "earliest" argument: getEarliestRange and getLatestRange
// From there you can patch in the funcs below, eliminating this function. The result will be a lot more code, but should be easier to understand.
func getLoopingFuncs(earliest bool, first, next, available, rc, capacity int) (func(int) bool, func(int) bool, func(int) int, func(int) int, func() (int, int, int)) {
	if earliest {
		return func(i int) bool {
				return i+first < capacity
			},
			func(i int) bool { return i < next },
			func(i int) int {
				return first + i
			},
			func(i int) int {
				return i
			},
			func() (int, int, int) {
				newFirst := first + rc
				if newFirst >= capacity {
					// we removed enough entries to wrap the "first" pointer
					newFirst -= capacity
				}

				return newFirst, next, available - rc
			}
	}

	return func(i int) bool {
			return next-(i+1) >= 0
		},
		func(i int) bool { return capacity-(i+1) >= first },
		func(i int) int {
			return next - (i + 1)
		},
		func(i int) int {
			return capacity - (i + 1)
		},
		func() (int, int, int) {
			newNext := next - rc
			if newNext < 0 {
				// we removed enough entries to wrap the "next" pointer
				newNext += capacity
			}

			return first, newNext, available - rc
		}
}

func (sb *rBuffer) addChannel(c chan bool) {
	if c != nil {
		sb.blockingChannels = append(sb.blockingChannels, c)
	}
}

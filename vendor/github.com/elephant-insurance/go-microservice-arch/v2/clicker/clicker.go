package clicker

import (
	"strconv"
	"sync"
)

// Clicker is a very simple thread-safe counter
// Use it for updating counts of events for logging and diagnostics
// Seems silly to create a package for this one little type, but this
// keeps us out of dependency trouble
type Clicker struct {
	Clicks int
	sync.Mutex
}

// Click adds count clicks to the counter and returns the result
func (c *Clicker) Click(count int) int {
	if count != 0 {
		c.Lock()
		defer c.Unlock()
		c.Clicks += count
	}

	return c.Clicks
}

// Set sets the Clicker's Clicks to a specified int
func (c *Clicker) Set(count int) *Clicker {
	if count != c.Clicks {
		c.Lock()
		defer c.Unlock()
		c.Clicks = count
	}

	return c
}

// MarshalJSON lets us to output a Clicker struct as if it were a scalar int
func (c *Clicker) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Itoa(c.Clicks)), nil
}

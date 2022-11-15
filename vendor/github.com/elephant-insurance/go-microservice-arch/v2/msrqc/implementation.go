package msrqc

import (
	"context"
	"sync"
	"time"
)

// msrqContext is a simple implementation of the Context interface.
// It's main advantage over the Golang context.Context is that it has a Set method.
// It is also a reference type, making it unnecessary to update and replace.
// Because it implements the most useful features of the gin.Context, we can use this interface
// as an abstraction of the Gin context in packages where we don't want to make a direct reference to Gin.
type msrqContext struct {
	context.Context
	valueMap map[string]interface{}
	// this is just here to protect the valueMap from contention
	sync.Mutex
}

// newMSRQContext creates a msrqc.Context from a given context.Context (or nil)
// By pre-populating the namespaces in the valueMap, we prevent a possible threading issue where
// two different threads attempt to add a namespace to the context at the same time
func newMSRQContext(c context.Context) *msrqContext {
	myContext := c
	if myContext == nil {
		myContext = context.Background()
	}

	return &msrqContext{
		myContext,
		map[string]interface{}{
			namespaceKeyConfig: newNamespace(),
			NamespaceKeyLog:    newNamespace(),
			NamespaceKeySec:    newNamespace(),
		},
		sync.Mutex{},
	}
}

func (c *msrqContext) Deadline() (deadline time.Time, ok bool) {
	if c != nil && c.Context != nil {
		return c.Context.Deadline()
	}

	return
}

func (c *msrqContext) Done() <-chan struct{} {
	if c != nil && c.Context != nil {
		return c.Context.Done()
	}

	return nil
}

func (c *msrqContext) Err() error {
	if c != nil && c.Context != nil {
		return c.Context.Err()
	}

	return nil
}

func (c *msrqContext) Value(key interface{}) interface{} {
	if c == nil || key == nil {
		return nil
	}

	if keyAsString, ok := key.(string); ok {
		val, _ := c.Get(keyAsString)
		return val
	}
	return nil
}

func (c *msrqContext) Get(key string) (value interface{}, exists bool) {
	if c == nil || key == `` {
		return nil, false
	}

	c.Lock()
	defer c.Unlock()

	if val, ok := c.valueMap[key]; ok {
		return val, ok
	}

	if c.Context != nil && c.Context.Value(key) != nil {
		return c.Context.Value(key), true
	}

	return nil, false
}

func (c *msrqContext) Set(key string, value interface{}) {
	if c != nil {
		c.Lock()
		defer c.Unlock()
		c.valueMap[key] = value
	}
}

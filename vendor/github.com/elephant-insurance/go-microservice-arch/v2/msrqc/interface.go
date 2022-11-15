package msrqc

import "time"

type Context interface {
	// These are the base methods for the context.Context interface
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key interface{}) interface{}

	// These are implemented by *gin.Context
	Get(key string) (value interface{}, exists bool)
	Set(key string, value interface{})
}

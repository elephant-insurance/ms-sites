package mbuf

import (
	"encoding/json"
	"fmt"

	"github.com/elephant-insurance/go-microservice-arch/v2/clicker"
)

func NewDefaultFlusher() Flusher {
	return newDefaultFlusher()
}

func newDefaultFlusher() *defaultFlusher {
	return &defaultFlusher{
		messagesFlushed: &clicker.Clicker{},
		flushing:        true,
	}
}

// defaultFlusher is the default implementation of the Flusher interface
// it prints flushed messages to the console using json.Marshal
type defaultFlusher struct {
	messagesFlushed *clicker.Clicker
	flushing        bool
}

func (f *defaultFlusher) Flush(msgs []interface{}) {
	cnt := len(msgs)
	if cnt < 1 {
		return
	}

	for i := 0; i < cnt; i++ {
		if msgs[i] != nil {
			bslice, err := json.Marshal(msgs[i])
			if err == nil && len(bslice) > 0 {
				if f.flushing {
					fmt.Println(string(bslice))
				}
			}
		}
	}

	f.messagesFlushed.Click(cnt)
}

func (f *defaultFlusher) Diagnostics() map[string]interface{} {
	return map[string]interface{}{
		diagnosticsFieldMessagesFlushed: f.messagesFlushed.Clicks,
	}
}

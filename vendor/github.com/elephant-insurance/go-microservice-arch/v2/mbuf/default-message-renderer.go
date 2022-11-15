package mbuf

import (
	"encoding/json"

	"github.com/elephant-insurance/go-microservice-arch/v2/clicker"
)

func NewDefaultMessageRenderer() MessageRenderer {
	return &defaultMessageRenderer{
		failedToRender: &clicker.Clicker{},
	}
}

// defaultMessageRenderer is the default implementation of the MessageRenderer interface
// it renders individual messgaes as JSON by using JSON.Marshal
type defaultMessageRenderer struct {
	failedToRender *clicker.Clicker
}

func (r *defaultMessageRenderer) Render(msg interface{}) []byte {
	rtn, err := json.Marshal(msg)
	if err == nil && len(rtn) > 0 {
		return rtn
	}

	r.failedToRender.Click(1)
	return nil
}

func (r *defaultMessageRenderer) Diagnostics() map[string]interface{} {
	return map[string]interface{}{
		diagnosticsFieldFailuresToRender: r.failedToRender.Clicks,
	}
}

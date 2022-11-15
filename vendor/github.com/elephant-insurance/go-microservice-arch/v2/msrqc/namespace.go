package msrqc

import "sync"

// Namespace represents a semi-private context storage container
type Namespace interface {
	// Dump returns a copy of the entire contents of the Namespace
	// Use this only when all fields are needed
	Dump() map[string]interface{}
	Get(key string) (value interface{}, exists bool)
	Set(key string, val interface{}, overwrite bool)
}

// namespaceType is an implementation of Namespace
type namespaceType struct {
	valDict map[string]interface{}
	sync.Mutex
}

// newNamespace creates a valid namespaceType
func newNamespace() *namespaceType {
	return &namespaceType{
		valDict: map[string]interface{}{},
	}
}

func (n *namespaceType) Dump() map[string]interface{} {
	rtn := make(map[string]interface{}, len(n.valDict))
	n.Lock()
	defer n.Unlock()

	for k, v := range n.valDict {
		rtn[k] = v
	}

	return rtn
}

func (n *namespaceType) Get(key string) (value interface{}, exists bool) {
	if n == nil || key == `` || n.valDict == nil || len(n.valDict) == 0 {
		return nil, false
	}

	n.Lock()
	defer n.Unlock()

	value, exists = n.valDict[key]

	return
}

func (n *namespaceType) Set(key string, val interface{}, overwrite bool) {
	n.Lock()
	defer n.Unlock()
	// can't call Get here or we get thread locked 1
	if _, exists := n.valDict[key]; !exists || overwrite {
		n.valDict[key] = val
	}
}

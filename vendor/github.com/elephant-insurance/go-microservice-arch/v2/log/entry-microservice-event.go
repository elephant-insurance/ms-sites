package log

import (
	"runtime"
	"time"

	enum "github.com/elephant-insurance/enumerations/v2"
)

type MSEventLogEntry struct {
	// uf.Event fields
	// can't make this a nested struct, unfortunately, because those don't serialize well
	Description   string                `json:"Description,omitempty"`
	ID            *enum.EventID         `json:"ID"`
	Severity      *enum.EventSeverityID `json:"Severity"`
	TargetService *enum.ServiceID       `json:"TargetService,omitempty"`

	CommonErrorCode      *enum.CommonErrorCodeID `json:"CommonErrorCode,omitempty"`
	DurationMicroseconds *int64                  `json:"DurationMicroseconds,omitempty"`
	Error                *string                 `json:"err,omitempty"`
	HostName             string                  `json:"HostName"`
	InstanceName         string                  `json:"InstanceName"`
	RAMAlloc             uint64                  `json:"RAMBytesAllocated"`
	RAMHeapObjects       uint64                  `json:"RAMHeapObjectCount"`
	RAMSys               uint64                  `json:"RAMBytesReserved"`
	Time                 time.Time               `json:"Time"`
	TransactionID        *string                 `json:"TransactionID,omitempty"`
	TransactionType      *string                 `json:"TransactionType,omitempty"`
}

func (e *MSEventLogEntry) addMemStats() {
	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)

	e.RAMAlloc = stats.Alloc
	e.RAMHeapObjects = stats.HeapObjects
	e.RAMSys = stats.Sys
}

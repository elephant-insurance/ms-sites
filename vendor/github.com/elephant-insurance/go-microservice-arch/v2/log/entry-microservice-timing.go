package log

import (
	"time"

	enum "github.com/elephant-insurance/enumerations/v2"
	"github.com/elephant-insurance/go-microservice-arch/v2/msrqc"
	"github.com/elephant-insurance/go-microservice-arch/v2/timing"
)

type MSTimingLogEntry struct {
	BusinessType         *enum.BusinessTypeID     `json:"BusinessType,omitempty"`
	DurationMicroseconds int64                    `json:"DurationMicroseconds"`
	HostName             string                   `json:"HostName"`
	InstanceName         string                   `json:"InstanceName"`
	Label                *string                  `json:"Label"`
	Path                 *string                  `json:"Path"`
	SequenceNumber       uint64                   `json:"SequenceNumber"`
	Status               int                      `json:"Status"`
	TimeInitiated        time.Time                `json:"TimeInitiated"`
	TransactionID        *string                  `json:"TransactionID"`
	TransactionSource    *enum.SourceOfBusinessID `json:"txsource,omitempty"`
	Type                 timing.TimingType        `json:"TimingType"`
}

// NewTimingEntry makes a safe copy of all fields in the passed-in timing
// and adds in HostName, InstanceName, Source of Business, and TXID
func NewTimingEntry(c msrqc.Context, t *timing.Timing) *MSTimingLogEntry {
	if t == nil {
		return nil
	}

	var lab, pat *string
	var txid string
	var txsrc enum.SourceOfBusinessID
	var txbt enum.BusinessTypeID

	if t.Label != nil {
		foo := *t.Label
		lab = &foo
	}

	if t.Path != nil {
		foo := *t.Path
		pat = &foo
	}

	txid = msrqc.GetTransactionID(c)
	if val := msrqc.GetTransactionSource(c); val != nil {
		txsrc = *val
	}

	if tbt, ok := msrqc.NamespaceGet(c, &msrqc.NamespaceKeyLog, &fieldNameMSBusinessType); ok {
		if foo, ok := tbt.(*enum.BusinessTypeID); ok && foo != nil {
			txbt = *foo
		}
	}

	return &MSTimingLogEntry{
		DurationMicroseconds: t.DurationMicroseconds,
		HostName:             defaultTimingLog.hostName,
		InstanceName:         defaultTimingLog.instanceName,
		Label:                lab,
		Path:                 pat,
		SequenceNumber:       t.SequenceNumber,
		Status:               t.Status,
		TimeInitiated:        t.TimeInitiated,
		TransactionID:        &txid,
		TransactionSource:    &txsrc,
		Type:                 t.Type,
		BusinessType:         &txbt,
	}
}

package log

import "time"

type RTBLogEntry struct {
	AppName       *string     `json:"app,omitempty"`
	AuctionId     *string     `json:"id,omitempty"`
	DurationMs    *int        `json:"duration,omitempty"`
	Headers       interface{} `json:"headers,omitempty"`
	HeadersJSON   string      `json:"headers_json,omitempty"`
	HTTPStatus    *int        `json:"status,omitempty"`
	Impression    *bool       `json:"impression,omitempty"`
	IP            *string     `json:"ip,omitempty"`
	OfferCents    *int        `json:"cents,omitempty"`
	Rank          *int64      `json:"rank,omitempty"`
	RequestJSON   string      `json:"request_json,omitempty"`
	RequestObject interface{} `json:"request,omitempty"`
	SlotName      *string     `json:"slot,omitempty"`
	Timestamp     *time.Time  `json:"timestamp,omitempty"`
}

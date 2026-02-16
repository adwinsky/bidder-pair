package producer

import (
	"time"

	"bidder-pair/internal/openrtb"
)

type Event struct {
	Type       string
	Request    *openrtb.BidRequest
	RawRequest []byte
	Meta       map[string]string
	At         time.Time
}

package parser

import (
	"encoding/json"
	"errors"

	"bidder-pair/internal/openrtb"
)

func ParseBidRequest(raw []byte) (*openrtb.BidRequest, error) {
	var br openrtb.BidRequest
	if err := json.Unmarshal(raw, &br); err != nil {
		return nil, err
	}
	if br.ID == "" {
		return nil, errors.New("missing id")
	}
	if len(br.Imp) == 0 {
		return nil, errors.New("missing imp")
	}
	if br.User.ID == "" {
		br.User.ID = "anon"
	}
	return &br, nil
}

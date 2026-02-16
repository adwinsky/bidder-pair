package filter

import (
	"bytes"
	"strconv"

	"bidder-pair/internal/openrtb"
)

func buildFeatures(br *openrtb.BidRequest) []byte {
	var buf bytes.Buffer

	buf.WriteString(br.User.ID)
	buf.WriteByte('|')

	for _, imp := range br.Imp {
		buf.WriteString(imp.ID)
		buf.WriteByte(',')
	}

	buf.WriteString(strconv.Itoa(len(br.Device.UA)))

	return buf.Bytes()
}

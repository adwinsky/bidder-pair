package response

import (
	"bytes"
	"strconv"

	"bidder-pair/internal/decision"
	"bidder-pair/internal/openrtb"
)

func BuildJSON(br *openrtb.BidRequest, dec decision.DecisionResponse) []byte {
	var buf bytes.Buffer

	buf.WriteString(`{"id":"`)
	buf.WriteString(br.ID)
	buf.WriteString(`","seatbid":[{"bid":[{`)

	buf.WriteString(`"id":"bid-1",`)
	buf.WriteString(`"impid":"`)
	buf.WriteString(br.Imp[0].ID)
	buf.WriteString(`",`)

	buf.WriteString(`"price":`)
	buf.WriteString(strconv.FormatFloat(dec.Price, 'f', 3, 64))
	buf.WriteByte(',')

	buf.WriteString(`"crid":"`)
	buf.WriteString(dec.CreativeID)
	buf.WriteString(`"`)

	buf.WriteString(`}]}]}`)

	return buf.Bytes()
}

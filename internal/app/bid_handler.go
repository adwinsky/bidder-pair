package app

import (
	"io"
	"net/http"
	"time"

	"bidder-pair/internal/decision"
	"bidder-pair/internal/filter"
	"bidder-pair/internal/parser"
	"bidder-pair/internal/producer"
	"bidder-pair/internal/response"
)

func BidHandler(
	f *filter.Filter,
	dc *decision.Client,
	kp *producer.Producer,
) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		br, err := parser.ParseBidRequest(body)
		if err != nil {
			http.Error(w, "invalid request", http.StatusBadRequest)
			return
		}

		ok, reason := f.Filter(br)
		if !ok {
			kp.Send(producer.Event{
				Type:       "filtered",
				Request:    br,
				RawRequest: body,
				At:         time.Now(),
				Meta:       map[string]string{"reason": reason},
			})
			w.WriteHeader(http.StatusNoContent)
			return
		}

		decResp, err := dc.Decide(r.Context(), decision.DecisionRequest{
			RequestID: br.ID,
			UserID:    br.User.ID,
			ImpIDs:    br.ImpIDs(),
		})
		if err != nil {
			http.Error(w, "decision error", http.StatusBadGateway)
			return
		}

		respBody := response.BuildJSON(br, decResp)

		kp.Send(producer.Event{
			Type:       "bid",
			Request:    br,
			RawRequest: body,
			At:         time.Now(),
			Meta: map[string]string{
				"latency": time.Since(start).String(),
			},
		})

		w.Header().Set("Content-Type", "application/json")
		w.Write(respBody)
	}
}

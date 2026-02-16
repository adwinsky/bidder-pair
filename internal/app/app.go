package app

import (
	"net/http"

	"bidder-pair/internal/decision"
	"bidder-pair/internal/filter"
	"bidder-pair/internal/producer"
)

type App struct {
	handler http.Handler
}

type Config struct {
	Filter   *filter.Filter
	Decision *decision.Client
	Producer *producer.Producer
}

func New(cfg Config) *App {
	mux := http.NewServeMux()

	mux.HandleFunc("/openrtb2/bid", BidHandler(
		cfg.Filter,
		cfg.Decision,
		cfg.Producer,
	))

	return &App{
		handler: mux,
	}
}

func (a *App) Handler() http.Handler {
	return a.handler
}

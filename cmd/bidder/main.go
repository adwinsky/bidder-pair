package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"bidder-pair/internal/app"
	"bidder-pair/internal/decision"
	"bidder-pair/internal/filter"
	"bidder-pair/internal/producer"
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags)

	ds := decision.NewServer(decision.ServerConfig{
		Addr:   ":8081",
		Logger: logger,
	})
	go ds.Run()

	kp := producer.NewProducer(producer.ProducerConfig{
		Logger:        logger,
		FlushInterval: 200 * time.Millisecond,
		SlowSend:      40 * time.Millisecond,
	})

	f := filter.New(filter.Config{MaxPerMinute: 120})

	dc := decision.NewClient(decision.ClientConfig{
		BaseURL: "http://127.0.0.1:8081",
		Client:  &http.Client{Timeout: 300 * time.Millisecond},
	})

	a := app.New(app.Config{
		Filter:   f,
		Decision: dc,
		Producer: kp,
	})

	srv := app.NewServer(":8080", a.Handler())

	go srv.Run()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	srv.Shutdown(ctx)
	ds.Shutdown(ctx)
	kp.Close(ctx)
}

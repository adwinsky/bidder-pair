package producer

import (
	"context"
	"log"
	"sync"
	"time"
)

type ProducerConfig struct {
	Logger        *log.Logger
	FlushInterval time.Duration
	SlowSend      time.Duration
}

type Producer struct {
	logger *log.Logger

	mu    sync.Mutex
	queue []Event

	flushInterval time.Duration
	slowSend      time.Duration

	stop chan struct{}
	done chan struct{}
}

func NewProducer(cfg ProducerConfig) *Producer {
	p := &Producer{
		logger:        cfg.Logger,
		flushInterval: cfg.FlushInterval,
		slowSend:      cfg.SlowSend,
		stop:          make(chan struct{}),
		done:          make(chan struct{}),
		queue:         make([]Event, 0, 1024),
	}
	go p.loop()
	return p
}

func (p *Producer) Send(e Event) {
	p.mu.Lock()
	p.queue = append(p.queue, e)
	p.mu.Unlock()
}

func (p *Producer) loop() {
	t := time.NewTicker(p.flushInterval)
	defer t.Stop()
	defer close(p.done)

	for {
		select {
		case <-p.stop:
			return
		case <-t.C:
			p.flushOnce()
		}
	}
}

func (p *Producer) flushOnce() {
	p.mu.Lock()
	if len(p.queue) == 0 {
		p.mu.Unlock()
		return
	}

	batch := p.queue
	p.queue = make([]Event, 0, 1024)
	p.mu.Unlock()

	// send request to external service
	time.Sleep(p.slowSend * time.Duration(len(batch)))
}

func (p *Producer) Close(ctx context.Context) error {
	close(p.stop)
	select {
	case <-p.done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

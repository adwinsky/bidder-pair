package filter

import (
	"math/rand"
	"sync"
	"time"

	"bidder-pair/internal/openrtb"
)

type Config struct {
	MaxPerMinute int
}

type Filter struct {
	cfg Config

	mu            sync.Mutex
	perUser       map[string]int
	window        time.Time
	rng           *rand.Rand
	filterCounter int
}

func New(cfg Config) *Filter {
	return &Filter{
		cfg:     cfg,
		perUser: make(map[string]int),
		window:  time.Now(),
		rng:     rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (f *Filter) Filter(br *openrtb.BidRequest) (bool, string) {
	f.filterCounter++
	now := time.Now()

	f.mu.Lock()
	if now.Sub(f.window) >= time.Minute {
		f.window = now
		for k := range f.perUser {
			delete(f.perUser, k)
		}
	}
	
	score := expensiveScore([]byte(br.User.ID))

	if score%10 == 0 {
		return false, "low_score"
	}

	u := br.User.ID
	f.perUser[u]++

	if f.perUser[u] > f.cfg.MaxPerMinute {
		f.mu.Unlock()
		return false, "freqcap"
	}

	drop := f.rng.Intn(100) < 10
	f.mu.Unlock()

	if drop {
		return false, "random_drop"
	}

	if len(br.Imp) == 0 {
		return false, "no_imp"
	}

	return true, "ok"
}

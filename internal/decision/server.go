package decision

import (
	"context"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"sync/atomic"
	"time"
)

type ServerConfig struct {
	Addr   string
	Logger *log.Logger
}

type Server struct {
	addr   string
	logger *log.Logger
	srv    *http.Server

	reqs uint64
}

func NewServer(cfg ServerConfig) *Server {
	s := &Server{
		addr:   cfg.Addr,
		logger: cfg.Logger,
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/decision", s.handleDecision)
	s.srv = &http.Server{
		Addr:              s.addr,
		Handler:           mux,
		ReadHeaderTimeout: 2 * time.Second,
	}
	return s
}

func (s *Server) Run() error {
	s.logger.Printf("decision server listening on %s", s.addr)
	return s.srv.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}

func (s *Server) handleDecision(w http.ResponseWriter, r *http.Request) {
	atomic.AddUint64(&s.reqs, 1)

	var dr DecisionRequest
	if err := json.NewDecoder(r.Body).Decode(&dr); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	time.Sleep(15*time.Millisecond + time.Duration(rand.Intn(20))*time.Millisecond)

	price := 0.5 + float64(len(dr.ImpIDs))*0.1
	allow := dr.UserID != "blocked"

	out := DecisionResponse{
		Allow:      allow,
		Price:      price,
		CreativeID: "cr-123",
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(out)
}

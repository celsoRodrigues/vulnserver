package server

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"time"
)

type server struct {
	http.Server
	logger *slog.Logger
}

type Options func(s *server) error

func New(host, port string, opts ...Options) *server {
	addr := fmt.Sprintf("%s:%s", host, port)
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	s := &server{}
	s.Addr = addr
	s.Handler = mux
	s.logger = slog.Default()

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func (s *server) Start() {
	log.Printf("server listening on http://%s\n", s.Addr)
	log.Println("server failed with:", s.ListenAndServe())
}

func (s *server) PrintIdleTimeout() {
	log.Println("IdleTimeout is:", s.IdleTimeout)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("NO ROUTES REGISTERED"))
}

func WithIdleTimeout(t time.Duration) Options {
	return func(s *server) error {
		s.IdleTimeout = t
		return nil
	}
}

func WithRoutes(m map[string]func(w http.ResponseWriter, r *http.Request)) Options {
	return func(s *server) error {
		mux := http.NewServeMux()
		for k, v := range m {
			mux.HandleFunc(k, v)
		}
		s.Handler = mux
		return nil
	}
}

func WithLogger(logger *slog.Logger) Options {
	slog.SetDefault(logger)
	return func(s *server) error {
		s.logger = logger
		return nil
	}
}

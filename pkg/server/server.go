package server

import (
	"context"
	"net/http"
	"time"

	"github.com/chatbot/pkg/logger"
)

const (
	_defaultReadTimeout     = 5 * time.Second
	_defaultWriteTimeout    = 5 * time.Second
	_defaultAddr            = ":80"
	_defaultShutdownTimeout = 3 * time.Second
)

// Server ...
type Server struct {
	server          *http.Server
	notify          chan error
	shutdownTimeout time.Duration
	logger          logger.Interface
}

// New ...
func New(handler http.Handler, logger logger.Interface, opts ...Option) *Server {
	httpServer := &http.Server{
		Handler:      handler,
		Addr:         _defaultAddr,
	}
	ShutdownTimeout(_defaultWriteTimeout)

	s := &Server{
		server:          httpServer,
		notify:          make(chan error, 1),
		logger:          logger,
	}
	ReadTimeout(_defaultReadTimeout)
	WriteTimeout(_defaultWriteTimeout)
	ShutdownTimeout(_defaultShutdownTimeout)

	// Custom options
	for _, opt := range opts {
		opt(s)
	}

	s.start()

	return s
}

// start ...
func (s *Server) start() {
	go func() {
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

// Notify ...
func (s *Server) Notify() <-chan error {
	return s.notify
}

// Shutdown ...
func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}

package DBDocument

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run(port string, handlers *mux.Router) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handlers,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1 mb
	}

	log.Printf("ListenAndServe on localhost:%s", port)

	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

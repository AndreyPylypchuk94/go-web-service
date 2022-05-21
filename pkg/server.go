package pkg

import (
	"context"
	"net/http"
)

type Server struct {
	server *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.server = &http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}

	return s.server.ListenAndServe()
}

func (s *Server) Off(cnt context.Context) error {
	return s.server.Shutdown(cnt)
}

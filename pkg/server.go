package pkg

import (
	"context"
	"net/http"
	"pylypchuk.home/internal/api"
	context2 "pylypchuk.home/pkg/context"
)

type Server struct {
	server  *http.Server
	handler *api.Handler
}

func NewServer() *Server {
	return &Server{
		handler: context2.Get("handler").(*api.Handler),
	}
}

func (s *Server) Run(port string) error {
	s.server = &http.Server{
		Addr:    ":" + port,
		Handler: s.handler.InitRouts(),
	}

	return s.server.ListenAndServe()
}

func (s *Server) Off(cnt context.Context) error {
	return s.server.Shutdown(cnt)
}

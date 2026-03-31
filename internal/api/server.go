package api

import (
	"net/http"

	"github.com/ryoshu404/gorelate/internal/store"
)

// Server is the read-only HTTP API layer.
type Server struct {
	store store.Store
	mux   *http.ServeMux
}

func New(s store.Store) *Server {
	srv := &Server{
		store: s,
		mux:   http.NewServeMux(),
	}
	srv.registerRoutes()
	return srv
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

package api

import (
	"encoding/json"
	"net/http"
)

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (s *Server) handleGetIOC(w http.ResponseWriter, r *http.Request) {
	// TODO: implement
}

func (s *Server) handleGetIOCFull(w http.ResponseWriter, r *http.Request) {
	// TODO: implement
}

func (s *Server) handleGetIOCSources(w http.ResponseWriter, r *http.Request) {
	// TODO: implement
}

func (s *Server) handleGetIOCConfidence(w http.ResponseWriter, r *http.Request) {
	// TODO: implement
}

func (s *Server) handleQueryIOCs(w http.ResponseWriter, r *http.Request) {
	// TODO: implement — query params: q, type, confidence, tag
}

func (s *Server) handleListSummaries(w http.ResponseWriter, r *http.Request) {
	// TODO: implement — paginated
}

func (s *Server) handleLatestSummaries(w http.ResponseWriter, r *http.Request) {
	// TODO: implement — one per type (emerging, daily)
}

func (s *Server) handleGetSummary(w http.ResponseWriter, r *http.Request) {
	// TODO: implement
}

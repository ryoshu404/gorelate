package api

func (s *Server) registerRoutes() {
	s.mux.HandleFunc("GET /health", s.handleHealth)
	s.mux.HandleFunc("GET /ioc/{indicator}", s.handleGetIOC)
	s.mux.HandleFunc("GET /ioc/{indicator}/full", s.handleGetIOCFull)
	s.mux.HandleFunc("GET /ioc/{indicator}/sources", s.handleGetIOCSources)
	s.mux.HandleFunc("GET /ioc/{indicator}/confidence", s.handleGetIOCConfidence)
	s.mux.HandleFunc("GET /iocs", s.handleQueryIOCs)
	s.mux.HandleFunc("GET /summaries", s.handleListSummaries)
	s.mux.HandleFunc("GET /summaries/latest", s.handleLatestSummaries)
	s.mux.HandleFunc("GET /summaries/{id}", s.handleGetSummary)
}

package api

import (
	"net/http"

	"github.com/MatthewAraujo/health-checker-website/cmd/service/check"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{addr: addr}
}

func (s *APIServer) Start() error {

	router := http.NewServeMux()
	checkHandler := check.NewHandler()
	checkHandler.RegisterRoutes(router)

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	return http.ListenAndServe(s.addr, router)
}

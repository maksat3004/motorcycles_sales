// 4. infrastructure/http/http_server.go
package http

import (
	"log"
	"net/http"
)

type Server struct {
	Address string
}

func NewServer(address string) *Server {
	return &Server{Address: address}
}

func (s *Server) Start(routes func()) {
	routes()
	log.Printf("Server running on %s", s.Address)
	if err := http.ListenAndServe(s.Address, nil); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}

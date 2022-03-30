package sever

import "github.com/gorilla/mux"

type Server struct {
	*mux.Router
}

func NewServer(r *mux.Router) *Server {
	return &Server{
		Router: r,
	}
}

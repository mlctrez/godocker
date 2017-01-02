package server

import "net/http"

type Server struct {
}

func New() *Server {
	return &Server{}
}

func (s *Server) Start() {
	http.HandleFunc("/", func(rw http.ResponseWriter, rq *http.Request) {
		rw.Write([]byte("hello world\n"))
	})
	http.ListenAndServe(":8080", nil)
}

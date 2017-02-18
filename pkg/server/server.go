package server

import (
	"fmt"
	"io"
	"net/http"
)

// Server is currently just an empty struct that can be added to
type Server struct{}

// New constructs a new server
func New() *Server {
	return &Server{}
}

func (s *Server) sslTest(rw http.ResponseWriter, rq *http.Request) {
	resp, err := http.Get("https://mlctrez.com")
	if err != nil {
		fmt.Fprintf(rw, "error retrieving https://mlctrez.com  %s", err)
		return
	}
	io.Copy(rw, resp.Body)
}

// Start initiates http.ListenAndServe on this server
func (s *Server) Start() error {

	http.HandleFunc("/", func(rw http.ResponseWriter, rq *http.Request) {
		rw.Write([]byte("hello world\n"))
	})

	http.HandleFunc("/ssltest", s.sslTest)

	return http.ListenAndServe(":8080", nil)
}

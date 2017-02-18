package server

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
)

// CloseableServer contains the quit channel and the http.Server
type CloseableServer struct {
	quit chan os.Signal
	srv  *http.Server
}

// New constructs a new server
func New() *CloseableServer {
	s := &CloseableServer{
		quit: make(chan os.Signal),
	}
	signal.Notify(s.quit, os.Interrupt)
	return s
}

func (s *CloseableServer) sslTest(rw http.ResponseWriter, rq *http.Request) {
	resp, err := http.Get("https://mlctrez.com")
	if err != nil {
		fmt.Fprintf(rw, "error retrieving https://mlctrez.com  %s", err)
		return
	}
	io.Copy(rw, resp.Body)
}

func (s *CloseableServer) index(rw http.ResponseWriter, rq *http.Request) {
	rw.Write([]byte("hello world"))
}

func (s *CloseableServer) favIcon(rw http.ResponseWriter, rq *http.Request) {
	http.ServeFile(rw, rq, "static/favicon.ico")
}

func (s *CloseableServer) shutdownRoutine() {
	<-s.quit
	log.Println("shutting down server on interrupt")
	if err := s.srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("could not shut down server: %v", err)
	}
}

// Start initiates http.ListenAndServe on this server
func (s *CloseableServer) Start() error {

	mux := http.NewServeMux()
	mux.HandleFunc("/", s.index)
	mux.HandleFunc("/favicon.ico", s.favIcon)
	mux.HandleFunc("/ssltest", s.sslTest)

	s.srv = &http.Server{Addr: ":8080", Handler: mux}

	go s.shutdownRoutine()

	err := s.srv.ListenAndServe()

	if err.Error() != http.ErrServerClosed.Error() {
		return err
	}
	return nil

}

package server

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/go-errors/errors"
	"io"
	"io/ioutil"
	"net/http"
)

type Server struct {
	httpClient *http.Client
}

func New() *Server {
	return &Server{}
}

func (s *Server) setupClient() error {
	pool := x509.NewCertPool()

	if pems, err := ioutil.ReadFile("ca-certificates.crt"); err != nil {
		return err
	} else {
		if !pool.AppendCertsFromPEM(pems) {
			return errors.New("unable to load certificates")
		}
	}

	s.httpClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{RootCAs: pool},
		},
	}
	return nil
}

func (s *Server) sslTest(rw http.ResponseWriter, rq *http.Request) {
	resp, err := s.httpClient.Get("https://mlctrez.com")
	if err != nil {
		fmt.Fprintf(rw, "error retrieving https://mlctrez.com  %s", err)
		return
	}
	io.Copy(rw, resp.Body)
}

func (s *Server) Start() error {

	if err := s.setupClient(); err != nil {
		return err
	}

	http.HandleFunc("/", func(rw http.ResponseWriter, rq *http.Request) {
		rw.Write([]byte("hello world\n"))
	})

	http.HandleFunc("/ssltest", s.sslTest)

	return http.ListenAndServe(":8080", nil)
}

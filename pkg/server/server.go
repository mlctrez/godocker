package server

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
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

func (s *CloseableServer) sslTest(ctx *gin.Context) {
	resp, err := http.Get("https://example.com")
	if err != nil {
		_ = ctx.Error(fmt.Errorf("error retrieving https://example.com %s", err))
		return
	}
	_, _ = io.Copy(ctx.Writer, resp.Body)
}

func (s *CloseableServer) shutdownRoutine() {
	<-s.quit
	fmt.Println("shutting down server on interrupt")
	if err := s.srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("could not shut down server: %v", err)
	}
}

// Start initiates http.ListenAndServe on this server
func (s *CloseableServer) Start() error {

	gin.SetMode(gin.ReleaseMode)
	mux := gin.Default()

	mux.Use(static.Serve("/", static.LocalFile("./static", false)))
	mux.GET("/ssltest", s.sslTest)

	s.srv = &http.Server{Addr: ":8080", Handler: mux}

	go s.shutdownRoutine()

	err := s.srv.ListenAndServe()

	if err.Error() != http.ErrServerClosed.Error() {
		return err
	}
	return nil

}

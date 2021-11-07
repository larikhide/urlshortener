package server

import (
	"context"
	"net/http"
	"time"

	"github.com/larikhide/urlshortener/app/repos/urls"
	"github.com/larikhide/urlshortener/app/shortener"
)

type Server struct {
	srv http.Server
	us  *urls.URLs
	sh  *shortener.Shortener
}

func NewServer(addr string, h http.Handler) *Server {
	s := &Server{}

	s.srv = http.Server{
		Addr:              addr,
		Handler:           h,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
	}
	return s
}

func (s *Server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	s.srv.Shutdown(ctx)
	cancel()
}

func (s *Server) Start(us *urls.URLs, sh *shortener.Shortener) {
	s.us = us
	s.sh = sh
	go s.srv.ListenAndServe()
}

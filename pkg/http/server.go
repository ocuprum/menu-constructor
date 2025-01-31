package http

import (
	"fmt"
	"net/http"
)

var srv *http.Server							   

func NewMux() *http.ServeMux {
	mux := http.NewServeMux()

	return mux
}

func NewServer(conf Config, hh ...Handler) *http.Server {
	mux := NewMux()
	for _, h := range hh {
		h.Register(mux)
	}

	srv = &http.Server{
		Addr: fmt.Sprintf(":%v", conf.Port),
		Handler: mux,
	}

	return srv 
}
package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
)

type Backend interface {
	SetAlive(bool)
	IsAlive() bool
	GetURL() *url.URL
	GetActiveConnections() int
	Serve(http.ResponseWriter, *http.Request)
}

type backend struct {
	url         *url.URL
	alive       bool
	mux         sync.RWMutex
	connections int
	reveseProxy *httputil.ReverseProxy
}

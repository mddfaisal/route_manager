// File: routing.go
// @Author: Faisal
// @Date: 2023-10-01
// @License: MIT License
// @Package: main
// @Description: This file contains the routing logic for the HTTP server.

package main

import (
	"net/http"

	routemanager "github.com/mddfaisal/route_manager"
)

type Server struct {
	Property1 string
	Property2 int
	Property3 bool
	Property4 float64
}

type HttpHandler func(*Server, http.ResponseWriter, *http.Request)

func (h HttpHandler) GoToNext(s *Server, rw http.ResponseWriter, r *http.Request) {
	h(s, rw, r)
}

type HandlerElement struct {
	Handler     HttpHandler
	Middlewares []Middleware
}

type router struct {
	// handler -> map[http_method][url:handler]
	Routes      map[string]routemanager.Node
	HttpHandler *Controller
	Server      *Server
}

func NewRouter(server *Server) *router {
	r := &router{
		Routes:      make(map[string]routemanager.Node),
		HttpHandler: &Controller{},
		Server:      server,
	}
	return r
}

func (r router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// Find the route for the incoming request

}

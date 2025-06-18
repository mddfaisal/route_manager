package main

import "net/http"

type Middleware func(next HttpHandler) HttpHandler

func SampleMiddleware(next HttpHandler) HttpHandler {
	return func(s *Server, rw http.ResponseWriter, r *http.Request) {

		next.GoToNext(s, rw, r)
	}
}

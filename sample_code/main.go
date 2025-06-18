// File: main.go
// @Author: Faisal
// @Date: 2023-10-01
// @License: MIT License
// @Package: main
// @Description: This file initializes the HTTP server and starts listening for incoming requests.

package main

import (
	"net/http"
)

func main() {
	srv := http.Server{
		Handler: NewRouter(&Server{
			Property1: "value1",
			Property2: 42,
			Property3: true,
			Property4: 3.14,
		}),
		Addr: "localhost:8080",
	}
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}

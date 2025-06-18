// File: homecontroller.go
// @Author: Faisal
// @Date: 2023-10-01
// @License: MIT License
// @Package: main
// @Description: This file contains the home controller logic for the HTTP server.

package main

import (
	"fmt"
	"net/http"
)

type HomeController struct {
	h *Controller
}

func NewHomeController() *HomeController {
	a := new(HomeController)
	a.h = NewBaseController()
	return a
}

func (h *HomeController) Get(s *Server, rw http.ResponseWriter, r *http.Request) {
	fmt.Println("HomeController called...")
	h.h.SetPayload("Hello World").Response200(&rw)
}

func (h *HomeController) Post(s *Server, rw http.ResponseWriter, r *http.Request) {
	fmt.Println("HomeController Post called...")
	h.h.SetPayload("Hello World from Post").Response200(&rw)
}
func (h *HomeController) Update(s *Server, rw http.ResponseWriter, r *http.Request) {
	fmt.Println("HomeController Post called...")
	h.h.SetPayload("Hello World from Post").Response200(&rw)
}
func (h *HomeController) Delete(s *Server, rw http.ResponseWriter, r *http.Request) {
	fmt.Println("HomeController Delete called...")
	h.h.SetPayload("Hello World from Delete").Response200(&rw)
}

// Package main provides a basic HTTP controller with methods for handling responses.
// It includes methods for setting payloads, handling errors, and rendering JSON responses.
// It also supports CORS handling and provides a NotFound response method.
// The Controller struct is designed to be used in a web server context, allowing for easy response
// @Author: Faisal
// @Date: 2023-10-01
// @License: MIT License
// @Package: main
// @Description: This file contains the base controller logic for the HTTP server.

package main

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Payload interface{}       `json:"payload"`
	Errors  interface{}       `json:"errors"`
	headers map[string]string `json:"-"`
}

type Controller struct {
	Payload *response
}

func NewBaseController() *Controller {
	h := new(Controller)
	h.Payload = new(response)
	h.Payload.headers = make(map[string]string, 0)
	return h
}

func (h *Controller) NotFound(rw *http.ResponseWriter) {
	if h.Payload == nil {
		h = NewBaseController()
	}
	h.Payload.Payload = "Not Found"
	(*rw).WriteHeader(http.StatusNotFound)
	h.Render(rw)
}

func (h *Controller) HandleCors(rw *http.ResponseWriter) {
	(*rw).Header().Set("Access-Control-Allow-Origin", "*")
	(*rw).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*rw).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization, token, verify")
}

func (h *Controller) SetPayload(i interface{}) *Controller {
	h.Payload.Payload = i
	return h
}

func (h *Controller) SetErrors(e interface{}) *Controller {
	h.Payload.Errors = e
	return h
}

func (h *Controller) InternalServerErr(rw *http.ResponseWriter) {
	(*rw).WriteHeader(http.StatusInternalServerError)
	h.Render(rw)
}

func (h *Controller) BadRequest(rw *http.ResponseWriter) {
	(*rw).WriteHeader(http.StatusBadRequest)
	h.Render(rw)
}

func (h *Controller) RenderWithStatus(rw *http.ResponseWriter, statusCode int) {
	(*rw).WriteHeader(statusCode)
	h.Render(rw)
}

func (h *Controller) Render(rw *http.ResponseWriter) error {
	data, err := json.Marshal(h.Payload)
	if err != nil {
		return err
	}
	(*rw).Write(data)
	return err
}

func (h *Controller) Response200(rw *http.ResponseWriter) {
	(*rw).WriteHeader(http.StatusOK)
	h.Render(rw)
}

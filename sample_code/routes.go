// File: routes.go
// Description: This file defines the routes for the application, including HTTP methods, URLs, handlers
// @Author: Faisal
//

package main

import "net/http"

type Route struct {
	Method      string
	Url         string
	Handler     HttpHandler
	Middlewares []Middleware
}

var AllRoutes = []Route{
	{
		Method:      http.MethodGet,
		Url:         "/home",
		Handler:     NewHomeController().Get,
		Middlewares: []Middleware{SampleMiddleware},
	},
	{
		Method:      http.MethodPost,
		Url:         "/auth/login/:id",
		Handler:     NewHomeController().Post,
		Middlewares: []Middleware{SampleMiddleware},
	},
	{
		Method:      http.MethodPut,
		Url:         "/auth/register/:id/create",
		Handler:     NewHomeController().Update,
		Middlewares: []Middleware{SampleMiddleware},
	},
	{
		Method:      http.MethodDelete,
		Url:         "/tags/:id/modify",
		Handler:     NewHomeController().Delete,
		Middlewares: []Middleware{SampleMiddleware},
	},
}

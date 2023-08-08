package router

import (
	"net/http"

	"github.com/justinpjose/cushon-assignment/internal/handlers"
)

type Router interface {
	// HandlerFunc allows the usage of an http.HandlerFunc as a request handle
	HandlerFunc(method string, path string, handler handlers.Handler)

	// ServeHTTP allows the router to implement the http.Handler
	ServeHTTP(http.ResponseWriter, *http.Request)
}

// Copyright 2014 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	// If you add new routers please:
	// - Keep the benchmark functions etc. alphabetically sorted
	// - Make a pull request (without benchmark results) at
	//   https://github.com/julienschmidt/go-http-routing-benchmark

	"github.com/meilihao/water"
)

type route struct {
	method string
	path   string
}

type mockResponseWriter struct{}

func (m *mockResponseWriter) Header() (h http.Header) {
	return http.Header{}
}

func (m *mockResponseWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (m *mockResponseWriter) WriteString(s string) (n int, err error) {
	return len(s), nil
}

func (m *mockResponseWriter) WriteHeader(int) {}

var nullLogger *log.Logger

// flag indicating if the normal or the test handler should be loaded
var loadTestHandler = false

func init() {
	// makes logging 'webscale' (ignores them)
	log.SetOutput(new(mockResponseWriter))
	nullLogger = log.New(new(mockResponseWriter), "", 0)
}

// water
func waterHandler(ctx *water.Context) {}

func waterHandlerWrite(ctx *water.Context) {
	ctx.WriteString(ctx.Params.String("name"))
}

func waterHandlerTest(ctx *water.Context) {
	ctx.WriteString(ctx.Req.RequestURI)
}

func loadWater(routes []route) http.Handler {
	h := waterHandler
	if loadTestHandler {
		h = waterHandlerTest
	}

	water.LogClose = true
	router := water.Classic()
	for _, route := range routes {
		switch route.method {
		case "GET":
			router.Get(route.path, h)
		case "POST":
			router.Post(route.path, h)
		case "PUT":
			router.Put(route.path, h)
		case "PATCH":
			router.Patch(route.path, h)
		case "DELETE":
			router.Delete(route.path, h)
		default:
			panic("Unknow HTTP method: " + route.method)
		}
	}
	return router
}

func loadWaterSingle(method, path string, handler water.HandlerFunc) http.Handler {
	water.LogClose = true
	router := water.Classic()
	switch method {
	case "GET":
		router.Get(path, handler)
	case "POST":
		router.Post(path, handler)
	case "PUT":
		router.Put(path, handler)
	case "PATCH":
		router.Patch(path, handler)
	case "DELETE":
		router.Delete(path, handler)
	default:
		panic("Unknow HTTP method: " + method)
	}
	return router
}

// Usage notice
func main() {
	fmt.Println("Usage: go test -bench=. -timeout=20m")
	os.Exit(1)
}

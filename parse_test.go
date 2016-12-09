// Copyright 2013 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package main

import (
	"net/http"
	"testing"
)

// Parse
// https://parse.com/docs/rest#summary
var parseAPI = []route{
	// Objects
	{"POST", "/1/classes/<className>"},
	{"GET", "/1/classes/<className>/<objectId>"},
	{"PUT", "/1/classes/<className>/<objectId>"},
	{"GET", "/1/classes/<className>"},
	{"DELETE", "/1/classes/<className>/<objectId>"},

	// Users
	{"POST", "/1/users"},
	{"GET", "/1/login"},
	{"GET", "/1/users/<objectId>"},
	{"PUT", "/1/users/<objectId>"},
	{"GET", "/1/users"},
	{"DELETE", "/1/users/<objectId>"},
	{"POST", "/1/requestPasswordReset"},

	// Roles
	{"POST", "/1/roles"},
	{"GET", "/1/roles/<objectId>"},
	{"PUT", "/1/roles/<objectId>"},
	{"GET", "/1/roles"},
	{"DELETE", "/1/roles/<objectId>"},

	// Files
	{"POST", "/1/files/<fileName>"},

	// Analytics
	{"POST", "/1/events/<eventName>"},

	// Push Notifications
	{"POST", "/1/push"},

	// Installations
	{"POST", "/1/installations"},
	{"GET", "/1/installations/<objectId>"},
	{"PUT", "/1/installations/<objectId>"},
	{"GET", "/1/installations"},
	{"DELETE", "/1/installations/<objectId>"},

	// Cloud Functions
	{"POST", "/1/functions"},
}

var (
	parseWater http.Handler
)

func init() {
	println("#ParseAPI Routes:", len(parseAPI))

	calcMem("Water", func() {
		parseWater = loadWater(parseAPI)
	})

	println()
}

// Static
func BenchmarkWater_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseWater, req)
}

// One Param
func BenchmarkWater_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseWater, req)
}

// Two Params
func BenchmarkWater_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseWater, req)
}

// All Routes
func BenchmarkWater_ParseAll(b *testing.B) {
	benchRoutes(b, parseWater, parseAPI)
}

// Copyright 2013 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package main

import (
	"net/http"
	"testing"
)

// Google+
// https://developers.google.com/+/api/latest/
// (in reality this is just a subset of a much larger API)
var gplusAPI = []route{
	// People
	{"GET", "/people/<userId>"},
	{"GET", "/people"},
	{"GET", "/activities/<activityId>/people/<collection>"},
	{"GET", "/people/<userId>/people/<collection>"},
	{"GET", "/people/<userId>/openIdConnect"},

	// Activities
	{"GET", "/people/<userId>/activities/<collection>"},
	{"GET", "/activities/<activityId>"},
	{"GET", "/activities"},

	// Comments
	{"GET", "/activities/<activityId>/comments"},
	{"GET", "/comments/<commentId>"},

	// Moments
	{"POST", "/people/<userId>/moments/<collection>"},
	{"GET", "/people/<userId>/moments/<collection>"},
	{"DELETE", "/moments/<id>"},
}

var (
	gplusWater http.Handler
)

func init() {
	println("#GPlusAPI Routes:", len(gplusAPI))

	calcMem("Water", func() {
		gplusWater = loadWater(gplusAPI)
	})

	println()
}

// Static
func BenchmarkWater_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusWater, req)
}

// One Param
func BenchmarkWater_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusWater, req)
}

// Two Params
func BenchmarkWater_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusWater, req)
}

// All Routes
func BenchmarkWater_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusWater, gplusAPI)
}

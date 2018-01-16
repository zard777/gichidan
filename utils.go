// Copyright 2017 hIMEI

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// Output colorizing
const (
	RED   string = "\x1B[31m"
	GRN          = "\x1B[32m"
	YEL          = "\x1B[33m"
	BLU          = "\x1B[34m"
	MAG          = "\x1B[35m"
	CYN          = "\x1B[36m"
	WHT          = "\x1B[97m"
	RESET        = "\x1B[0m"
	BOLD         = "\x1B[1m"
	LINE         = "\x1B[4m"
	INV          = "\x1B[7m"
	ITAL         = "\x1B[3m"
)

// Connect to search engine
const (
	ICHIDAN string = "ichidanv34wrx7m7.onion:80"
	SEARCH         = "/search?query="
)

// Html parsing
const (
	ADDED          string = "Added on "
	LONGFORM              = "2017-09-09 01:30:35 UTC"
	PRE                   = "//pre"
	SPAN                  = "//span"
	LINK                  = "//a"
	HREF                  = "href"
	H2                    = "//h2"
	H3                    = "//h3"
	VERSION               = "//small"
	NONE                  = " "
	CURRENT               = "//em[@class='current']"
	DISABLED              = "//span[@class='next_page disabled']"
	SEARCHRESULT          = "//div[@class='search-results']"
	PAGINATION            = "//div[@class='pagination']"
	DETAILS               = "//a[@class='details']"
	SUMMARY               = "//div[@class='search-result-summary col-xs-4']"
	ONION                 = "//div[@class='onion']"
	TOTALR                = "//div[@class='bignumber']"
	SERVICE               = "//div[@class='service']"
	SERVICES              = "//div[@class='services']"
	SERVICELONG           = "//li[@class='service service-long']"
	SERVICEDETAILS        = "//div[@class='service-details col-sm-2']"
	HOST                  = "//div[@class='search-result row-fluid']"
	NORESULT              = "//div[@class='msg alert alert-info']"
	RESULT                = "//div[@class='col-sm-9']"
	PORT                  = "//div[@class='port']"
	PROTO                 = "//div[@class='protocol']"
	STATE                 = "//div[@class='state']"
	PREVIOUS              = "← Previous"
	NEXT                  = "Next →"
)

var SLEEPER = sleeper()

// Toint converts string to int and handle errors
func toInt(str string) int {
	intCount, err := strconv.Atoi(str)
	ErrFatal(err)

	return intCount
}

// ErrFatal is the basic errors handler
func ErrFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// trimString trims trailing and leading spaces from string
func trimString(str string) string {
	return strings.TrimSpace(str)
}

// Sleeper is a closure which calls time.Sleep with random time
// range between 300 and 359 milliseconds. It will be used in spider.Crawl() to avoid
// highly load to Ichidan server.
func sleeper() func() {
	return func() {
		s := rand.NewSource(time.Now().UnixNano())
		r := rand.New(s)
		p := time.Duration(300 + r.Intn(59))
		time.Sleep(p * time.Millisecond)
	}
}

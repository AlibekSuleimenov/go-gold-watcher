package main

import (
	"net/http"
	"os"
	"testing"
)

var testApp Config

func TestMain(m *testing.M) {

	os.Exit(m.Run())
}

var jsonToReturn = `
{
  "ts": 1709973861292,
  "tsj": 1709973855557,
  "date": "Mar 9th 2024, 03:44:15 am NY",
  "items": [
    {
      "curr": "USD",
      "xauPrice": 2178.795,
      "xagPrice": 24.31,
      "chgXau": 17.085,
      "chgXag": -0.0725,
      "pcXau": 0.7903,
      "pcXag": -0.2973,
      "xauClose": 2161.71,
      "xagClose": 24.3825
    }
  ]
}
`

type RoundTripFunc func(request *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(request *http.Request) (*http.Response, error) {
	return f(request), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{Transport: fn}
}

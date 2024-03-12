package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// currency represents the currency in which the gold price is fetched.
var currency = "KZT"

// Gold represents the structure of the gold prices fetched from the API.
type Gold struct {
	Prices []Price `json:"items"`
	Client *http.Client
}

// Price represents an individual gold price item.
type Price struct {
	Currency      string    `json:"currency"`
	Price         float64   `json:"xauPrice"`
	Change        float64   `json:"chgXau"`
	PreviousClose float64   `json:"xauClose"`
	Time          time.Time `json:"-"`
}

// GetPrices fetches the current gold prices from the API.
func (g *Gold) GetPrices() (*Price, error) {
	// If client is nil, create a new HTTP client.
	if g.Client == nil {
		g.Client = &http.Client{}
	}

	client := g.Client
	url := fmt.Sprintf("https://data-asg.goldprice.org/dbXRates/%s", currency)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("error creating request to goldprice.org", err)
		return nil, err
	}
	response, err := client.Do(request)
	if err != nil {
		log.Println("error contacting goldprice.org", err)
		return nil, err
	}
	defer response.Body.Close()

	// Read the response body.
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("error reading JSON from goldprice.org", err)
		return nil, err
	}

	// Unmarshal the JSON response into the Gold struct.
	gold := Gold{}
	var previous, current, change float64
	err = json.Unmarshal(body, &gold)
	if err != nil {
		log.Println("error unmarshalling JSON from goldprice.org", err)
		return nil, err
	}

	// Extract current gold price information.
	previous, current, change = gold.Prices[0].PreviousClose, gold.Prices[0].Price, gold.Prices[0].Change
	var currentInfo = Price{
		Currency:      currency,
		Price:         current,
		Change:        change,
		PreviousClose: previous,
		Time:          time.Now(),
	}

	return &currentInfo, nil
}

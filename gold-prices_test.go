package main

import (
	"testing"
)

func TestGold_GetPrices(t *testing.T) {
	g := Gold{
		Prices: nil,
		Client: client,
	}

	p, err := g.GetPrices()
	if err != nil {
		t.Error("failed to get prices:", err)
	}

	if p.Price != 2178.795 {
		t.Error("wrong price returned:", p.Price)
	}
}

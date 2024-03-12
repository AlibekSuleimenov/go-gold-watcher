package main

import (
	"fyne.io/fyne/v2/test"
	"testing"
)

func TestApp_getToolbar(t *testing.T) {
	tb := testApp.getToolbar()

	if len(tb.Items) != 4 {
		t.Error("wrong number of items in toolbar")
	}
}

func TestApp_addHoldingsDialog(t *testing.T) {
	testApp.addHoldingsDialog()

	test.Type(testApp.AddHoldingsPurchaseAmountEntry, "1")
	test.Type(testApp.AddHoldingsPurchasePriceEntry, "1000")
	test.Type(testApp.AddHoldingsPurchaseDateEntry, "2024-03-12")

	if testApp.AddHoldingsPurchaseDateEntry.Text != "2024-03-12" {
		t.Error("date is not correct; expected 2024-03-12, but got:", testApp.AddHoldingsPurchaseDateEntry.Text)
	}

	if testApp.AddHoldingsPurchaseAmountEntry.Text != "1" {
		t.Error("amount is not correct; expected 1, but got:", testApp.AddHoldingsPurchaseAmountEntry.Text)
	}

	if testApp.AddHoldingsPurchasePriceEntry.Text != "1000" {
		t.Error("price is not correct; expected 1000, but got:", testApp.AddHoldingsPurchasePriceEntry.Text)
	}
}

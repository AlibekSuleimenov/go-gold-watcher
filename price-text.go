package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"image/color"
)

// getPriceText retrieves the latest gold price information and creates canvas.Text objects
// to display the open price, current price, and price change. If an error occurs while
// fetching the price information, it creates canvas.Text objects with "Unreachable" text
// and sets their color to grey. Otherwise, it formats the price information and creates
// canvas.Text objects with the formatted text and appropriate colors based on the price
// change. The canvas.Text objects are aligned for display and returned to the caller.
func (app *Config) getPriceText() (*canvas.Text, *canvas.Text, *canvas.Text) {
	var g Gold
	var open, current, change *canvas.Text
	g.Client = app.HttpClient

	gold, err := g.GetPrices()
	if err != nil {
		grey := color.NRGBA{
			R: 155,
			G: 155,
			B: 155,
			A: 255,
		}
		open = canvas.NewText("Open: Unreachable", grey)
		current = canvas.NewText("Current: Unreachable", grey)
		change = canvas.NewText("Change: Unreachable", grey)
	} else {
		displayColor := color.NRGBA{
			R: 0,
			G: 180,
			B: 0,
			A: 255,
		}

		if gold.Price < gold.PreviousClose {
			displayColor = color.NRGBA{
				R: 180,
				G: 0,
				B: 0,
				A: 255,
			}
		}

		openText := fmt.Sprintf("Open $%.4f %s", gold.PreviousClose, currency)
		currentText := fmt.Sprintf("Current $%.4f %s", gold.Price, currency)
		changeText := fmt.Sprintf("Change $%.4f %s", gold.Change, currency)

		open = canvas.NewText(openText, nil)
		current = canvas.NewText(currentText, displayColor)
		change = canvas.NewText(changeText, displayColor)
	}

	open.Alignment = fyne.TextAlignLeading
	current.Alignment = fyne.TextAlignCenter
	change.Alignment = fyne.TextAlignTrailing

	return open, current, change
}

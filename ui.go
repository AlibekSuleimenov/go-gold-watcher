package main

import "fyne.io/fyne/v2/container"

// makeUI creates the user interface for displaying the current price of gold.
func (app *Config) makeUI() {
	// get the current price of gold
	openPrice, currentPrice, priceChange := app.getPriceText()

	// put info into a container
	priceContent := container.NewGridWithColumns(3,
		openPrice,
		currentPrice,
		priceChange,
	)
	app.PriceContainer = priceContent

	// get toolbar
	toolbar := app.getToolbar(app.MainWindow)

	// add container to window
	finalContent := container.NewVBox(priceContent, toolbar)
	app.MainWindow.SetContent(finalContent)
}

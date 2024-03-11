package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"time"
)

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
	toolbar := app.getToolbar()
	app.Toolbar = toolbar

	priceTabContent := app.pricesTab()
	holdingsTabContent := app.holdingsTab()

	// get app tabs
	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Prices", theme.HomeIcon(), priceTabContent),
		container.NewTabItemWithIcon("Holdings", theme.InfoIcon(), holdingsTabContent),
	)
	tabs.SetTabLocation(container.TabLocationTop)

	// add container to window
	finalContent := container.NewVBox(priceContent, toolbar, tabs)
	app.MainWindow.SetContent(finalContent)

	// go routine to refresh the prices in background
	go func() {
		for range time.Tick(time.Second * 5) {
			app.refreshPriceContent()
		}
	}()
}

// refreshPriceContent updates the price and chart content displayed in the application
func (app *Config) refreshPriceContent() {
	open, current, change := app.getPriceText()
	app.PriceContainer.Objects = []fyne.CanvasObject{open, current, change}
	app.PriceContainer.Refresh()

	chart := app.getChart()
	app.PriceChartContainer.Objects = []fyne.CanvasObject{chart}
	app.PriceChartContainer.Refresh()
}

// refreshHoldingsTable refreshes the holdings table with updated data.
func (app *Config) refreshHoldingsTable() {
	app.Holdings = app.getHoldingSlice()
	app.HoldingsTable.Refresh()
}

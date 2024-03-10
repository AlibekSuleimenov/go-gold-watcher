package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"log"
	"net/http"
	"os"
)

// Config struct holds configuration data for the application
type Config struct {
	App                 fyne.App        // App holds the Fyne application instance.
	InfoLog             *log.Logger     // InfoLog is the logger for informational messages.
	ErrorLog            *log.Logger     // ErrorLog is the logger for error messages.
	MainWindow          fyne.Window     // MainWindow is the main application window.
	PriceContainer      *fyne.Container // PriceContainer holds the container for displaying price information.
	Toolbar             *widget.Toolbar // Toolbar holds the toolbar for the application.
	PriceChartContainer *fyne.Container // PriceChartContainer holds the container for displaying the price chart.
	HttpClient          *http.Client    // HttpClient is the HTTP client used for making API requests.
}

var myApp Config

// main is the entry point of the application
func main() {
	// create fyne app
	fyneApp := app.NewWithID("go-gold-watcher")
	myApp.App = fyneApp
	myApp.HttpClient = &http.Client{}

	// create loggers
	myApp.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	myApp.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// open a connection to db

	// create a database repository

	// create and size a fyne window
	myApp.MainWindow = fyneApp.NewWindow("GoGoldWatcher")
	myApp.MainWindow.Resize(fyne.NewSize(770, 410))
	myApp.MainWindow.SetFixedSize(true)
	myApp.MainWindow.SetMaster()

	myApp.makeUI()

	// show and run app
	myApp.MainWindow.ShowAndRun()
}

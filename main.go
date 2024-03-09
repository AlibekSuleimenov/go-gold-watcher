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
	App            fyne.App
	InfoLog        *log.Logger
	ErrorLog       *log.Logger
	MainWindow     fyne.Window
	PriceContainer *fyne.Container
	Toolbar        *widget.Toolbar
	HttpClient     *http.Client
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

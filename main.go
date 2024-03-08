package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"log"
	"os"
)

// Config struct holds configuration data for the application
type Config struct {
	App      fyne.App
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

var myApp Config

// main is the entry point of the application
func main() {
	// create fyne app
	fyneApp := app.NewWithID("go-gold-watcher")
	myApp.App = fyneApp

	// create loggers
	myApp.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	myApp.ErrorLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile)

	// open a connection to db

	// create a database repository

	// create and size a fyne window
	window := fyneApp.NewWindow("GoGoldWatcher")

	// show and run app
	window.ShowAndRun()
}

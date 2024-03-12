package main

import (
	"database/sql"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"github.com/alibeksuleimenov/go-gold-watcher/repository"
	_ "github.com/glebarez/go-sqlite"
	"log"
	"net/http"
	"os"
)

// Config struct holds configuration data for the application
type Config struct {
	App                            fyne.App              // App holds the Fyne application instance.
	InfoLog                        *log.Logger           // InfoLog is the logger for informational messages.
	ErrorLog                       *log.Logger           // ErrorLog is the logger for error messages.
	DB                             repository.Repository // DB is the database repository for interacting with the database.
	MainWindow                     fyne.Window           // MainWindow is the main application window.
	PriceContainer                 *fyne.Container       // PriceContainer holds the container for displaying price information.
	Toolbar                        *widget.Toolbar       // Toolbar holds the toolbar for the application.
	PriceChartContainer            *fyne.Container       // PriceChartContainer holds the container for displaying the price chart.
	Holdings                       [][]interface{}       // Holdings holds the data for displaying holdings information.
	HoldingsTable                  *widget.Table         // HoldingsTable is the table widget for displaying holdings information.
	HttpClient                     *http.Client          // HttpClient is the HTTP client used for making API requests.
	AddHoldingsPurchaseAmountEntry *widget.Entry         // AddHoldingsPurchaseAmountEntry is the widget entry for entering purchase amount.
	AddHoldingsPurchaseDateEntry   *widget.Entry         // AddHoldingsPurchaseDateEntry is the widget entry for entering purchase date.
	AddHoldingsPurchasePriceEntry  *widget.Entry         // AddHoldingsPurchasePriceEntry is the widget entry for entering purchase price.

}

// main is the entry point of the application
func main() {
	var myApp Config

	// create fyne app
	fyneApp := app.NewWithID("go-gold-watcher")
	myApp.App = fyneApp
	myApp.HttpClient = &http.Client{}

	// create loggers
	myApp.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	myApp.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// open a connection to db
	sqlDB, err := myApp.ConnectSQL()
	if err != nil {
		log.Panic(err)
	}

	// create a database repository
	myApp.setupDB(sqlDB)

	currency = fyneApp.Preferences().StringWithFallback("currency", "USD")

	// create and size a fyne window
	myApp.MainWindow = fyneApp.NewWindow("GoGoldWatcher")
	myApp.MainWindow.Resize(fyne.NewSize(770, 410))
	myApp.MainWindow.SetFixedSize(true)
	myApp.MainWindow.SetMaster()

	myApp.makeUI()

	// show and run app
	myApp.MainWindow.ShowAndRun()
}

// ConnectSQL establishes a connection to the SQL database and returns a pointer to the database connection.
func (app *Config) ConnectSQL() (*sql.DB, error) {
	path := ""

	if os.Getenv("DB_PATH") != "" {
		path = os.Getenv("DB_PATH")
	} else {
		path = app.App.Storage().RootURI().Path() + "/sql.db"
		app.InfoLog.Println("DB in:", path)
	}

	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// setupDB initializes the database repository with the given SQL database instance.
// It migrates the database schema if necessary and logs any errors.
func (app *Config) setupDB(sqlDB *sql.DB) {
	app.DB = repository.NewSQLiteRepository(sqlDB)

	err := app.DB.Migrate()
	if err != nil {
		app.ErrorLog.Println(err)
		log.Panic()
	}
}

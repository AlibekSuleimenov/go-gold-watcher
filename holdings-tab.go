package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/alibeksuleimenov/go-gold-watcher/repository"
	"strconv"
)

// holdingsTab returns a container containing the holdings table.
func (app *Config) holdingsTab() *fyne.Container {
	app.HoldingsTable = app.getHoldingsTable()

	holdingsContainer := container.NewVBox(app.HoldingsTable)

	return holdingsContainer
}

// getHoldingsTable creates a widget.Table containing holdings data.
// It formats the holdings data into a slice of slices of interface{} using getHoldingSlice().
// Each row in the table corresponds to a holding, with the last cell containing a delete button.
func (app *Config) getHoldingsTable() *widget.Table {
	data := app.getHoldingSlice()
	app.Holdings = data

	table := widget.NewTable(
		func() (int, int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {
			ctr := container.NewVBox(widget.NewLabel(""))
			return ctr
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			if i.Col == (len(data[0])-1) && i.Row != 0 {
				// last cell - put in a button
				w := widget.NewButtonWithIcon("Delete", theme.DeleteIcon(), func() {
					dialog.ShowConfirm("Delete?", "", func(deleted bool) {
						id, _ := strconv.Atoi(data[i.Row][0].(string))
						err := app.DB.DeleteHolding(int64(id))
						if err != nil {
							app.ErrorLog.Println(err)
						}
						app.refreshHoldingsTable()
					}, app.MainWindow)
				})
				w.Importance = widget.HighImportance

				o.(*fyne.Container).Objects = []fyne.CanvasObject{
					w,
				}
			} else {
				// text info
				o.(*fyne.Container).Objects = []fyne.CanvasObject{
					widget.NewLabel(data[i.Row][i.Col].(string)),
				}
			}
		})

	colWidth := []float32{50, 200, 200, 200, 100}
	for i := 0; i < len(colWidth); i++ {
		table.SetColumnWidth(i, colWidth[i])
	}

	return table
}

// getHoldingSlice retrieves holdings from the database and formats them into a slice of slices of interface{}.
// Each inner slice represents a row of data, with the first row containing headers.
func (app *Config) getHoldingSlice() [][]interface{} {
	var slice [][]interface{}

	holdings, err := app.currentHoldings()
	if err != nil {
		app.ErrorLog.Println(err)
	}

	slice = append(slice, []interface{}{"ID", "Amount", "Price", "Date", "Delete"})

	for _, x := range holdings {
		var currentRow []interface{}

		currentRow = append(currentRow, strconv.FormatInt(x.ID, 10))
		currentRow = append(currentRow, fmt.Sprintf("%d toz", x.Amount))
		currentRow = append(currentRow, fmt.Sprintf("$%2f", float32(x.PurchasePrice/100)))
		currentRow = append(currentRow, x.PurchaseDate.Format("2006-01-02"))
		currentRow = append(currentRow, widget.NewButton("Delete", func() {}))

		slice = append(slice, currentRow)
	}

	return slice
}

// currentHoldings retrieves all holdings from the database and returns them.
func (app *Config) currentHoldings() ([]repository.Holdings, error) {
	holdings, err := app.DB.AllHoldings()
	if err != nil {
		app.ErrorLog.Println(err)
		return nil, err
	}

	return holdings, nil
}

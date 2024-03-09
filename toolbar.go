package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// getToolbar creates and returns a new toolbar widget
func (app *Config) getToolbar(window fyne.Window) *widget.Toolbar {
	toolbar := widget.NewToolbar(
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {

		}),
		widget.NewToolbarAction(theme.ViewRefreshIcon(), func() {

		}),
		widget.NewToolbarAction(theme.SettingsIcon(), func() {

		}),
	)

	return toolbar
}

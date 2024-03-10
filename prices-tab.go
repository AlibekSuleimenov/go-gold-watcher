package main

import (
	"bytes"
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"image"
	"image/png"
	"io"
	"os"
	"strings"
)

// pricesTab returns a container containing the price chart.
func (app *Config) pricesTab() *fyne.Container {
	// Retrieve the price chart.
	chart := app.getChart()
	// Create a vertical box container for the chart.
	chartContainer := container.NewVBox(chart)
	// Set the PriceChartContainer field of the app to the chart container.
	app.PriceChartContainer = chartContainer
	// Return the chart container.
	return chartContainer
}

// getChart retrieves the price chart image from API
func (app *Config) getChart() *canvas.Image {
	// Construct the API URL for the price chart image based on the selected currency.
	apiURL := fmt.Sprintf("https://goldprice.org/charts/gold_3d_b_o_%s_x.png", strings.ToLower(currency))
	var img *canvas.Image

	// Download the chart image file from the API URL.
	err := app.downloadFile(apiURL, "gold.png")
	if err != nil {
		// Use a bundled image if downloading fails.
		img = canvas.NewImageFromResource(resourceUnreachablePng)
	} else {
		// Create a new image from the downloaded file.
		img = canvas.NewImageFromFile("gold.png")
	}

	// Set the minimum size of the image.
	img.SetMinSize(fyne.Size{Width: 770, Height: 410})
	// Set the fill mode of the image to fill original size.
	img.FillMode = canvas.ImageFillOriginal

	// return the image
	return img
}

// downloadFile downloads a file from the specified URL and saves it with the given fileName.
func (app *Config) downloadFile(URL, fileName string) error {
	// get the response bytes from calling a url
	response, err := app.HttpClient.Get(URL)
	if err != nil {
		return err
	}

	// Check if the response status code is valid.
	if response.StatusCode != 200 {
		return errors.New("received wrong response code when downloading image")
	}

	// Read the response body into a byte slice.
	b, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Decode the byte slice into an image.
	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		return err
	}

	// Create the output file.
	out, err := os.Create(fmt.Sprintf("./%s", fileName))
	if err != nil {
		return err
	}

	// Encode the image as a PNG and write it to the output file.
	err = png.Encode(out, img)
	if err != nil {
		return err
	}

	return nil
}

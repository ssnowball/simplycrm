package main

import (
	"github.com/kbinani/screenshot"
	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

func main() {

	var width int
	var height int

	n := screenshot.NumActiveDisplays()
	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)

		width = bounds.Dx()
		height = bounds.Dy()

		// fmt.Printf("Display #%d resolution is %d x %d\n", i, width, height)
	}

	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:  width,
		Height: height,
		// Width:  1500,
		// Height: 850,
		Title:  "simplycrm",
		JS:     js,
		CSS:    css,
		Colour: "#FFFFFF",
	})
	app.Bind(readData)
	app.Bind(createData)
	app.Bind(deleteData)
	app.Bind(getUpdateData)
	app.Bind(updateData)
	app.Bind(readSubData)
	app.Bind(databaseBuild)
	app.Run()
}

package main

import (
	"fmt"
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"
)

func main() {
	go func() {

		w := new(app.Window)
		w.Option(app.Title("Grid example - A grid of widgets"))
		w.Option(app.Size(unit.Dp(810), unit.Dp(810)))

		if err := draw(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

const sideLength int = 8
const cellSize int = 90

type (
	C = layout.Context
	D = layout.Dimensions
)

func draw(w *app.Window) error {
	th := material.NewTheme()
	var (
		ops  op.Ops
		grid component.GridState
	)

	clickers := []widget.Clickable{}
	for i := 0; i < sideLength*sideLength; i++ {
		clickers = append(clickers, widget.Clickable{})
	}

	for {

		windowevent := w.Event()
		switch e := windowevent.(type) {
		case app.DestroyEvent:
			return e.Err

		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			for i := range clickers {
				if clickers[i].Clicked(gtx) {
					fmt.Println("You clicked button", i)
				}
			}

			component.Grid(th, &grid).Layout(gtx, sideLength, sideLength,
				func(axis layout.Axis, index, constraint int) int {
					return gtx.Dp(unit.Dp(cellSize))
				},
				func(gtx C, row, col int) D {
					clk := &clickers[row*sideLength+col]
					btn := material.Button(th, clk, fmt.Sprintf("R%d C%d", row, col))
					color := color.NRGBA{
						R: uint8(255 / sideLength * row),
						G: uint8(255 / sideLength * col),
						B: uint8(255 * row * col / (sideLength * sideLength)),
						A: 255,
					}
					btn.Background = color
					btn.CornerRadius = 0
					return btn.Layout(gtx)
				})

			e.Frame(gtx.Ops)
		}
	}
}

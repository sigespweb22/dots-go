package ui

import (
	"dots-go/internal/app/engine"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Ligando os Pontos",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	e := engine.NewEngine()

	e.DrawLine(pixel.V(100, 100), pixel.V(300, 300))

	for !win.Closed() {
		e.Draw()
		win.Update()
	}
}

func StartWindow() {
	pixelgl.Run(run)
}

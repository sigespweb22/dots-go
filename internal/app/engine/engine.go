package engine

import (
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

type Engine struct {
	canvas *imdraw.IMDraw
}

func NewEngine() *Engine {
	return &Engine{
		canvas: imdraw.New(nil),
	}
}

func (e *Engine) DrawLine(from pixel.Vec, to pixel.Vec) {
	e.canvas.Color = color.Black
	e.canvas.Push(from, to)
	e.canvas.Line(5)
}

func (e *Engine) Draw() {
	e.canvas.Clear()

	// TODO: Adicionar código para desenhar as linhas do jogo

	e.canvas.Draw(pixel.IM)
}

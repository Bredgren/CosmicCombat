package main

import (
	"log"

	"github.com/Bredgren/gogame"
)

func main() {
	log.SetPrefix("")
	log.SetFlags(0)

	ready := gogame.Ready()
	go func() {
		<-ready
		initGG()
		start()
	}()
}

func initGG() {
	width, height := 900, 600
	canvas := gogame.GetDisplay()
	canvas.SetMode(width, height)
	canvas.Fill(&gogame.FillStyle{Colorer: gogame.Black})
}

func start() {
	log.Println("start")
	display := gogame.GetDisplay()
	display.DrawRect(&gogame.Rect{X: 11, Y: 11, W: 48, H: 48}, &gogame.StrokeStyle{Colorer: gogame.White, Width: 4})
	display.DrawRect(&gogame.Rect{X: 70, Y: 10, W: 50, H: 50},
		&gogame.FillStyle{Colorer: &gogame.LinearGradient{
			X1: 70, Y1: 10, X2: 70 + 50, Y2: 10 + 50,
			ColorStops: []gogame.ColorStop{
				{0.0, gogame.Red},
				{1.0, gogame.Blue},
			},
		}})
	display.DrawRect(&gogame.Rect{X: 130, Y: 10, W: 50, H: 50},
		&gogame.FillStyle{Colorer: &gogame.RadialGradient{
			X1: 130 + 50/2, Y1: 10 + 50/2, R1: 40, X2: 130 + 50/4, Y2: 10 + 50/4, R2: 1,
			ColorStops: []gogame.ColorStop{
				{0.0, gogame.Blue},
				{1.0, gogame.Green},
			},
		}})
}

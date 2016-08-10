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
	canvas.Fill(gogame.Black)
}

func start() {
	log.Println("start")
	gogame.GetDisplay().DrawRect(&gogame.Rect{X: 10, Y: 10, W: 50, H: 50}, gogame.White, 0)
	gogame.GetDisplay().DrawRect(&gogame.Rect{X: 70, Y: 11, W: 48, H: 48}, gogame.White, 4)
}

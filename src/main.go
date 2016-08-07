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
	canvas := gogame.GetCanvas()
	canvas.SetWidth(width)
	canvas.SetHeight(height)
	canvas.Fill(gogame.Black)
}

func start() {
	log.Println("start")
	gogame.Log(gogame.GetCanvas())
}

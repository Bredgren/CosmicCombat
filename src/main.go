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
		start()
	}()
}

func start() {
	log.Println("start")
}

package main

import (
	"log"
	"os"
	"os/signal"
	"test/customModule/htmlToPdf"
)

func main() {
	println("test")
	htmlToPdf.ToPdf()
	sigc := make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt)
	select {
	case <-sigc:
	}

	log.Println("exiting...")
}

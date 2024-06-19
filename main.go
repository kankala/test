package main

import (
	"log"
	"os"
	"os/signal"
)

func main() {
	println("test")
	sigc := make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt)
	select {
	case <-sigc:
	}

	log.Println("exiting...")
}

package main

import (
	"log"
	"os"
	"os/signal"
	"searchservice/cmd"
	"searchservice/config"
)

func main() {
	config.ReadConfig()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	go func() {
		if err := cmd.Execute(); err != nil {
			log.Fatal(err)
		}
	}()
	sig := <-c
	log.Println("GOT SIGNAL: " + sig.String())
}

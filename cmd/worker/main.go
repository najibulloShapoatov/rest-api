package main

import (
	"business/cmd/worker/app"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	name := flag.String("name", "Worker", "name to print")
	flag.Parse()
	log.Printf("Starting sleepservice for %s", *name)
	// setup signal catching
	sigs := make(chan os.Signal, 1)
	// catch all signals since not explicitly listing
	// signal.Notify(sigs)
	signal.Notify(sigs, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGINT)
	//signal.Notify(sigs,syscall.SIGQUIT)
	// method invoked upon seeing signal
	go func() {
		s := <-sigs
		log.Printf("RECEIVED SIGNAL: %s", s)

		os.Exit(1)
	}()

	app.Run()
}

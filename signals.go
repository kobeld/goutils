package goutils

import (
	"log"
	"os"
	"os/signal"
)

func CleanupOnInterrupt(jobs ...func()) chan bool {
	var (
		signalChan = make(chan os.Signal, 1)
		doneChan   = make(chan bool)
	)

	signal.Notify(signalChan, os.Interrupt)

	go func() {
		for _ = range signalChan {
			println("")
			log.Println("===== Received an interrupt, handling cleanup services. =====")
			for _, job := range jobs {
				job()
			}
			log.Println("===== All jobs were finished, existing the program. =====")
			doneChan <- true
		}

	}()

	println("Program started, waiting interrupt signal.")
	return doneChan
}

func CleanupAndExitOnInterrupt(jobs ...func()) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	go func() {
		for _ = range signalChan {
			println("")
			log.Println("===== Received an interrupt, handling cleanup services. =====")
			for _, job := range jobs {
				job()
			}
			log.Println("===== All jobs were finished, existing the program. =====")
			os.Exit(0)
		}
	}()
}

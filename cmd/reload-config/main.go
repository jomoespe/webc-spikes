package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGHUP)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		fmt.Println("Waiting for SIGHUB signal (kill -1 <pid>)...")
		for {
			select {
			case s := <-sigc:
				fmt.Printf("received signal %v. Here reoad and restart.\n", s)
			}
		}
	}()
	fmt.Println("Waiting forever...")
	wg.Wait()
}

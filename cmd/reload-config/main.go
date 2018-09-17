// Copyright 2018 Jomoespe. All rights reserved.
// Use of this source code is governed by a WTFPL-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGHUP)

	go func() {
		fmt.Println("Waiting for SIGHUB signal (kill -SIGHUP <pid>)...")
		for {
			select {
			case s := <-sigc:
				fmt.Printf("Received signal %s. Here reload config and continue.\n", s)
			}
		}
	}()
	fmt.Println("Waiting forever...")
	c := make(chan bool)
	<-c
}

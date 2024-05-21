package main

import (
	"fmt"
	"time"
)

func ping(pingCh chan<- bool, pongCh <-chan bool) {
	for {
		<-pongCh
		fmt.Println("ping")
		time.Sleep(500 * time.Millisecond)
		pingCh <- true
	}
}

func pong(pingCh <-chan bool, pongCh chan<- bool) {
	for {
		<-pingCh
		fmt.Println("pong")
		time.Sleep(500 * time.Millisecond)
		pongCh <- true
	}
}

func main() {
	pingCh := make(chan bool)
	pongCh := make(chan bool)

	go ping(pingCh, pongCh)
	go pong(pingCh, pongCh)

	pongCh <- true

	time.Sleep(10 * time.Second)
}

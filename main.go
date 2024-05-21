
package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup

    pingChannel := make(chan string)
    pongChannel := make(chan string)

    wg.Add(2)

    go func() {
        for {
            pingChannel <- "PING"
            wg.Done()
        }
    }()

    go func() {
        for {
            <-pongChannel
            fmt.Println("PING")
            pingChannel <- "PONG"
        }
    }()

    wg.Wait()

    close(pingChannel)
    close(pongChannel)
}

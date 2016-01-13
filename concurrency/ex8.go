package main

import (
	"fmt"
	"runtime"
	"os"
	"time"
	"log"
)

func main() {

    runtime.GOMAXPROCS(4)

    f, _ := os.Create("log.txt")
    f.Close()

    logCh := make(chan string, 50)

    go func() {
        for {
            msg, ok := <-logCh
            if ok {
                f, err := os.OpenFile("log.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
                if err != nil {
                    log.Fatal(err)
                }
                logTime := time.Now().Format(time.RFC3339)
                if _, err := f.WriteString(logTime + " - " + msg); err != nil {
                    log.Fatal(err)
                }
                f.Close()
            } else {
                break
            }
        }
    }()

    mutex := make(chan bool, 1)

    for i := 1; i < 10; i++ {
        for j := 1; j < 10; j++ {
            mutex <- true
            go func() {
                msg := fmt.Sprintf("%d + %d = %d\n", i, j, i + j)
                logCh <- msg
                fmt.Print(msg)
                <-mutex
            }()
        }
    }

    fmt.Scanln()

}

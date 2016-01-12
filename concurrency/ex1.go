package main

import (
	"runtime"
	"time"
)

func main() {

	godur, _ := time.ParseDuration("10ms")
	runtime.GOMAXPROCS(2)
	go func() {
		for i := 0; i < 100; i++ {
			println("hello")
			time.Sleep(godur)
		}

	}()

	go func() {
		for i := 0; i < 100; i++ {
			println("go")
			time.Sleep(godur)
		}
	}()

	dur, _ := time.ParseDuration("1s")
	time.Sleep(dur)
}

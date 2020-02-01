package main

import (
	"fmt"
	"time"
)

func main() {
	testChan := make(chan struct{})

	go func(ch chan struct{}) {
		for {
			select {
				case _, ok := <- ch:
					fmt.Println("!", ok)
					return
			}
		}
	}(testChan)

	go func(ch chan struct{}) {
		for {
			select {
				case _, ok := <- ch:
					fmt.Println("!!", ok)
					return
			}
		}
	}(testChan)

	time.Sleep(1*time.Second)
	testChan <- struct{}{}

	time.Sleep(2*time.Second)
}
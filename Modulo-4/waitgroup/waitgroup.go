package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(4)

	go func() {
		escrever("GoRoutine 1 	| Execução:")
		waitGroup.Done() // -1
	}()

	go func() {
		escrever("GoRoutine 2 	| Execução:")
		waitGroup.Done() // -1
	}()

	go func() {
		escrever("GoRoutine 3 	| Execução:")
		waitGroup.Done() // -1
	}()

	go func() {
		escrever("Goroutine 4 	| Execução:")
		waitGroup.Done() // -1
	}()

	waitGroup.Wait() // 0
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto, i)
		time.Sleep(time.Second / 2)
	}
}

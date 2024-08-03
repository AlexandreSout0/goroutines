package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(2) // Quantide de goRoutines que vão trabalhar em paralelo

	go func() {
		myFuncWrite("Hello Word")
		waitGroup.Done() // -1 do waitGroup
	}()

	go func() {
		myFuncWrite("Learning Go lang!")
		waitGroup.Done() // -1 do waitGroup
	}()

	waitGroup.Wait() // 0 do waitGroup
	//execução termina
}

func myFuncWrite(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}

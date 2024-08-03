package main

import (
	"fmt"
	"time"
)

func main(){
	go myFuncWrite("Hello Word") // GoRoutine
	myFuncWrite("Learning Go lang!")
}


func myFuncWrite(text string){
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
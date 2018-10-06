package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 100; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main(){
	go say("word")
	say("hello")
}

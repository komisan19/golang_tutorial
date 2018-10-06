package main

import "fmt"

func main() {
	ch := make(chan int, 7)
	ch <- 3
	ch <- 4
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

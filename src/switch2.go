package main

import "fmt"

func main() {
	i := 5
	switch i {
	case 4:
		fmt.Println("number 4")
	case 5:
		fmt.Println("number 5")
	case 6:
		fmt.Println("number 6")
	default:
		fmt.Println("default")
	}
}

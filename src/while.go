package main

import "fmt"

func main(){
	x := 0
	for x < 100{
		if x % 15 == 0{
			fmt.Println("fizzbuzz")
		}else if x % 3 == 0{
			fmt.Println("fizz")
		}else if x % 5 == 0{
			fmt.Println("buzz")
		}else{
			fmt.Println(x)
		}
		x++
	}
}

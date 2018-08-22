package main

import "fmt"

for i := 0 ; i < 10000; i++ {
	if i % 15 == 0{
		fmt.Println("fizzbuzz")
	}else if i % 3 == 0{
		fmt.Println("fizz")
	}else if i % 5 == 0{
		fmt.Println("buzz")
	}else{
		fmt.Println("sum")
	}
}

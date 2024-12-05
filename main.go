package main

import (
	"fmt"
	"go-sample/fizzbuzz"
)

func main() {
	for i := 1; i <= 20; i++ {
		fmt.Println(fizzbuzz.FizzBuzz(i))
	}

}

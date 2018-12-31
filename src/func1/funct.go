package main

import (
	"fmt"
	"mymath111"
)

func main() {
	c,err := mymath.Add(1, 3)
	fmt.Println(c)
	if err == nil {
		fmt.Println("No errors")
	}
}
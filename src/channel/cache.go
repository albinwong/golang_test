package main

import "fmt"

func main() {
	c := make(chan int, 1024)

	for i := range c {
		fmt.Println("Received:", i)
		// fmt.Println(k)
	}
}
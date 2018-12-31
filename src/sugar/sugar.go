package main

import (
	"fmt"
)

func myfunc(args ...int) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}

func main() {
	myfunc(2, 3, 4)
	myfunc(1, 3, 7, 13)
}
package main

import (
	"fmt"
)

// 条件语句 if 
/*func test(b int) string{
	var a string
	if b > 5 {
		a =  "true"
	} else {
		a =  "false"
	}
	return a
}

func example(x int) int { 
	if x == 0 {
		return 5 
	} else {
		return x 
	}
}
func main() {
	
	fmt.Println(example(5))
}*/

// 条件语句 switch
func main() {
	i := 5
	switch i { 
		case 0:
			fmt.Printf("0") 
		case 1:
			fmt.Printf("1") 
		case 2:
			fallthrough 
		case 3:
			fmt.Printf("3") 
		case 4, 5, 6:
			if(i == 4) {
				fmt.Printf("4")
			} else {
				fmt.Printf("5 or 6")
			}
			// fmt.Printf("4, 5, 6") 
		default:
	        fmt.Printf("Default")
	}
}


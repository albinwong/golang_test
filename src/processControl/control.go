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
/*func main() {
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
}*/


// 循环语句
func main() {
	// case one
	/*sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Println(sum)
	}*/

	// case two
	/*sum := 0
	i := 0
	for {
		sum += i
		fmt.Println(sum)
		i++
		if i > 10 {
			break
		}
	}*/

	// case three
	/*a := []int{1,2,3,4,5,6,7}
	for i,j := 0,len(a) - 1; i < j; i,j = i + 1, j -1 {
		a[i], a[j] = a[j], a[i]
	}
	fmt.Println(a)*/
	/*for j := 0; j < 5; j++ {
		for i :=0; i <10; i++ {
			if i > 5 {
				continue
			}
	 		fmt.Println(i)
 		}
	}*/

	// go to 
	i := 0
	HERE:
		fmt.Println(i)
		i++
	if i < 10 {
		goto HERE
	}
	fmt.Println("test")

}


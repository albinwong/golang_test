// type system 类型系统
package main

import (
	"fmt"
)
type Integer int

// func (a Integer) Less(b Integer) bool {
// 	res := a < b
// 	return res
// }

// func main() {
// 	var a Integer = 1
// 	if a.Less(2) {
// 		fmt.Println(a, "Less 2")
// 	}


/*func (a *Integer) Add(b Integer) {
	*a += b
	return
}*/

func (a Integer) Add(b Integer) {
	a += b
	return
}

func main() {
	var a Integer = 3
	a.Add(2)
		fmt.Println("a = ",a)
}

/*func main() {
	var v1 interface{} = 1
	fmt.Println(v1)
}*/
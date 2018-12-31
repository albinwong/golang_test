package main

import "fmt"

/*func modify(array [5]int) {
	array[0] = 10 // 试图修改数组的第一个元素 
	fmt.Println("In modify(), array values:", array)
}

func main() {
	array := [5]int{1,2,3,4,5} // 定义并初始化一个数组
	modify(array) // 传递给一个函数，并试图在函数体内修改这个数组内容
    fmt.Println("In main(), array values:", array)
}*/


func main(){
	// 数组创建
	/*mySlice1 := make([]int, 5, 5)
	mySlice1 = append(mySlice1, 1, 5, 3, 3, 2, 1, 10, 1, 3, 5)*/

	// 元素遍历
	/*for i,j := range mySlice1 {
		fmt.Println(i,j)
	}
	for i := 0; i < len(mySlice1) ; i++{
		fmt.Println(i, mySlice1[i])
	}*/

	// mySlice1[10] = 3
    // mySlice3 := []int{"t", "3", "3", "4", "5"}
    /*mySlice2 := make([]int, 5)
    mySlice2[2] = 3
    fmt.Println(mySlice2)*/

    // 基于数组切片，赋值给信的数组
    /*oldSlice := []int{1, 2, 3, 4, 5}
	newSlice := oldSlice[:1]
	fmt.Println(newSlice)*/


	// 数组复制
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	// copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置true
	// fmt.Println(slice1)
	fmt.Println(slice2)
}
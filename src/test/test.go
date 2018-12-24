package main

import "fmt"
import "reflect"

func main() {
	/*fval := 110.48
	ival := 200
	sval := "This is a string"
	fmt.Println("The value of fval is", fval)
	fmt.Printf("fval=%f, ival=%d, sval=%s\n", fval, ival, sval)
	fmt.Printf("fval=%f, ival=%d, sval=%v\n", fval, ival, sval)*/
	const (
		c0 = "w"
		/*c1 = iota
		c2 = 'w'
		c3 = iota*/
	)
	/*const c0 = iota
	const c1 = iota
	const c2 = iota
	const c3 = iota*/
	/**
	16 8 4 2 1
	1 1 1 0 0
	0 0 0 1 0
	 */
	// var a float32 = 1.00000000000002
	// var a = 32+12i
	str := "Hello,世界"
	n := len(str)
	fmt.Println(reflect.TypeOf(n))
	fmt.Println(n)
	for i := 0; i < n; i++ {
		ch := str[i] // 依据下标取字符串中的字符，类型为byte
        fmt.Println(i, ch)
    }
    // arr := [] int
	// fmt.Println(arr)

}
/*
0000 0000 0000 0010 2 

1111 1111 1111 1101 

1000 0000 0000 0011  -3 */
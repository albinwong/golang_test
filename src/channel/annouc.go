package main

import (
	"fmt"
)

/*func main() {
	// channel声明方式
	var chanName chan ElementType // 声明方式
	var a chan int // 声明一个类型为int的channel
	var b map[string] chan bool // 声明一个map, 元素是bool类型
	c := make(chan int) // 声明并初始化一个int类型名为c的channel

	ch <- value // channel写入

	value := <-ch // channel读取
}*/


func main() {
	
	ch := make(chan int, 1) 
	for {
		select {
			case ch <- 0:
			case ch <- 1: 
		}
		i := <-ch
        fmt.Println("Value received:", i)
        if i == 0 {
        	break
        }
    }
}
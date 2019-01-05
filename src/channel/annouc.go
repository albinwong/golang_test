package main

import (
	"fmt"
)

func main() {
	// channel声明方式
	var chanName chan ElementType // 声明方式
	var a chan int // 声明一个类型为int的channel
	var b map[string] chan bool // 声明一个map, 元素是bool类型
	c := make(chan int) // 声明并初始化一个int类型名为c的channel

	ch <- value // channel写入

	value := <-ch // channel读取
}
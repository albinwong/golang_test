package main

import (
	"fmt"
	"sync"
	"runtime"
)

var counter int = 0

func Count(lock *sync.Mutex) {
	lock.Lock() // 加锁
	counter++
	fmt.Println(lock)
	lock.Unlock() // 解锁
}

func main() {
	lock := &sync.Mutex{}

	// 10个goroutine是并发执行的
	for i := 0; i < 10; i++ {
		go Count(lock) // goroutine
	}

	// 用for 循环来不断检查counter的值(同样需要加锁)
	for {
		lock.Lock()

		c := counter

		lock.Unlock()

		runtime.Gosched()

		if c >= 10 {
			break // 当其值达到10时，说明所有goroutine都执行完 毕了，这时主函数返回，程序退出
		}
	}

}
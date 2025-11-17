package stage2

import (
	"fmt"
	"sync"
	"time"
)

func chanMethod1() {
	var channel chan int = make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			channel <- i
		}
		close(channel)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for val := range channel {
			fmt.Printf("%d ", val)
		}
	}()

	time.Sleep(time.Second)
}

func chanMethod2() {
	channel2 := make(chan int, 10)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 100; i++ {
			channel2 <- i
			fmt.Printf("发送: %d, 缓冲区当前长度: %d\n", i, len(channel2))
		}
		close(channel2)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for val := range channel2 {
			fmt.Printf("接收: %d, 缓冲区剩余长度: %d\n", val, len(channel2))
		}
	}()
	wg.Wait()
	fmt.Println("所有任务完成")
}

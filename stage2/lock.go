package stage2

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func lockMethod1() {
	var count int
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for j := 0; j < 1000; j++ {
				mu.Lock()
				count++
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	fmt.Printf("最终计数器值: %d\n", count)
}

func lockMethod2() {
	var count int64
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&count, 1)
			}
		}()
	}

	wg.Wait()
	fmt.Printf("最终计数器值: %d\n", atomic.LoadInt64(&count))
}

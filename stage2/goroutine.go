package stage2

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func ExeGoroutine() {
	var wg sync.WaitGroup
	wg.Add(1)
	go printOldNumber(&wg)
	wg.Add(1)
	go printEvenNumber(&wg)
	wg.Wait()
	fmt.Println("All goroutines finished!")
}

func printEvenNumber(wg *sync.WaitGroup) {
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			randomNum := rand.Intn(100)
			time.Sleep(time.Duration(randomNum) * time.Millisecond)
			fmt.Printf("Even Number : %d \n", i)
		}
	}
}

func printOldNumber(wg *sync.WaitGroup) {
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			randomNum := rand.Intn(100)
			time.Sleep(time.Duration(randomNum) * time.Millisecond)
			fmt.Printf("Old Number : %d \n", i)
		}
	}
}

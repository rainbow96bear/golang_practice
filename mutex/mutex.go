package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	var sharedValue int

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go increment(&sharedValue)
	}

	wg.Wait()
	fmt.Println("Final value:", sharedValue)
}

func increment(value *int) {
	defer wg.Done()
	fmt.Println("Lock의 범위가 아닌 곳은 병렬처리")
	mu.Lock()
	defer mu.Unlock()

	newValue := *value + 1
	fmt.Println("여기는 순차적인 직렬처리")
	time.Sleep(time.Second) // 작업 시뮬레이션
	*value = newValue
}

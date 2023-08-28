package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mu1 sync.Mutex
var mu2 sync.Mutex

func firstFunc(x, y *int){
	defer wg.Done()
	fmt.Println("firstFunc 시작")
	mu1.Lock()
	*x += 1
	fmt.Println("x 증가")
	time.Sleep(time.Second)
	mu2.Lock()
	*y += 1
	fmt.Println("y 증가")
	fmt.Println("firstFunc 결과 : ", *x * *y)
	mu1.Unlock()
	mu2.Unlock()
}

func secondFunc(x, y *int){
	defer wg.Done()
	fmt.Println("secondFunc 시작")
	mu2.Lock()
	*y += 1
	fmt.Println("y 증가")
	time.Sleep(time.Second)
	mu1.Lock()
	*x += 1
	fmt.Println("x 증가")
	mu2.Unlock()
	mu1.Unlock()
}

func main(){
	fmt.Println("시작")
	x := 1
	y := 1

	wg.Add(2)
	go firstFunc(&x, &y)
	go secondFunc(&x, &y)

	wg.Wait()
}
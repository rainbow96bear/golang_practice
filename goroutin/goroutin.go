package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func twoSecondTrigger() {
    defer wg.Done()
    defer fmt.Println("twoSecondTrigger 종료")
    for i:=0 ; i<5 ; i++ {
        time.Sleep(time.Second * 2)
        fmt.Printf("%d번째 2초 트리거\n",i+1)
    }
}

func threeSecondTrigger() {
    defer wg.Done()
    defer fmt.Println("threeSecondTrigger 종료")
    for i:=0 ; i<5 ; i++ {
        time.Sleep(time.Second * 3)
        fmt.Printf("%d번째 3초 트리거\n",i+1)
    }
}

func main() {
    wg.Add(2)
    go twoSecondTrigger()	// 고루틴 작동
    go threeSecondTrigger()	// 고루틴 작동
    wg.Wait()
}
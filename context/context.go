package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker(ctx context.Context) {
	tick := time.Tick(time.Second)
	fmt.Println("parent Data는 : ",ctx.Value("parent Data"))
	fmt.Println("child Data는 : ",ctx.Value("child Data"))
	fmt.Println("child Data는 : ",ctx.Value("child Data"))
	for {
		select {
		case <-ctx.Done() :
			fmt.Println("child context Done")
			wg.Done()
		case <- tick:
			fmt.Println("tick toc")
		}
	}
}

func main() {
	wg.Add(1)
	parentCtx := context.Background()
	parentCtx = context.WithValue(parentCtx, "parent Data", "Parent")
	childCtx := context.WithValue(parentCtx, "child Data", "Child")
	childCtx, _ = context.WithTimeout(childCtx, 5*time.Second)

	go worker(childCtx)
	
	if data := parentCtx.Value("parent Data"); data == "Parent"{
		fmt.Println("Main completed")
	}
	wg.Wait()
}






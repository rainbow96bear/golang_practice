package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
func testFunc(s string){
	for i:=0; i<5 ; i++{
		fmt.Println(s,i+1)
		time.Sleep(100*time.Millisecond)
	}
}
func main(){
	wg.Add(3)
	testFunc("Sync")

	go testFunc("Async1")
	go testFunc("Async2")
	go testFunc("Async3")
	wg.Wait()	
}
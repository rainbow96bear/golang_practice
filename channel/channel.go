package main

import (
	"fmt"
	"sync"
	"time"
)
func readData(wg *sync.WaitGroup, data1 chan string, data2 chan string) {
	for {
            select {
            case getData := <-data1 :
                fmt.Println(getData)
				time.Sleep(time.Second)
            case quit := <-data2:
				if quit == "quit" {
					fmt.Println("data2의 값 : ",quit)
					wg.Done()
					return
				}
			}
	}
}
func main() {
	data1 := make(chan string)
    data2 := make(chan string)
	var wg sync.WaitGroup
	wg.Add(1)

	go readData(&wg, data1, data2)
	for i:=0 ; i<5 ; i++ {
		data1 <- fmt.Sprintf("%d번째 Data",i+1)
	}
	data2 <- "quit"
    
    wg.Wait()
}
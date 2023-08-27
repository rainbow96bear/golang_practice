package main

import (
	"fmt"
)

func nextValue() func() int {

	i := 0
 	return func() int {
 		i++
 		return i
 	}
}

func practiceTest (x,y int) int {
	return x * y
}
func main() {
	var x, y int
	_,err := fmt.Scan(&x, &y)
	if err != nil {
		fmt.Println(err)
	}
	practiceTest(x,y)
}	

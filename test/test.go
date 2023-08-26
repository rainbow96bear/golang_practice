package main

import "fmt"

func nextValue() func() int {

	i := 0
 	return func() int {
 		i++
 		return i
 	}
}

func main() {
	
	next := nextValue()

	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())

	anotherNext := nextValue()
	println(anotherNext()) // 1 다시 시작
	println(anotherNext()) // 2	
}	

package main

import "fmt"
func sum(a, b int) int {
	return a+b
}
func main(){

    if i := sum(5, 2);i < 10 && i > 5 {
    	fmt.Println("i는 10보다 작고 5보다 크다.")
    } else if i > 10 {
    	fmt.Println("i는 10보다 크다.")
    } else {
    	fmt.Println("i는 10입니다.")
    }

	
}
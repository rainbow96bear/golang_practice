package main

import "fmt"

	func main(){
		i:=7
	switch i{
		case 6, 7 :
			fmt.Println("i는 6 또는 7입니다.")
		case 8:
			fallthrough
		case 9:
			fmt.Println("i는 8 또는 9입니다.")
	}
	switch {
	case i <10  :
		fmt.Println("i는 10보다 작다.")
	case i > 10:
		fmt.Println("i는 10보다 크다.")
	default:
		fmt.Println("i는 10입니다.")
	}
}
package main

import "fmt"


func main(){
	checkArr1 := []int{1, 3, 5, 7, 9, 11}
	checkArr2 := append(checkArr1[:3],checkArr1[4:]...)
	fmt.Println(checkArr2)
	fmt.Println(checkArr1)
}
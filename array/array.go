package main

import (
	"fmt"
)

func main(){
	var testArr1 [5]int
	var testArr2 [4]int = [4]int{1,2,3}
	var testArr3 = [6]int{2:10,4:20}
	testArr4 := [...]string{"안녕","반가워","잘가", 5:"또 봐"}
	testArr5 := [...]int{10, 20, 30, 5:50}
	testArr6 := []float64{12.3, 45.6, 8:78.9}
	fmt.Println(testArr1)
	fmt.Println(testArr2)
	fmt.Println(testArr3)
	fmt.Println(testArr4)
	fmt.Println(testArr5)
	fmt.Println(testArr5[3])
	fmt.Println(testArr6)


	checkArr1 := []int{1, 3, 5, 7, 9, 11}
	fmt.Println(checkArr1[4:])
	checkArr1 = append(checkArr1[:3],checkArr1[4:]...)
	fmt.Println(append([]int{11,23},checkArr1[4:]...))
	fmt.Println(checkArr1)

	for i :=0; i<len(testArr4); i++{
		fmt.Println(i,testArr4[i])
	}
	for i,v := range checkArr1{
		fmt.Println(i,v)
	}
}

package main

import "fmt"

func changeValue(testInt *int) int {
    *testInt = *testInt+10
    return *testInt
}

func main() {
    var testNum int = 10
	testNum=changeValue(&testNum)
    fmt.Println(testNum)    
}
package main

import "fmt"

func changeValue(testInt *int) int {
    *testInt = *testInt+10
    return *testInt
}

func main() {

    num := 30
    ptr1 := &num

    fmt.Println(ptr1)
    fmt.Println(*ptr1)

    num = 40
    fmt.Println(ptr1)
    fmt.Println(*ptr1)

    *ptr1 = 50
    fmt.Println(num)
    fmt.Println(&num)
    fmt.Println(ptr1)
    fmt.Println(*ptr1)

    testNum := 15
    ptr1 = &testNum

    fmt.Println(&testNum)
    fmt.Println(ptr1)
    fmt.Println(*ptr1)
}
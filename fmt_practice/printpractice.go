package main

import (
	"bufio"
	"fmt"
	"os"
)
func main(){
	var a int
var b int
var c int
a, b, c = 10, 20, 30

a, b = b, a
	fmt.Println(a,b,c)
	stdin := bufio.NewReader(os.Stdin)
	var name1 string = "tetetete"
    var age1 int
	var name2 string 
    fmt.Print("Enter your name and age: ")
    
    fmt.Scan(&name1, &age1)
    fmt.Printf("Scan - Name: %s\nAge: %d\n", name1, age1)
	fmt.Println("----------")
	stdin.ReadString('\n')
    fmt.Scanln(&name2)
    fmt.Printf("Scanln - Name: %s\n", name2)
	fmt.Println("----------")
}








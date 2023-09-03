package main

import (
	"bufio"
	"fmt"
	"os"
)


func main(){
	rx := bufio.NewReader(os.Stdin)
	wx := bufio.NewWriter(os.Stdout)
	defer wx.Flush()

	var testNum int

	fmt.Fscanln(rx, &testNum)
	var x1, y1, x2, y2 int
	fmt.Fscanln(rx, &x1, &y1, &x2, &y2)
	var planetsNum int
	fmt.Fscanln(rx,&planetsNum)
	
	fmt.Println(testNum)
	fmt.Println(x1, y1, x2, y2)
	fmt.Println(planetsNum)
}
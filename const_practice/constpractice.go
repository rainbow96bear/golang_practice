package main

import "fmt"
const (
	Red int = iota
	Green int = iota
	Blue int = iota
)
const (
	Blind uint8 = 1 << iota
	Burn 
	Cold 
	Paralyze
	Poison
	Silence
	Slow
	Stun
)
const (
	testInt int = 8
	testNum = 10
	NonTypeFloat = 3.14
	PI float64 = 3.14
)
func main(){

	fmt.Println("Red : ",Red)
	fmt.Println("Green : ",Green)
	fmt.Println("Blue : ",Blue)
	fmt.Println("--------------------------")
	fmt.Println("Blind : ",Blind )
	fmt.Println("Burn : ",Burn)
	fmt.Println("Cold : ",Cold)
	fmt.Println("Paralyze : ",Paralyze)
	fmt.Println("Poison : ",Poison)
	fmt.Println("Silence : ",Silence)
	fmt.Println("Slow : ",Slow)
	fmt.Println("Stun : ",Stun)

	var a = testInt * 10
	var b uint16 = testNum * 10

	fmt.Println(a,b)
}
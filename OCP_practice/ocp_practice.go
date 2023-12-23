package main

import (
	"fmt"
)

type Animal interface{
	Speak()
}

type Dog struct{}

func (d *Dog) Speak() {
	fmt.Println("왈왈")
}

type Cat struct{}

func (c *Cat) Speak(){
	fmt.Println("야옹")
}

func AnimalSpeak(a Animal){
	a.Speak()
}

func main(){
	dog := &Dog{}
	cat := &Cat{}

	AnimalSpeak(dog)
	AnimalSpeak(cat)
}
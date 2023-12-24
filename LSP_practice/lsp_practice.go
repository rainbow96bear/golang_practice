package main

import "fmt"

type Dog struct{}

func (d *Dog) Speak() {
	fmt.Println("왈왈")
}

type Puddle Dog

func (p *Puddle) Speak(){
	fmt.Println("으르렁")
}

func (p *Puddle) Sleep(){
	fmt.Println("푸들이 잔다.")
}

func main() {
	dog := &Dog{}
	// dog := &Puddle{}
	dog.Speak()
	// if
	// puddle := &Puddle{}
	// puddle.Sleep()
	// if
	// if 사이의 코드에서 &Puddle{}를 &Dog{}로 변경 시 오류
}
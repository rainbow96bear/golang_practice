package main

import (
	"golang_practice/DIP_practice/Developer"
	IPerson "golang_practice/DIP_practice/I_Person"
	"golang_practice/DIP_practice/Singer"
)

func main() {
	developer := &Developer.Developer{}
	singer := &Singer.Singer{}

	// Worker 인터페이스를 구현한 객체를 주입하여 Work 함수 호출
	IPerson.Work(developer)
	IPerson.Work(singer)
}

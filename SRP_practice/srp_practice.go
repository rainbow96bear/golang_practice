package main

import (
	"fmt"
)

// 행동 클래스
type Act struct{}

func (a *Act) Walk() {fmt.Println("걸어요")}
func (a *Act) Stop() {fmt.Println("멈춰요")}


// 감정 클래스
type Emotion struct{}

func (e *Emotion) Smile() {fmt.Println("웃어요")}
func (e *Emotion) Cry() {fmt.Println("울어요")}

// 사람 클래스 (구조체 내장)
type Person struct {
	Act
	Emotion
}

func main(){
	var person Person
	person.Walk()
	person.Stop()
	person.Smile()
	person.Cry()
}
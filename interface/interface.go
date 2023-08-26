package main

import (
	"fmt"
)

type LegacyPrinter struct{}

func (lp *LegacyPrinter) Print(s string) {
    fmt.Println("Legacy Printer:", s)
}

type LegacyPrinterAdapter struct{
	la LegacyPrinter
}

func (l LegacyPrinterAdapter) PrintStored(s ...string){
	for _,v :=range s {
		l.la.Print(v)
		break
	}
}

type NewPrinter struct{}

func (np *NewPrinter) newPrint(s1, s2 string) {
    fmt.Printf("newPrinter로 출력 : %s 그리고 %s \n",s1, s2)
}

type NewAdapter struct{
	na NewPrinter
}

func (n NewAdapter) PrintStored(s ...string){
	for i,_ :=range s {
		n.na.newPrint(s[i],s[i+1])
		break
	}
	
}

type ModernPrinter interface {
    PrintStored(s ...string)
}

func Print (m ModernPrinter, s ...string){
	m.PrintStored(s...)
}

func main(){
	legacy := LegacyPrinter{}
	legacy.Print("legacy입니다.")

	new := NewPrinter{}
	new.newPrint("첫 번째 string","두 번째 string")

	legacyAdapter  := LegacyPrinterAdapter{legacy}
	legacyAdapter.PrintStored("legacy Adapter입니다.","Adapter로 프린팅 중입니다.")
	
	NewAdapter  := NewAdapter{new}
	NewAdapter.PrintStored("newAdapter 입력1","newAdapter 입력2")
	
	Print(legacyAdapter,"출력1","출력2","출력3")
	Print(NewAdapter,"출력1","출력2","출력3")
}
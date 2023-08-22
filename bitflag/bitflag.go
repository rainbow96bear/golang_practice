package main

import "fmt"

const (
	Blind uint8 = 1 << iota
	Burn 
	Cold 
	Paralyze
)

func SetState(states, state uint8) uint8{
    return states | state
}

func ResetState(states, state uint8) uint8{
    return states &^ state
}

func IsStateOn(states, state uint8) bool {
    return states & state == state
}

func CheckState (states uint8) {
    
    if IsStateOn(states, Blind) {
        fmt.Println("Blind 적용중")
    }
    if IsStateOn(states, Burn) {
        fmt.Println("Burn 적용중")
    }
    if IsStateOn(states, Cold) {
        fmt.Println("Cold 적용중")
    }
    if IsStateOn(states, Paralyze) {
        fmt.Println("Paralyze 적용중")
    }
}

func main () {
    var states uint8 = 0
    states = SetState(states, Blind)
    states = SetState(states, Paralyze)
    states = SetState(states, Cold)
    
    states = ResetState(states, Blind)
    
    CheckState(states)
}	
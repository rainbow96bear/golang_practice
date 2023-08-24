package main

import "fmt"

type Character struct {
	Name  string
	Level uint
	Exe   float64
}
type Address struct {
    Street  string
    City    string
    Country string
}

type Person struct {
    Name string
    Address
	Country string
}

type Building struct {
    Name string
    Address
}

func main() {
	var character1 Character = Character{"무지개곰1", 1, 0.0}
	character2 := Character{"무지개곰2", 1, 0.0}
	character3 := Character{Name:"무지개곰3", Exe:1.0}
	character4 := new(Character)
	character4.Name="무지개곰4"
	fmt.Println(character1)
	fmt.Println(character2)
	fmt.Printf("%T %v\n",character3, character3)
	fmt.Printf("%T %v\n",character4, character4)

	// person1 := Person{"무지개곰" ,Address{"거리1", "도시1", "나라1"}}
	building := Building{"무지개 타워", Address{"거리2", "도시2", "나라1"}}
	// fmt.Println(person1)
	fmt.Println(building)
	// fmt.Println(person1.LivingAddress.Country)
	fmt.Println(building.Address.Country)
	person := Person{"무지개곰", Address{"거리1", "도시1", "거주지"}, "출생지"}

    fmt.Println(person.Country,person.Address.Country)
}
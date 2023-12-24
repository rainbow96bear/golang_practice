package Developer

import "fmt"

type Developer struct{}

func (d *Developer) Work() {
	d.Programming()
}

func (d *Developer) Programming() {
	fmt.Println("프로그래밍 합니다.")
}
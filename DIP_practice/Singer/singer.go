package Singer

import "fmt"

type Singer struct{}

func (s *Singer) Work() {
	s.Singing()
}

func (s *Singer) Singing() {
	fmt.Println("노래를 부릅니다.")
}
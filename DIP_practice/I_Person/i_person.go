package IPerson

type Person interface {
	Work()
}

func Work(p Person) {
	p.Work()
}
package main

func nextValue() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	next := nextValue()

	println("next", next()) // 1
	println("next", next()) // 2
	println("next", next()) // 3

	anotherNext := nextValue()
	println("anotherNext", anotherNext()) // 1 다시 시작
	println("anotherNext", anotherNext()) // 2

	println("next", next())
	println("anotherNext", anotherNext())
}
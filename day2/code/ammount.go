package main

import "fmt"

func add(input string, payments ...int) (str string) {
	var sum int
	for _, i := range payments {
		sum += i
	}
	str = fmt.Sprint(input, sum)
	return
}

type CustomErr struct {
	errno   int
	message string
}

func (c CustomErr) Error() string {
	return c.message
}

func (c CustomErr) ErrNo() int {

}

func main() {
	println(add("T: ", 10, 20, 30, -20))
	a := 1
	f := func() {
		println(a)
	}
	a = 2
	f() //2
}

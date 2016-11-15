package main

import (
	"fmt"
)

// START OMIT
type Adder interface {
	Add(int)
	fmt.Stringer
}

type IntAdder int

func (a *IntAdder) Add(v int) {
	*a += IntAdder(v)
}

func (a IntAdder) String() string {
	return fmt.Sprintf("%d", a)
}

//END OMIT

type StructAdder struct {
	value int
}

func (a *StructAdder) Add(v int) {
	a.value += v
}

func (a StructAdder) String() string {
	return fmt.Sprintf("%d", a.value)
}

func addAndPrint(adder Adder, v int) {
	adder.Add(v)
	fmt.Println(adder)
}

func main() {
	aInt := new(IntAdder)
	addAndPrint(aInt, 123)
	sInt := new(StructAdder)
	addAndPrint(sInt, 987)
	s2Int := &StructAdder{800}
	addAndPrint(s2Int, 200)
}

package main

func a() {
	v := 1
	defer println(v)
	v = 2
	defer func() {
		println(v)
	}()
	v = 3
}

func main() {
	a()
}

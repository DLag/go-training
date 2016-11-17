package main

func add(payments ...int) (sum int) {
	sum = len(payments)
	for _, i := range payments {
		sum += i
	}
	return
}

func main() {
	println("Total:", add(10, 20, 30, -20))
}

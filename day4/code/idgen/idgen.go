package main

import (
	"fmt"
	"time"

	"./generator"
)

func main() {
	chanGenerator := generator.NewChanGenerator()
	for {
		fmt.Println(chanGenerator.Generate())
		time.Sleep(time.Millisecond * 200)
	}
}

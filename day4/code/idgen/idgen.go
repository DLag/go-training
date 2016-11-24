package main

import (
	"fmt"
	"time"

	g "./generator"
)

func main() {
	chanGenerator := g.NewChanGenerator()
	for {
		fmt.Println(chanGenerator.Generate())
		time.Sleep(time.Millisecond * 200)
	}
}

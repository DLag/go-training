package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"

	"./generator"
)

func main() {
	http.Handle("/atomic", generator.NewIdGenHandler(generator.NewIdGeneratorAtomic()))
	http.Handle("/mutex", generator.NewIdGenHandler(generator.NewIdGeneratorMutex()))
	http.Handle("/chan", generator.NewIdGenHandler(generator.NewIdGeneratorChan()))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

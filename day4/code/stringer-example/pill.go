//go:generate stringer -type=Pill

package main

type Pill int

const (
	Placebo Pill = iota
	Aspirin
	Ibuprofen
	LSD
	Paracetamol
	Acetaminophen = Paracetamol
)

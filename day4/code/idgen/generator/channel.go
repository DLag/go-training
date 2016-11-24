package generator

type ChanGenerator struct {
	c chan int
}

func (g *ChanGenerator) Close() {
	println("closed")
}

// Generate generates uniq
func (g *ChanGenerator) Generate() int {
	return <-g.c
}

func (g *ChanGenerator) worker() {
	var i int
	for {
		i++
		g.c <- i
	}
}

// NewChanGenerator creates structure of ChanGenerator and initializes internal worker
func NewChanGenerator() *ChanGenerator {
	generator := ChanGenerator{make(chan int)}
	go generator.worker()
	return &generator
}

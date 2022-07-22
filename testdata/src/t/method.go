package t

type counter struct {
	i int
}

func (c *counter) inc() {
	c.i++
}

func failMethod(counter, kCounter *counter) {
	counter.inc()
	kCounter.inc() // want "write to const variable 'kCounter'"
}

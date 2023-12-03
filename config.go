package pqr

// config struct for parsing options on initialization.
type config struct {
	l Logger
	q SQL
}

// setLogger to config.
func (c *config) setLogger(l Logger) {
	c.l = l
}

// setQuerier to config. Panic on duplicate.
func (c *config) setQuerier(q SQL) {
	if c.q != nil {
		panic("connection has already been attached")
	}
	c.q = q
}

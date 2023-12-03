package pqr

// Key interface represents a list of allowed keys.
type Key interface {
	int32 | uint32 | int64 | uint64 | string
}

type keyInfo struct {
	index int
	name  string
}

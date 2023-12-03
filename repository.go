package pqr

// Repository interface is the main set of methods to the repository.
type Repository[T any, K Key] interface {
	querier
	transactioner[T, K]
	orm[T, K]
}

type repository[T any, K Key] struct {
	name string
	m    *meta
	q    SQL
	l    Logger
}

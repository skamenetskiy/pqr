package pqr

const (
	OrderNone Order = ""
	OrderASC  Order = "ASC"
	OrderDESC Order = "DESC"
)

type Order string

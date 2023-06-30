package db

type UpdateAccountParams struct {
	ID          int
	Title       string
	Description string
	Value       int
}

type Store interface {
	UpdateAccount(arg UpdateAccountParams) error
}

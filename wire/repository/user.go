package repository

type User interface {
	Get(id string) (string, error)
}
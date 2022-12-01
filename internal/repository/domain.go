package repository

type Repo interface {
	GetQuestion() ([]string, error)
}

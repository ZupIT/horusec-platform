package repository

type IRepository interface {
}

type Repository struct {
}

func NewRepository() IRepository {
	return &Repository{}
}

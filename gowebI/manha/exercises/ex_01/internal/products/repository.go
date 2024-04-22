package products

import "github.com/batatinha123/bootcamp-meli-web/pkg/store"

type Filter struct {
	Name      string
	Published *bool
}

type Repository interface {
	GetAll(filter Filter) ([]Product, error)
	Store(name, code, color string, count int, price float64, published bool) (Product, error)
	Update(id uint64, name, code, color string, count int, price float64, published bool) (Product, error)
	UpdateName(id uint64, name string) (Product, error)
	LastID() (uint64, error)
	Delete(id uint64) error
}

func NewRepository(db store.Store) Repository {
	return &FileRepository{db}
}

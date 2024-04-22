package products

type Repository interface {
	GetAll() ([]Product, error)
	Store(name, category string, count int, price float64) (Product, error)
	Update(id uint64, name, productType string, count int, price float64) (Product, error)
	UpdateName(id uint64, name string) (Product, error)
	LastID() (uint64, error)
	Delete(id uint64) error
}

func NewRepository() Repository {
	return &MemoryRepository{}
}

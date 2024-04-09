package products

type Repository interface {
	GetAll() ([]Product, error)
	Store(name, category string, count int, price float64) (Product, error)
	LastID() (uint64, error)
}

func NewRepository() Repository {
	return &MemoryRepository{}
}

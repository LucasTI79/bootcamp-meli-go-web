package products

var ps []Product
var lastID uint64 = 0

type MemoryRepository struct {
}

func (m *MemoryRepository) GetAll() ([]Product, error) {
	return ps, nil
}

func (m *MemoryRepository) Store(name, category string, count int, price float64) (Product, error) {
	lastID++
	p := Product{
		ID:       lastID,
		Name:     name,
		Category: category,
		Count:    count,
		Price:    price,
	}
	ps = append(ps, p)
	return p, nil
}

func (m *MemoryRepository) LastID() (uint64, error) {
	return lastID, nil
}

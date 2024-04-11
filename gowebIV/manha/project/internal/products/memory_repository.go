package products

import (
	"fmt"
)

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

func (m *MemoryRepository) Update(id uint64, name, category string, count int, price float64) (Product, error) {
	p := Product{
		Name:     name,
		Category: category,
		Count:    count,
		Price:    price,
	}
	updated := false
	for i := range ps {
		if ps[i].ID == id {
			p.ID = id
			ps[i] = p
			updated = true
		}
	}
	if !updated {
		return Product{}, fmt.Errorf("product %d not found", id)
	}
	return p, nil
}

func (m *MemoryRepository) UpdateName(id uint64, name string) (Product, error) {
	var p Product
	updated := false
	for i := range ps {
		if ps[i].ID == id {
			ps[i].Name = name
			updated = true
			p = ps[i]
		}
	}
	if !updated {
		return Product{}, fmt.Errorf("product %d not found", id)
	}

	return p, nil
}

func (r *MemoryRepository) Delete(id uint64) error {
	deleted := false
	var index int
	for i := range ps {
		if ps[i].ID == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("product %d not found", id)
	}

	//[1,2,3]
	// spread operator -> ellipsis
	ps = append(ps[:index], ps[index+1:]...)
	return nil
}

func (m *MemoryRepository) LastID() (uint64, error) {
	return lastID, nil
}

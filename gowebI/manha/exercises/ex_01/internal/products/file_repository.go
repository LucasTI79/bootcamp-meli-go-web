package products

import (
	"time"

	"github.com/batatinha123/bootcamp-meli-web/pkg/store"
)

type FileRepository struct {
	db store.Store
}

func NewFileRepository(db store.Store) Repository {
	return &FileRepository{
		db: db,
	}
}

func (r *FileRepository) GetAll() ([]Product, error) {
	var ps []Product
	r.db.Read(&ps)
	return ps, nil
}

func (r *FileRepository) Store(name, code, color string, count int, price float64, published bool) (Product, error) {
	p := Product{
		Name:      name,
		Code:      code,
		Color:     color,
		Count:     count,
		Price:     price,
		Published: published,
		CreatedAt: time.Now().String(),
	}

	var ps []Product

	// primeiro lemos o arquivo
	r.db.Read(&ps)

	// calculamos qual o próximo ID
	lastIdInserted := len(ps)
	lastIdInserted++
	p.ID = uint64(lastIdInserted)

	// inserimos o produto a ser cadastrado no slice de produtos
	ps = append(ps, p)

	// gravamos no arquivo novamente com o novo produto inserido
	err := r.db.Write(ps)
	if err != nil {
		return Product{}, err
	}
	return p, nil
}

func (r *FileRepository) Delete(id uint64) error {
	return nil
}

func (r *FileRepository) Update(id uint64, name, code, color string, count int, price float64, published bool) (Product, error) {
	return Product{}, nil
}
func (r *FileRepository) UpdateName(id uint64, name string) (Product, error) {
	return Product{}, nil
}

func (r *FileRepository) LastID() (uint64, error) {
	var ps []Product
	if err := r.db.Read(&ps); err != nil {
		return 0, err
	}

	if len(ps) == 0 {
		return 0, nil
	}

	return ps[len(ps)-1].ID, nil

}

package products

import "errors"

var (
	ErrProductAlreadyExists = errors.New("product already exists")
	ErrProductNotFound      = errors.New("product not found")
)

type Service interface {
	GetAll() ([]Product, error)
	Store(name, code, color string, count int, price float64, published bool) (Product, error)
	Update(id uint64, name, code, color string, count int, price float64, published bool) (Product, error)
	UpdateName(id uint64, name string) (Product, error)
	Delete(id uint64) error
}

type service struct {
	repository Repository
}

func (s *service) GetAll() ([]Product, error) {
	produtos, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return produtos, nil

}

func (s *service) Store(name, code, color string, count int, price float64, published bool) (Product, error) {
	// aqui poderiamos também através da service enviar o id ao repositório caso quisessemos
	// lastID, err := s.repository.LastID()
	// if err != nil {
	// 	return Product{}, err
	// }

	// lastID++

	product, err := s.repository.Store(name, code, color, count, price, published)
	if err != nil {
		return Product{}, err
	}

	return product, nil
}

func (s *service) Update(id uint64, name, code, color string, count int, price float64, published bool) (Product, error) {
	return s.repository.Update(id, name, code, color, count, price, published)
}

func (s *service) UpdateName(id uint64, name string) (Product, error) {
	return s.repository.UpdateName(id, name)
}

func (s *service) Delete(id uint64) error {
	return s.repository.Delete(id)
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

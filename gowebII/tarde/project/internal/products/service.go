package products

type Service interface {
	GetAll() ([]Product, error)
	Store(name, category string, count int, price float64) (Product, error)
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

func (s *service) Store(name, category string, count int, price float64) (Product, error) {
	// aqui poderiamos também através da service enviar o id ao repositório caso quisessemos
	// lastID, err := s.repository.LastID()
	// if err != nil {
	// 	return Product{}, err
	// }

	// lastID++

	product, err := s.repository.Store(name, category, count, price)
	if err != nil {
		return Product{}, err
	}

	return product, nil
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

package product

type ProductRepository interface {
	ProductInsert(Product) error
	FindProductByID(int) (Product, error)
}

type Service struct {
	repo ProductRepository
}

func NewService(repo ProductRepository) *Service {
	return &Service{repo: repo}
}

var runid = 0

func (s Service) Register(name string, amount, stock int) (int, error) {
	runid++
	m := Product{
		ID:     runid,
		Name:   name,
		Amount: amount,
		Stock:  stock,
	}
	if err := s.repo.ProductInsert(m); err != nil {
		return runid, err
	}
	return runid, nil
}

func (s Service) Information(id int) (Product, error) {
	m, err := s.repo.FindProductByID(id)
	if err != nil {
		return Product{}, err
	}
	return m, nil
}

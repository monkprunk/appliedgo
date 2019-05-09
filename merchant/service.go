package merchant

import (
	"Advance-Golang-Programming/advanced/final/product"
	"errors"
)

var (
	IDCannotBeEmpty = errors.New("id cannot be empty")
)

type merchantRepository interface {
	MerchantInsert(Merchant) error
	FindMerchantByID(int) (Merchant, error)
	AddProductByID(int, product.Product) error
	FindProductByMerchantID(int) ([]product.Product, error)
}

type Service struct {
	repo merchantRepository
}

func NewService(repo merchantRepository) *Service {
	return &Service{repo: repo}
}

func (s Service) Register(name, bankaccount string) (int, error) {
	runid++
	m := Merchant{
		ID:          runid,
		Name:        name,
		BankAccount: bankaccount,
	}
	if err := s.repo.MerchantInsert(m); err != nil {
		return runid, err
	}
	return runid, nil
}

func (s Service) Information(id int) (Merchant, error) {
	m, err := s.repo.FindMerchantByID(id)
	if err != nil {
		return Merchant{}, err
	}
	return m, nil
}

func (s Service) Add(id int, name string, amount int, stock int) (int, error) {
	runidproduct++
	p := product.Product{
		ID:     runidproduct,
		Name:   name,
		Amount: amount,
		Stock:  stock,
	}
	if err := s.repo.AddProductByID(id, p); err != nil {
		return runidproduct, err
	}
	return runidproduct, nil
}

func (s Service) Find(id int) ([]product.Product, error) {
	p, err := s.repo.FindProductByMerchantID(id)
	if err != nil {
		return []product.Product{}, err
	}
	return p, nil
}

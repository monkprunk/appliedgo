package merchant

import "Advance-Golang-Programming/advanced/final/product"

var merchantDB []Merchant

type MemoryRepo struct{}

func NewMemoryDB() *MemoryRepo {
	return &MemoryRepo{}
}

func (mr MemoryRepo) MerchantInsert(m Merchant) error {
	merchantDB = append(merchantDB, m)
	return nil
}

func (mr MemoryRepo) FindMerchantByID(id int) (Merchant, error) {
	for _, m := range merchantDB {
		if m.ID == id {
			return m, nil
		}
	}
	return Merchant{}, nil
}

func (mr MemoryRepo) AddProductByID(id int, product product.Product) error {
	for i := 0; i < len(merchantDB); i++ {
		if merchantDB[i].ID == id {
			merchantDB[i].Products = append(merchantDB[i].Products, product)
			return nil
		}
	}
	return nil
}

func (mr MemoryRepo) FindProductByMerchantID(id int) ([]product.Product, error) {
	for i := 0; i < len(merchantDB); i++ {
		if merchantDB[i].ID == id {
			return merchantDB[i].Products, nil
		}
	}
	return []product.Product{}, nil
}

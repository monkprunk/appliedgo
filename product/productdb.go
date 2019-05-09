package product

var productDB []Product

type MemoryRepo struct{}

func (mr MemoryRepo) ProductInsert(merId int, m Product) error {
	productDB = append(productDB, m)
	return nil
}

func (mr MemoryRepo) FindProductByID(id int) (Product, error) {
	for _, m := range productDB {
		if m.ID == id {
			return m, nil
		}
	}
	return Product{}, nil
}

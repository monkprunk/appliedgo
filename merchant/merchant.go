package merchant

import (
	"Advance-Golang-Programming/advanced/final/product"
)

type Merchant struct {
	ID          int               `json:"id"`
	Name        string            `json:"name"`
	BankAccount string            `json:"bankaccount"`
	Username    string            `json:"username"`
	Password    string            `json:"password"`
	Products    []product.Product `json:"product"`
}

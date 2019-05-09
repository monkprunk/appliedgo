package product

type Product struct {
	ID     int    `json:"id"`
	Name   string `json:"name" bson:"name"`
	Amount int    `json:"amount" bson:"amount"`
	Stock  int    `json:"stock" bson:"stock"`
}

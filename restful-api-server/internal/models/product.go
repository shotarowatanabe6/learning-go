package models

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type Products []Product

func (p *Products) GetAll() Products {
	return *p
}

func (p *Products) Get(id int) Product {
	return (*p)[id]
}

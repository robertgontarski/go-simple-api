package models

import "time"

type Product struct {
	ID      int64      `json:"id"`
	Name    string     `json:"name"`
	Price   float64    `json:"price"`
	Created *time.Time `json:"created"`
	Updated *time.Time `json:"updated"`
	Deleted *time.Time `json:"deleted"`
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		Name:  name,
		Price: price,
	}
}

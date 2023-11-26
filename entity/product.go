package entity

import (
	"encoding/json"
	"os"
)

type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	IsAvailable bool    `json:"isAvailable"`
}

func GetProducts() ([]byte, error) {
	// Read JSON file
	data, err := os.ReadFile("../entity/data.json")
	if err != nil {
		return nil, err
	}
	return data, nil
}

func AddProduct(product Product) error {
	// Load existing products and append the data to product list
	var products []Product
	data, err := os.ReadFile("../entity/data.json")
	if err != nil {
		return err
	}
	// Load our JSON file to memory using array of products
	err = json.Unmarshal(data, &products)
	if err != nil {
		return err
	}
	// Add new Product to our list
	products = append(products, product)

	// Write Updated JSON file
	updatedData, err := json.Marshal(products)
	if err != nil {
		return err
	}
	err = os.WriteFile("../entity/data.json", updatedData, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

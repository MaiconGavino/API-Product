package repository

import (
	"apiproducts/src/pb/products"
	"fmt"
	"google.golang.org/protobuf/proto"
	"os"
)

type ProductRepository struct{}

const filename string = "products.txt"

func (pr *ProductRepository) loadData() (products.ProductList, error) {
	productList := products.ProductList{}

	data, err := os.ReadFile(filename)
	if err != nil {
		return productList, fmt.Errorf("failed to read products file: %w", err)
	}
	err = proto.Unmarshal(data, &productList)
	if err != nil {
		return productList, fmt.Errorf("failed to unmarshal. error: %w", err)
	}
	return productList, nil
}

func (pr *ProductRepository) saveData(productList *products.ProductList) error {
	data, err := proto.Marshal(productList)
	if err != nil {
		return fmt.Errorf("failed to marshal. error: %w", err)
	}
	err = os.WriteFile(filename, data, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to save products. error: %w", err)
	}

	return nil
}
func (pr *ProductRepository) Create(product products.Product) (products.Product, error) {
	productList, err := pr.loadData()
	if err != nil {
		return product, err
	}
	product.Id = int32(len(productList.Product) + 1)
	productList.Product = append(productList.Product, &product)
	err = pr.saveData(&productList)

	return product, err

}

func (pr *ProductRepository) FindAll() (products.ProductList, error) {
	return pr.loadData()
}

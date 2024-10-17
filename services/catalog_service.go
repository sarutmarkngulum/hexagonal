package services

import "goredis/repositories"

type catalogservice struct {
	productRepo repositories.ProductRepository
}

func NewCatalogService(productRepo repositories.ProductRepository) CatalogService {
	return catalogservice{productRepo}
}

func (s catalogservice) GetProducts() (products []Product, err error) {
	productsDB, err := s.productRepo.GetProducts()
	if err != nil {
		return nil, err
	}

	for _, p := range productsDB {
		products = append(products, Product{
			ID:       p.ID,
			Name:     p.Name,
			Quantity: p.Quantity,
		})
	}

	return products, nil
}

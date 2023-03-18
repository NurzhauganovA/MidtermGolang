package service

import (
	"github.com/NurzhauganovA/online-store/endpoint"
	"github.com/NurzhauganovA/online-store/package/repository"
)

type ProductService struct {
	repo         repository.Product
	categoryRepo repository.Category
}

func NewProductService(repo repository.Product, categoryRepo repository.Category) *ProductService {
	return &ProductService{repo: repo, categoryRepo: categoryRepo}
}

func (s *ProductService) Create(userId, categoryId int, product endpoint.Product) (int, error) {
	_, err := s.categoryRepo.GetById(userId, categoryId)
	if err != nil {
		// list does not exists or does not belongs to user
		return 0, nil
	}

	return s.repo.Create(categoryId, product)
}

func (s *ProductService) GetAll(userId, categoryId int) ([]endpoint.Product, error) {
	return s.repo.GetAll(userId, categoryId)
}

func (s *ProductService) GetById(userId, productId int) (endpoint.Product, error) {
	return s.repo.GetById(userId, productId)
}

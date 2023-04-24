package service

import (
	"github.com/NurzhauganovA/online-store/endpoint"
	"github.com/NurzhauganovA/online-store/package/repository"
)

type ProductService struct {
	repo         repository.Product
	categoryRepo repository.Category
}

func (s *ProductService) CreateRating(rating endpoint.Rating) ([]endpoint.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ProductService) CreateComment(comment, productId int) ([]endpoint.Product, error) {
	return s.repo.CreateComment(comment, productId), nil
}

func (s *ProductService) Rate(userId, productId int) (endpoint.Product, error) {
	return s.repo.Rate(userId, productId)
}

func (s *ProductService) GetFilteredProducts(price float64, price2 float64, rating int) (interface{}, interface{}) {
	return s.repo.GetFilteredProducts(price, price2, rating)
}

func NewProductService(repo repository.Product, categoryRepo repository.Category) *ProductService {
	return &ProductService{repo: repo, categoryRepo: categoryRepo}
}

func (s *ProductService) Create(userId, categoryId int, product endpoint.Product) (int, error) {
	_, err := s.categoryRepo.GetById(userId, categoryId)
	if err != nil {
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

func (s *ProductService) GetQueryParam(userId, productTitle string) (endpoint.Product, error) {
	return s.repo.GetQueryParam(userId, productTitle)
}

package service

import (
	"github.com/NurzhauganovA/online-store/endpoint"
	"github.com/NurzhauganovA/online-store/package/repository"
)

type CategoryProductService struct {
	repo repository.Category
}

func NewCategoryProductService(repo repository.Category) *CategoryProductService {
	return &CategoryProductService{repo: repo}
}

func (s *CategoryProductService) Create(userId int, category endpoint.Category) (int, error) {
	return s.repo.Create(userId, category)
}

func (s *CategoryProductService) GetAll(userId int) ([]endpoint.Category, error) {
	return s.repo.GetAll(userId)
}

func (s *CategoryProductService) GetById(userId, categoryId int) (endpoint.Category, error) {
	return s.repo.GetById(userId, categoryId)
}

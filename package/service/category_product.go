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

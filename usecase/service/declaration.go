package service

import (
	"slack-api/domain"
	"slack-api/usecase/repository"
)

type declService struct {
	r repository.DeclarationRepository
}

type DeclarationService interface {
	Create(decl *domain.Declaration) error
}

func NewDeclarationService(r repository.DeclarationRepository) DeclarationService {
	return &declService{r}
}

func (s *declService) Create(decl *domain.Declaration) error {
	err := s.r.Store(decl)
	return err
}

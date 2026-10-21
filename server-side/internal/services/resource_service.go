package services

import (
	"context"

	"github.com/kishanravi887321/pumkin/server-side/internal/models"
)

type ResourceRepository interface {
	List(ctx context.Context) ([]models.Resource, error)
}

type ResourceService struct {
	repository ResourceRepository
}

func NewResourceService(repository ResourceRepository) *ResourceService {
	return &ResourceService{repository: repository}
}

func (s *ResourceService) List(ctx context.Context) ([]models.Resource, error) {
	return s.repository.List(ctx)
}

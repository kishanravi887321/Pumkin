package repository

import (
	"context"

	"github.com/kishanravi887321/pumkin/server-side/internal/models"
)

type ResourceRepository struct {
}

func NewResourceRepository() *ResourceRepository {
	return &ResourceRepository{}
}

func (r *ResourceRepository) List(context.Context) ([]models.Resource, error) {
	return []models.Resource{}, nil
}

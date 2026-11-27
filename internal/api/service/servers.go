package service

import (
	"context"

	"github.com/asim9115/incidentix/internal/api/repository"
)

type ServersService struct {
	repo *repository.ServerRepo
}

func NewServersService(repo *repository.ServerRepo) *ServersService {
	return &ServersService{
		repo : repo,
	}
}

func (s *ServersService) CreateServer(ctx context.Context, server repository.Server) error {
	
	err := s.repo.AddServer(ctx, &server)
}
package service

import (
	"context"
	"crypto/rand"

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

func (s *ServersService) CreateServer(ctx context.Context, server *repository.Server) error {

	server.ID = rand.Text()[:8]
	server.Token = GenerateToken()
	if err := s.repo.AddServer(ctx, server); err != nil {
		return err
	}

	return nil
}

func (s *ServersService) GetAll(ctx context.Context) (*[]repository.Server, error) {
	servers, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err 
	}

	return servers, nil
}

func GenerateToken()(string) {
	prefix := "idx-"
	id := rand.Text()
	return prefix+id[:8]
} 
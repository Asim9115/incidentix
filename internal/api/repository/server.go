package repository

import (
	"context"

	"github.com/asim9115/incidentix/internal/api/database"
)

type ServerRepo struct {
	db *database.DataBase
}

func NewServersService(db *database.DataBase) *ServerRepo {
	return &ServerRepo{
		db: db,
	}
}

type Server struct {
	ID          string
	Name        string
	Token       string
	TotalCPU    int
	TotalMemory int
	TotalSpace  int
}

func (s *ServerRepo) AddServer(ctx context.Context, srv *Server) error {
	_, err := s.db.Pool.Exec(ctx, `INSERT INTO servers (id, name, token, total_cpu, total_ram, total_space)
	VALUES ($1, $2, $3, $4, $5, $6)`, srv.ID, srv.Name, srv.Token, srv.TotalCPU, srv.TotalMemory, srv.TotalSpace)
	if err != nil {
		return err
	}
	return nil
}

func (s *ServerRepo) GetServerByName(ctx context.Context, name string) (*Server, error) {
	row := s.db.Pool.QueryRow(ctx, `SELECT id, name, token, total_cpu, total_ram, total_space from 
					servers where name=$1`, name)

	var server Server

	err := row.Scan(
		&server.ID,
		&server.Name,
		&server.Token,
		&server.TotalCPU,
		&server.TotalMemory,
		&server.TotalSpace,
	)

	if err != nil {
		return nil, err
	}

	return &server, nil
}

func (s *ServerRepo) GetAll(ctx context.Context) (*[]Server, error) {
	rows, err := s.db.Pool.Query(ctx, `SELECT id, name, token, total_ram, total_cpu, total_space FROM servers`)

	if err != nil {
		return nil ,err
	}
	defer rows.Close()

	var servers []Server
	for rows.Next() {
		var server Server
		if err := rows.Scan(
			&server.ID, &server.Name, &server.Token, &server.TotalCPU, &server.TotalMemory, &server.TotalSpace,
		); err != nil {
			return nil, err
		}

		servers = append(servers, server)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &servers, nil
}
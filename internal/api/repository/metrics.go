package repository

import (
	"context"
	"time"

	"github.com/asim9115/incidentix/internal/api/database"
)

type MetricsRepo struct {
	db *database.DataBase
}

type Metrics struct {
	ID string
	ServerID string
	CPU float64
	RAM float64
	SPACE float64
	TS time.Time
}

func NewMetricsRepo(db *database.DataBase) *MetricsRepo {
	return &MetricsRepo{db: db}
}

func (r *MetricsRepo) Insert(ctx context.Context, id, serverID string, cpu, ram, space float64, ts time.Time) error {
	_, err := r.db.Pool.Exec(ctx, `
		INSERT INTO metrics (id, server_id, ts, cpu, ram, space)
		VALUES ($1, $2, $3, $4, $5, $6)
	`, id, serverID, ts, cpu, ram, space)

	if err != nil {
		return err
	}

	return nil
}

func (r *MetricsRepo) DisplayAll(ctx context.Context) (*[]Metrics, error) {
	rows, err := r.db.Pool.Query(ctx, `SELECT * FROM metrics`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var metrics []Metrics

	for rows.Next(){
		var m Metrics

		err := rows.Scan(
			&m.ID, &m.CPU, &m.RAM, &m.SPACE, &m.ServerID, &m.TS,
		)
		if err != nil {
			return nil, err
		}

		metrics = append(metrics, m)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &metrics, nil
}
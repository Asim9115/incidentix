package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DataBase struct{
	url string
	Pool *pgxpool.Pool
}

func New(url string) *DataBase{
	return &DataBase{
		url: url,
	}
}

func (d *DataBase)Connect(ctx context.Context) (error) {
	pool, err := pgxpool.New(ctx, d.url)
	if err != nil {
		return err
	}

	if err := pool.Ping(ctx); err != nil {
		return err
	}

	d.Pool = pool

	return nil
}

func (d *DataBase) Close() {
	if d.Pool != nil {
		d.Pool.Close()
	}
}

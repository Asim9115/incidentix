package database

import (
	"context"

_ "github.com/golang-migrate/migrate/v4/database/postgres"
_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4"
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
		pool.Close()
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

func (d *DataBase)RunMigrations() error {
	m, err := migrate.New("file://../../internal/api/database/migrations", d.url)
	if err != nil {
		return err
	}
	defer m.Close()

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}
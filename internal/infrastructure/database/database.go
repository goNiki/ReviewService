package database

import (
	"context"
	"fmt"
	"time"

	"github.com/goNiki/ReviewService/internal/infrastructure/config"
	"github.com/goNiki/ReviewService/models/errorapp"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Db struct {
	Pool *pgxpool.Pool
}

func InitDatabase(cfg *config.DBConfig) (*Db, error) {

	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name, cfg.SslMode)

	PoolCfg, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return &Db{}, fmt.Errorf("%w: %v", errorapp.ErrInitDatabase, err)
	}

	PoolCfg.MaxConns = cfg.MaxConns
	PoolCfg.MinConns = cfg.MinConns
	PoolCfg.MaxConnLifetime = cfg.MaxConnLifeTime
	PoolCfg.MaxConnIdleTime = cfg.MaxConnIdleTime
	PoolCfg.HealthCheckPeriod = cfg.HealthCheckPeriod

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	Pool, err := pgxpool.NewWithConfig(ctx, PoolCfg)
	if err != nil {
		return &Db{}, fmt.Errorf("%w: %v", errorapp.ErrInitDatabase, err)
	}

	if err := Pool.Ping(ctx); err != nil {
		return &Db{}, fmt.Errorf("%w: %v", errorapp.ErrInitDatabase, err)
	}

	return &Db{
		Pool: Pool,
	}, nil

}

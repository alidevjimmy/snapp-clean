package postgres

import (
	"context"
	"fmt"
	"github.com/alidevjimmy/snapp-clean/internal/config"
	"github.com/alidevjimmy/snapp-clean/internal/repository"
	"github.com/jackc/pgx/v4"
)

type postgres struct {
	db *pgx.Conn
	//logger logger.Logger
}

func New(cfg config.Postgres) (repository.Postgres, error) {
	db , err := pgx.Connect(context.Background(), dsn(cfg))
	if err != nil {
		// log
		panic(err)
	}
	defer db.Close(context.Background())
	return &postgres{
		db: db,
	}, nil
}

func dsn(cfg config.Postgres) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
}

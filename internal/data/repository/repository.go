package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Repository struct {
	UserRepo UserRepository
	DB       *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool, log *zap.Logger) Repository {
	return Repository{
		UserRepo: NewUserRepository(db),
		// ProductRepo:   NewProductRepository(db, log),
		// SaleRepo:      NewSaleRepository(db, log),
		// ReportRepo:    NewReportRepository(db),
		DB: db,
	}
}

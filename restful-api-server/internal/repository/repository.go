//go:generate mockgen -source=repository.go -destination=repository_mock.go -package=repository
package repository

import (
	"database/sql"
	"fmt"
	"log"
	"restful-api-server/internal/config"
	"restful-api-server/internal/models"

	_ "github.com/lib/pq"
)

type DBRepository interface {
	GetProducts() (models.Products, error)
}

func NewDBRepository(config *config.Config) DBRepository {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.Database.Host, config.Database.Port, config.Database.User, config.Database.Password, config.Database.Name)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return &PostgreSQL{db: db}
}

type PostgreSQL struct {
	db *sql.DB
}

func (p PostgreSQL) GetProducts() (models.Products, error) {
	return models.Products{
		{ID: 1, Name: "Product 1", Price: 100},
		{ID: 2, Name: "Product 2", Price: 200},
	}, nil
}

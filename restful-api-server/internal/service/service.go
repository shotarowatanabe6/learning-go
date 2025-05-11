//go:generate mockgen -source=service.go -destination=service_mock.go -package=service
package service

import (
	"restful-api-server/internal/config"
	"restful-api-server/internal/models"
	"restful-api-server/internal/repository"
)

type Service struct {
	DBRepository repository.DBRepository
}

func NewService(config *config.Config) *Service {
	return &Service{
		DBRepository: repository.NewDBRepository(config),
	}
}

type ProductService interface {
	GetProducts() (models.Products, error)
}

func (p *Service) GetProducts() (models.Products, error) {
	return p.DBRepository.GetProducts()
}

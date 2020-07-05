package news

import (
	"bareksa/models"
)

type INewsUsecase interface {
	Create(models.News) error
	GetAll(models.NewsFilterParams) ([]*models.News, *models.Pagination, error)
	Get(string) (*models.News, error)
	Update(string, models.News) error
	Delete(string) error
	Publish(string) error
}

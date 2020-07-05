package news

import (
	"bareksa/models"
)

type INewsRepository interface {
	Create(*models.News) error
	Update(*models.News, *models.News) error
	Delete(*models.News) error
	GetAll(models.NewsFilterParams) ([]*models.News, int, error)
	GetByID(int) (*models.News, error)
	GetBySlug(string) (*models.News, error)
	CountSlug(string) (models.CheckSlug, error)
	Publish(*models.News) error
}

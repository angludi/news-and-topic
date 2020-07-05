package topic

import (
	"bareksa/models"
)

type ITopicUsecase interface {
	GetAll(models.TopicFilterParameters) ([]*models.Topic, *models.Pagination, error)
	Get(string) (*models.Topic, error)
	Create(models.Topic) error
	Update(string, models.Topic) error
	Delete(string) error
	CountSlug(string) (*models.CheckSlug, error)
}

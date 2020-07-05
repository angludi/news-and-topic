package topic

import "bareksa/models"

type ITopicRepository interface {
	GetAll(models.TopicFilterParameters) ([]*models.Topic, int, error)
	GetByID(int) (*models.Topic, error)
	GetBySlug(string) (*models.Topic, error)
	Update(*models.Topic, *models.Topic) error
	Create(models.Topic) error
	Delete(models.Topic) error

	CountSlug(string) (models.CheckSlug, error)
}

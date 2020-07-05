package usecase

import (
	TopicInterface "bareksa/app/topic"
	"bareksa/models"
	"strconv"

	"github.com/asaskevich/govalidator"

	"github.com/pkg/errors"
)

type TopicUsecase struct {
	TopicRepository TopicInterface.ITopicRepository
}

func NewTopicUsecase(r TopicInterface.ITopicRepository) TopicInterface.ITopicUsecase {
	return &TopicUsecase{
		TopicRepository: r,
	}
}

func (u TopicUsecase) GetAll(params models.TopicFilterParameters) (topics []*models.Topic, pagination *models.Pagination, err error) {
	topics, total, err := u.TopicRepository.GetAll(params)
	if err != nil {
		return nil, nil, err
	}

	count := len(topics)

	pagination = models.BuildPagination(total, params.CurrentPage, params.PerPage, count)

	return topics, pagination, nil
}

func (u TopicUsecase) Get(identifier string) (topic *models.Topic, err error) {
	if govalidator.IsNumeric(identifier) {
		ID, _ := strconv.Atoi(identifier)
		topic, err = u.TopicRepository.GetByID(ID)
	} else {
		slug := identifier
		topic, err = u.TopicRepository.GetBySlug(slug)
	}

	if err != nil {
		return nil, err
	}

	if topic.ID == 0 {
		return nil, errors.New("record not found")
	}

	return topic, nil
}

func (u TopicUsecase) Create(topic models.Topic) (err error) {
	err = u.TopicRepository.Create(topic)
	if err != nil {
		return err
	}

	return nil
}

func (u TopicUsecase) Update(identifier string, req models.Topic) (err error) {
	var topic *models.Topic

	if govalidator.IsNumeric(identifier) {
		ID, _ := strconv.Atoi(identifier)
		topic, err = u.TopicRepository.GetByID(ID)
	} else {
		slug := identifier
		topic, err = u.TopicRepository.GetBySlug(slug)
	}
	if err != nil {
		return err
	}

	err = u.TopicRepository.Update(&req, topic)
	if err != nil {
		return err
	}

	return nil
}

func (u TopicUsecase) Delete(identifier string) (err error) {
	var topic *models.Topic

	if govalidator.IsNumeric(identifier) {
		ID, _ := strconv.Atoi(identifier)
		topic, err = u.TopicRepository.GetByID(ID)
	} else {
		slug := identifier
		topic, err = u.TopicRepository.GetBySlug(slug)
	}

	if err != nil {
		return err
	}

	err = u.TopicRepository.Delete(*topic)

	return nil
}

func (u TopicUsecase) CountSlug(slug string) (*models.CheckSlug, error) {
	check, err := u.TopicRepository.CountSlug(slug)
	if err != nil {
		return nil, err
	}

	return &check, nil
}

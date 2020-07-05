package repository

import (
	TopicInterface "bareksa/app/topic"
	"bareksa/models"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type TopicRepository struct {
	Conn *gorm.DB
}

func NewTopicRepository(Conn *gorm.DB) TopicInterface.ITopicRepository {
	return &TopicRepository{Conn}
}

func (r *TopicRepository) GetAll(params models.TopicFilterParameters) ([]*models.Topic, int, error) {
	var (
		topics []*models.Topic
		total  int
		err    error
	)

	tx := r.Conn

	txCount := tx
	txCount.Model(&topics).Count(&total)

	if err = tx.Find(&topics).Error; err != nil {
		return nil, total, err
	}

	return topics, total, err
}

func (r *TopicRepository) GetByID(id int) (*models.Topic, error) {
	var (
		topic models.Topic
		err   error
	)

	tx := r.Conn

	if err = tx.Where("id = ?", id).Find(&topic).Error; err != nil {
		return nil, err
	}

	return &topic, nil
}

func (r *TopicRepository) GetBySlug(slug string) (*models.Topic, error) {
	var (
		topic models.Topic
		err   error
	)

	tx := r.Conn

	if err = tx.Where("slug = ?", slug).First(&topic).Error; err != nil {
		return nil, err
	}

	return &topic, nil
}

func (r *TopicRepository) Create(t models.Topic) (err error) {
	topic := models.Topic{
		Name: t.Name,
		Slug: t.Slug,
	}
	tx := r.Conn.Begin()

	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()
	if err = tx.Create(&topic).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (r *TopicRepository) Update(data *models.Topic, t *models.Topic) (err error) {
	dt := map[string]interface{}{
		"name":       data.Name,
		"slug":       data.Slug,
		"updated_at": time.Now(),
	}

	tx := r.Conn.Begin()
	tx = tx.Where("id = ?", t.ID)
	if err = tx.Model(&t).UpdateColumns(dt).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (r *TopicRepository) Delete(t models.Topic) (err error) {
	tx := r.Conn.Begin()
	if err := tx.Delete(&t).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func (r *TopicRepository) CountSlug(slug string) (index models.CheckSlug, err error) {
	tx := r.Conn
	tx = tx.Select("COUNT(*) AS count, MAX(CAST(SUBSTRING(slug, -1, length(slug)) AS UNSIGNED)) AS max")
	if err := tx.Table("topics").Where("slug LIKE ?", slug+"%").Scan(&index).Error; err != nil {
		fmt.Println("Error:", err.Error())
		return index, err
	}

	return index, nil
}

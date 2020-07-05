package repository

import (
	NewsInterface "bareksa/app/news"
	"bareksa/models"
	"fmt"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
)

type NewsRepository struct {
	Conn *gorm.DB
}

func NewNewsRepository(Conn *gorm.DB) NewsInterface.INewsRepository {
	return &NewsRepository{Conn}
}

func (r *NewsRepository) GetAll(params models.NewsFilterParams) ([]*models.News, int, error) {
	var (
		news  []*models.News
		total int
		err   error
	)

	fmt.Printf("Params: %v:", params)

	tx := r.Conn
	tx = tx.Preload("Topic")

	if len(params.Topic) > 0 {
		tx = tx.Where("topic_id IN (?)", params.Topic)
	}

	if len(params.Title) > 0 {
		tx = tx.Where("title LIKE ?", "%"+params.Title+"%")
	}

	if len(params.Tags) > 0 {
		for i, tag := range params.Tags {
			if i == 0 {
				tx = tx.Where("tags LIKE ?", "%"+tag+"%")
			} else {
				tx = tx.Or("tags LIKE ?", "%"+tag+"%")
			}
		}
	}

	if len(params.Published) > 0 {
		isPublish, _ := strconv.ParseBool(params.Published)
		tx = tx.Where("is_publish = ?", isPublish)
	}

	if len(params.Deleted) > 0 {
		isDeleted, _ := strconv.ParseBool(params.Deleted)
		if isDeleted {
			tx = tx.Unscoped().Where("deleted_at IS NOT NULL")
		}

	}

	txCount := tx
	txCount.Model(&news).Count(&total)

	if err = tx.Limit(params.PerPage).Offset(params.Offset).Find(&news).Error; err != nil {
		return nil, total, err
	}

	return news, total, err
}

func (r NewsRepository) GetByID(ID int) (*models.News, error) {
	var (
		news models.News
		err  error
	)
	tx := r.Conn

	if err = tx.Where("id = ?", ID).Preload("Topic").First(&news).Error; err != nil {
		return nil, err
	}
	return &news, nil
}

func (r NewsRepository) GetBySlug(slug string) (*models.News, error) {
	var (
		news models.News
		err  error
	)

	tx := r.Conn

	if err = tx.Where("slug = ?", slug).Preload("Topic").First(&news).Error; err != nil {
		fmt.Println("Ini error Pak")
		return nil, err
	}
	return &news, nil
}

func (r NewsRepository) Create(n *models.News) (err error) {
	tx := r.Conn.Begin()

	n.CreatedAt = time.Now()
	n.UpdatedAt = time.Now()

	if err = tx.Create(&n).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (r NewsRepository) Update(data *models.News, n *models.News) error {
	dt := map[string]interface{}{
		"topic_id":   data.TopicID,
		"title":      data.Title,
		"slug":       data.Slug,
		"tags":       data.Tags,
		"is_publish": data.IsPublish,
		"updated_at": time.Now(),
	}

	tx := r.Conn.Begin()
	tx = tx.Where("id = ?", n.ID)
	if err := tx.Model(&n).UpdateColumns(dt).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (r NewsRepository) Publish(n *models.News) error {
	dt := map[string]interface{}{
		"is_publish": !n.IsPublish,
	}

	tx := r.Conn.Begin()
	tx = tx.Where("id = ?", n.ID)
	if err := tx.Model(&n).UpdateColumns(dt).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (r NewsRepository) Delete(n *models.News) error {
	tx := r.Conn.Begin()
	if err := tx.Delete(&n).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func (r *NewsRepository) CountSlug(slug string) (index models.CheckSlug, err error) {
	tx := r.Conn
	tx = tx.Select("COUNT(*) AS count, MAX(CAST(SUBSTRING(slug, -1, length(slug)) AS UNSIGNED)) AS max")
	if err := tx.Table("news").Where("slug LIKE ?", slug+"%").Scan(&index).Error; err != nil {
		return index, err
	}

	return index, nil
}

package transformers

import (
	"bareksa/models"
	"strings"
)

type NewsTopic struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type News struct {
	ID          int       `json:"id"`
	Topic       NewsTopic `json:"topic"`
	Title       string    `json:"title"`
	Slug        string    `json:"slug"`
	Description string    `json:"description"`
	Tags        []string  `json:"tags"`
	IsPublish   bool      `json:"is_publish"`
	CreatedAt   string    `json:"created_at"`
	UpdatedAt   string    `json:"updated_at"`
	DeletedAt   string    `json:"deleted_at"`
}

func (res *Transformer) NewsTransformer(data *models.News) *Transformer {
	res.Data = assignNews(data)
	return res
}

func (res *CollectionTransformer) NewsCollection(datas []*models.News, pagination *models.Pagination) {
	for _, data := range datas {
		res.Data = append(res.Data, assignNews(data))
	}

	if len(datas) == 0 {
		data := make([]interface{}, 0)
		res.Data = data
	}

	res.Meta = models.Meta{Pagination: pagination}
}

func assignNews(data *models.News) interface{} {
	result := News{}
	result.ID = data.ID
	result.Topic.ID = data.Topic.ID
	result.Topic.Name = data.Topic.Name
	result.Title = data.Title
	result.Slug = data.Slug
	result.Description = data.Description

	result.Tags = strings.Split(data.Tags, ",")
	result.IsPublish = data.IsPublish
	result.CreatedAt = data.CreatedAt.Format("2006-01-02 15:04:05")
	result.UpdatedAt = data.UpdatedAt.Format("2006-01-02 15:04:05")
	if data.DeletedAt.Valid {
		result.DeletedAt = data.DeletedAt.Time.Format("2006-01-02 15:04:05")
	}

	return result
}

package transformers

import (
	"bareksa/models"
)

type Topic struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (res *Transformer) TopicTransformer(data *models.Topic) *Transformer {
	res.Data = assignTopic(data)
	return res
}

func (res *CollectionTransformer) TopicCollection(datas []*models.Topic, pagination *models.Pagination) {
	for _, data := range datas {
		res.Data = append(res.Data, assignTopic(data))
	}

	if len(datas) == 0 {
		data := make([]interface{}, 0)
		res.Data = data
	}

	res.Meta = models.Meta{Pagination: pagination}
}

func assignTopic(data *models.Topic) interface{} {
	result := Topic{}
	result.ID = data.ID
	result.Name = data.Name
	result.Slug = data.Slug
	result.CreatedAt = data.CreatedAt.Format("2006-01-02 15:04:05")
	result.UpdatedAt = data.UpdatedAt.Format("2006-01-02 15:04:05")

	return result
}

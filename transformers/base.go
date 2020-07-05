package transformers

import (
	"bareksa/models"
)

type Transformer struct {
	Data interface{} `json:"data"`
}

type CollectionTransformer struct {
	Data []interface{} `json:"data"`
	Meta models.Meta   `json:"meta"`
}

package handler

import (
	Base "bareksa/app/api/handler"
	TopicInterface "bareksa/app/topic"
	"bareksa/helpers"
	"bareksa/models"
	"bareksa/transformers"
	"net/http"
	"strconv"

	"github.com/gosimple/slug"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type TopicRules struct {
	Name string `valid:"required~name required"`
}

type TopicHandler struct {
	TopicUsecase TopicInterface.ITopicUsecase
}

func (h TopicHandler) CreateTopic(c *gin.Context) {
	validation := ValidateTopic(c)
	if validation != nil {
		Base.RespondFailValidation(c, validation)
		return
	}

	topic := new(models.Topic)

	topic.Name = c.PostForm("name")

	slug := slug.Make(c.PostForm("name"))
	check, _ := h.TopicUsecase.CountSlug(slug)
	if check.Count > 0 {
		Base.RespondError(c, "Topic already exist", http.StatusBadRequest)
		return

	}
	topic.Slug = slug

	err := h.TopicUsecase.Create(*topic)
	if err != nil {
		Base.RespondFailValidation(c, err.Error())
		return
	}

	Base.RespondCreated(c, "Resource Created")
	return
}

func (h TopicHandler) GetTopic(c *gin.Context) {
	topic, err := h.TopicUsecase.Get(c.Param("identifier"))
	if err != nil {
		Base.RespondFailValidation(c, err.Error())
	}

	res := new(transformers.Transformer)
	res.TopicTransformer(topic)

	Base.RespondJSON(c, res)
	return
}

func (h TopicHandler) GetAllTopics(c *gin.Context) {
	var params models.TopicFilterParameters

	page, _ := strconv.Atoi(c.Query("page"))
	perPage, _ := strconv.Atoi(c.Query("perpage"))

	if page < 1 {
		page = 1
	}

	if perPage < 1 {
		perPage = 1000
	}

	offset := (page * perPage) - perPage

	params.CurrentPage = page
	params.PerPage = perPage
	params.Offset = offset

	topics, pagination, err := h.TopicUsecase.GetAll(params)
	if err != nil {
		Base.RespondFailValidation(c, err.Error())
	}

	res := new(transformers.CollectionTransformer)
	res.TopicCollection(topics, pagination)
	Base.RespondJSON(c, res)
	return
}

func (h TopicHandler) DeleteTopic(c *gin.Context) {
	err := h.TopicUsecase.Delete(c.Param("identifier"))
	if err != nil {
		Base.RespondFailValidation(c, err.Error())
		return
	}

	Base.RespondDeleted(c, "Resource Deleted")
	return
}

func (h TopicHandler) UpdateTopic(c *gin.Context) {
	validation := ValidateTopic(c)
	if validation != nil {
		Base.RespondFailValidation(c, validation)
		return
	}

	req := new(models.Topic)
	req.Name = c.PostForm("name")

	slug := slug.Make(c.PostForm("name"))
	check, _ := h.TopicUsecase.CountSlug(slug)
	if check.Count > 0 {
		Base.RespondError(c, "Topic already exist", http.StatusBadRequest)

	}
	req.Slug = slug

	err := h.TopicUsecase.Update(c.Param("identifier"), *req)
	if err != nil {
		Base.RespondFailValidation(c, err.Error())
		return
	}

	Base.RespondUpdated(c, "Resource Updated")
	return
}

func ValidateTopic(c *gin.Context) interface{} {
	rules := &TopicRules{
		Name: c.PostForm("name"),
	}

	_, err := govalidator.ValidateStruct(rules)
	if err != nil {
		respErr := helpers.ErrorValidation(rules, err)

		if respErr != nil {
			return respErr
		}
	}

	return nil
}

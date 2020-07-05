package handler

import (
	Base "bareksa/app/api/handler"
	NewsInterface "bareksa/app/news"
	"bareksa/helpers"
	"bareksa/models"
	"bareksa/transformers"
	"strconv"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type NewsRules struct {
	Topic       string `valid:"required~topic required"`
	Title       string `valid:"required~title required"`
	Description string `valid:"required~description required"`
	Tags        string `valid:"required~tags required"`
}

type NewsHandler struct {
	NewsUsecase NewsInterface.INewsUsecase
}

func (h NewsHandler) GetAllNews(c *gin.Context) {
	var params models.NewsFilterParams
	page, _ := strconv.Atoi(c.Query("page"))
	perPage, _ := strconv.Atoi(c.Query("perpage"))

	if page < 1 {
		page = 1
	}

	if perPage < 1 {
		perPage = 25
	}

	offset := (page * perPage) - perPage

	params.Title = c.Query("title")
	params.Topic = getSliceTopic(c.Query("topic_id"))
	params.Tags = strings.Split(c.Query("tags"), ",")
	params.Published = c.Query("published")
	params.Deleted = c.Query("deleted")
	params.PerPage = perPage
	params.Offset = offset
	params.CurrentPage = page

	news, pagination, err := h.NewsUsecase.GetAll(params)
	if err != nil {
		Base.RespondFailValidation(c, err.Error())
		return
	}

	res := new(transformers.CollectionTransformer)
	res.NewsCollection(news, pagination)

	Base.RespondJSON(c, res)
	return
}

func (h NewsHandler) GetNews(c *gin.Context) {
	news, err := h.NewsUsecase.Get(c.Param("identifier"))
	if err != nil {
		Base.RespondFailValidation(c, err.Error())
		return
	}

	res := new(transformers.Transformer)
	res.NewsTransformer(news)

	Base.RespondJSON(c, res)
	return
}

func (h NewsHandler) CreateNews(c *gin.Context) {
	validation := ValidateNews(c)
	if validation != nil {
		Base.RespondFailValidation(c, validation)
		return
	}

	req := h.MapNews(c)

	err := h.NewsUsecase.Create(*req)
	if err != nil {
		Base.RespondFailValidation(c, err.Error())
		return
	}

	Base.RespondCreated(c, "News Created")
	return
}

func (h NewsHandler) UpdateNews(c *gin.Context) {
	validation := ValidateNews(c)
	if validation != nil {
		Base.RespondFailValidation(c, validation)
		return
	}

	req := h.MapNews(c)
	err := h.NewsUsecase.Update(c.Param("identifier"), *req)
	if err != nil {
		Base.RespondFailValidation(c, err.Error())
		return
	}

	Base.RespondUpdated(c, "News Updated")
	return
}

func (h NewsHandler) DeleteNews(c *gin.Context) {

	err := h.NewsUsecase.Delete(c.Param("identifier"))
	if err != nil {
		Base.RespondFailValidation(c, err.Error())
		return
	}

	Base.RespondDeleted(c, "News Deleted")
	return
}

func (h NewsHandler) PublishNews(c *gin.Context) {
	err := h.NewsUsecase.Publish(c.Param("identifier"))
	if err != nil {
		Base.RespondFailValidation(c, err.Error())
		return
	}

	Base.RespondUpdated(c, "News Published")
	return
}

func ValidateNews(c *gin.Context) interface{} {
	rules := &NewsRules{
		Topic:       c.PostForm("topic_id"),
		Title:       c.PostForm("title"),
		Description: c.PostForm("description"),
		Tags:        c.PostForm("tags"),
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

func (h NewsHandler) MapNews(c *gin.Context) *models.News {
	news := new(models.News)

	news.TopicID, _ = strconv.Atoi(c.PostForm("topic_id"))
	news.Title = c.PostForm("title")
	news.Description = c.PostForm("description")
	news.Tags = c.PostForm("tags")

	return news
}

func getSliceTopic(strTopicID string) []int {
	var topicID []int
	arrTopicID := strings.Split(strTopicID, ",")
	for _, sid := range arrTopicID {
		id, err := strconv.Atoi(sid)
		if err == nil {
			topicID = append(topicID, id)
		}
	}

	return topicID
}

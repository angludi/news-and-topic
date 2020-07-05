package usecase

import (
	NewsInterface "bareksa/app/news"
	"bareksa/models"
	"fmt"
	"strconv"

	"github.com/asaskevich/govalidator"

	"github.com/gosimple/slug"
)

type NewsUsecase struct {
	NewsRepository NewsInterface.INewsRepository
}

func NewNewsUsecase(r NewsInterface.INewsRepository) NewsInterface.INewsUsecase {
	return &NewsUsecase{
		NewsRepository: r,
	}
}

// models.News
func (u NewsUsecase) Create(req models.News) (err error) {
	req.Slug = u.getSlug(req.Title)

	err = u.NewsRepository.Create(&req)
	if err != nil {
		return err
	}

	return nil
}

func (u NewsUsecase) GetAll(params models.NewsFilterParams) (news []*models.News, pagination *models.Pagination, err error) {
	news, total, err := u.NewsRepository.GetAll(params)
	if err != nil {
		return nil, nil, err
	}

	count := len(news)

	pagination = models.BuildPagination(total, params.CurrentPage, params.PerPage, count)

	return news, pagination, nil
}

func (u NewsUsecase) Get(identifier string) (news *models.News, err error) {
	if govalidator.IsNumeric(identifier) {
		ID, _ := strconv.Atoi(identifier)
		news, err = u.NewsRepository.GetByID(ID)
	} else {
		slug := identifier
		news, err = u.NewsRepository.GetBySlug(slug)
	}

	if err != nil {
		return nil, err
	}

	return news, nil
}

func (u NewsUsecase) Update(identifier string, req models.News) (err error) {
	var news *models.News
	if govalidator.IsNumeric(identifier) {
		ID, _ := strconv.Atoi(identifier)
		news, err = u.NewsRepository.GetByID(ID)
	} else {
		slug := identifier
		news, err = u.NewsRepository.GetBySlug(slug)
	}

	req.Slug = u.getSlug(req.Title)

	if req.Title == news.Title {
		req.Slug = news.Slug
	}

	err = u.NewsRepository.Update(&req, news)
	if err != nil {
		return err
	}
	return nil
}

func (u NewsUsecase) Delete(identifier string) (err error) {
	var news *models.News
	if govalidator.IsNumeric(identifier) {
		ID, _ := strconv.Atoi(identifier)
		news, err = u.NewsRepository.GetByID(ID)
	} else {
		slug := identifier
		news, err = u.NewsRepository.GetBySlug(slug)
	}

	err = u.NewsRepository.Delete(news)
	if err != nil {
		return err
	}
	return nil
}

func (u NewsUsecase) Publish(identifier string) (err error) {
	var news *models.News
	if govalidator.IsNumeric(identifier) {
		ID, _ := strconv.Atoi(identifier)
		news, err = u.NewsRepository.GetByID(ID)
	} else {
		slug := identifier
		news, err = u.NewsRepository.GetBySlug(slug)
	}

	err = u.NewsRepository.Publish(news)
	if err != nil {
		return err
	}

	return nil
}

func (u NewsUsecase) getSlug(text string) string {
	slug := slug.Make(text)
	check, _ := u.NewsRepository.CountSlug(slug)
	if check.Count > 0 {
		return fmt.Sprintf("%s-%d", slug, check.Max+1)
	}

	return slug
}

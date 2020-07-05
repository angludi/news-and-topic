package models

import "math"

type Meta struct {
	Pagination *Pagination `json:"pagination"`
}

type Pagination struct {
	Total       int `json:"total"`
	Count       int `json:"count"`
	PerPage     int `json:"per_page"`
	CurrentPage int `json:"current_page"`
	TotalPages  int `json:"total_pages"`
}

func BuildPagination(total, page, perPage, count int) *Pagination {
	pagination := new(Pagination)
	pagination.Total = total
	pagination.Count = count
	pagination.PerPage = perPage
	pagination.CurrentPage = page
	pagination.TotalPages = int(math.Ceil(float64(total) / float64(perPage)))

	return pagination
}

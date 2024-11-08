package models

type PaginationFilter struct {
	Offset *uint `json:"offset"`
	Limit  *uint `json:"limit"`
}

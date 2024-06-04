package common

type ListRes struct {
	CurrentPage int `json:"currentPage,omitempty"`
	Total       int `json:"total"`
}

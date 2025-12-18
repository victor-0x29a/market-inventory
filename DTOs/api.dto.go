package dtos

type ApiError struct {
	Errors []string `json:"Errors"`
}

type ApiPagination struct {
	Page    int
	PerPage int
}

type ApiPaginationResponse struct {
	Records    any   `json:"Records"`
	TotalPages int64 `json:"TotalPages"`
	ItemsCount int64 `json:"ItemsCount"`
}

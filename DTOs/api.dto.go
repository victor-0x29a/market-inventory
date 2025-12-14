package dtos

type ApiError struct {
	Errors []string `json:"errors"`
}

type ApiPagination struct {
	Page    int
	PerPage int
}

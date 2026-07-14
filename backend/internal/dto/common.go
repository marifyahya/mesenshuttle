package dto

type PaginatedResponse struct {
	Data       interface{} `json:"data"`
	TotalCount int64       `json:"total_count"`
	Page       int         `json:"page"`
	Limit      int         `json:"limit"`
}

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

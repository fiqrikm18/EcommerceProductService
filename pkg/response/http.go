package response

type PaginationResponse struct {
	Message   string      `json:"message"`
	TotalPage int         `json:"total_page"`
	PerPage   int         `json:"items_per_page"`
	Page      int         `json:"current_page"`
	Total     interface{} `json:"total_items"`
	Data      interface{} `json:"data"`
}

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"errors"`
}

func NewSuccessResponse(message string, data interface{}) *SuccessResponse {
	return &SuccessResponse{
		Message: message,
		Data:    data,
	}
}

func NewErrorResponse(message string, error ...interface{}) *ErrorResponse {
	return &ErrorResponse{
		Message: message,
		Data:    error,
	}
}

func NewPaginationResponse(total int, totalPage int, perPage int, pageSize int, data interface{}) *PaginationResponse {
	return &PaginationResponse{
		Message:   "success",
		TotalPage: totalPage,
		PerPage:   perPage,
		Page:      pageSize,
		Total:     total,
		Data:      data,
	}
}

package http

type CreateSubRequest struct {
	UserID      string  `json:"user_id" example:"60601fee-2bf1-4721-ae6f-7636e79a0cba"`
	ServiceName string  `json:"service_name" example:"Yandex Plus"`
	Price       int64   `json:"price" example:"400"`
	StartDate   string  `json:"start_date" example:"07-2025"`
	EndDate     *string `json:"end_date,omitempty" example:"12-2025"`
}

type CreateSubResponse struct {
	ID string `json:"id" example:"60601fee-2bf1-4721-ae6f-7636e79a0cba"`
}

type UpdateSubRequest struct {
	ServiceName *string `json:"service_name,omitempty" example:"Netflix"`
	Price       *int64  `json:"price,omitempty" example:"500"`
	StartDate   *string `json:"start_date,omitempty" example:"07-2025"`
	EndDate     *string `json:"end_date,omitempty" example:"12-2025"`
}

type CalculateTotalCostResponse struct {
	Total       int64   `json:"total" example:"500"`
	FromDate    string  `json:"from_date,omitempty" example:"07-2025"`
	ToDate      string  `json:"to_date,omitempty" example:"12-2025"`
	UserID      *string `json:"user_id,omitempty" example:"60601fee-2bf1-4721-ae6f-7636e79a0cba"`
	ServiceName *string `json:"service_name,omitempty" example:"Netflix"`
}

type ErrorResponse struct {
	Error string `json:"error" example:"error"`
}

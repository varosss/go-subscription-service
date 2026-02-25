package dto

type Subscription struct {
	ID          string  `json:"id" example:"60601fee-2bf1-4721-ae6f-7636e79a0cba"`
	UserID      string  `json:"user_id" example:"60601fee-2bf1-4721-ae6f-7636e79a0cba"`
	ServiceName string  `json:"service_name" example:"Netflix"`
	Price       int64   `json:"price" example:"500"`
	StartDate   string  `json:"start_date" example:"02-2026"`
	EndDate     *string `json:"end_date,omitempty" example:"05-2026"`
}

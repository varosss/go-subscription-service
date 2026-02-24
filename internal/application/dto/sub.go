package dto

type Subscription struct {
	ID          string
	UserID      string
	ServiceName string
	Price       int64
	StartDate   string
	EndDate     *string
}

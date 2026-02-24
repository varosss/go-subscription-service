package dto

import "time"

type Subscription struct {
	ID          string
	UserID      string
	ServiceName string
	Price       string
	StartDate   time.Time
	EndDate     *time.Time
}

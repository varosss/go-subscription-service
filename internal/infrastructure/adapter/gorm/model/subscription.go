package model

import (
	"time"
)

type Subscription struct {
	ID          string     `gorm:"tprimaryKey" json:"id"`
	ServiceName string     `gorm:"not null" json:"service_name"`
	Price       int64      `gorm:"not null" json:"price"`
	UserID      string     `gorm:"not null" json:"user_id"`
	StartDate   time.Time  `gorm:"not null" json:"start_date"`
	EndDate     *time.Time `json:"end_date,omitempty"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

package entity

import (
	"go-subscription-service/internal/domain/valueobject"
	"time"
)

type Subscription struct {
	id          valueobject.SubscriptionID
	userID      valueobject.UserID
	serviceName string
	price       int64
	startDate   time.Time
	endDate     *time.Time
}

func NewSubscription(
	userID valueobject.UserID,
	serviceName string,
	price int64,
	startDate time.Time,
	endDate *time.Time,
) *Subscription {
	return &Subscription{
		id:          valueobject.NewSubscriptionID(),
		userID:      userID,
		serviceName: serviceName,
		price:       price,
		startDate:   startDate,
		endDate:     endDate,
	}
}

func SubscribtionFromPrimitives(
	id valueobject.SubscriptionID,
	userID valueobject.UserID,
	serviceName string,
	price int64,
	startDate time.Time,
	endDate *time.Time,
) *Subscription {
	return &Subscription{
		id:          id,
		userID:      userID,
		serviceName: serviceName,
		price:       price,
		startDate:   startDate,
		endDate:     endDate,
	}
}

func (s *Subscription) ID() valueobject.SubscriptionID {
	return s.id
}

func (s *Subscription) UserID() valueobject.UserID {
	return s.userID
}

func (s *Subscription) ServiceName() string {
	return s.serviceName
}

func (s *Subscription) Price() int64 {
	return s.price
}

func (s *Subscription) StartDate() time.Time {
	return s.startDate
}

func (s *Subscription) EndDate() *time.Time {
	return s.endDate
}

func (s *Subscription) SetServiceName(serviceName string) {
	s.serviceName = serviceName
}

func (s *Subscription) SetPrice(price int64) {
	s.price = price
}

func (s *Subscription) SetStartDate(startDate time.Time) {
	s.startDate = startDate
}

func (s *Subscription) SetEndDate(endDate time.Time) {
	s.endDate = &endDate
}

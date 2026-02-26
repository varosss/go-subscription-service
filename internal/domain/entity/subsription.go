package entity

import (
	"go-subscription-service/internal/domain"
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
) (*Subscription, error) {
	if serviceName == "" {
		return nil, domain.ErrInvalidServiceName
	}

	if price <= 0 {
		return nil, domain.ErrInvalidPrice
	}

	if endDate != nil && endDate.Before(startDate) {
		return nil, domain.ErrInvalidDateRange
	}

	return &Subscription{
		id:          valueobject.NewSubscriptionID(),
		userID:      userID,
		serviceName: serviceName,
		price:       price,
		startDate:   startDate,
		endDate:     endDate,
	}, nil
}

func SubscriptionFromPrimitives(
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

func (s *Subscription) ChangeServiceName(serviceName string) error {
	if serviceName == "" {
		return domain.ErrInvalidServiceName
	}
	s.serviceName = serviceName
	return nil
}

func (s *Subscription) ChangePrice(price int64) error {
	if price <= 0 {
		return domain.ErrInvalidPrice
	}
	s.price = price
	return nil
}

func (s *Subscription) ChangeStartDate(startDate time.Time) error {
	if s.endDate != nil && s.endDate.Before(startDate) {
		return domain.ErrInvalidDateRange
	}
	s.startDate = startDate
	return nil
}

func (s *Subscription) ChangeEndDate(endDate *time.Time) error {
	if endDate != nil && endDate.Before(s.startDate) {
		return domain.ErrInvalidDateRange
	}
	s.endDate = endDate
	return nil
}

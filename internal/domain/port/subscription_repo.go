package port

import (
	"context"
	"go-subscription-service/internal/domain/entity"
	"go-subscription-service/internal/domain/valueobject"
	"time"
)

type SubscriptionRepo interface {
	Save(ctx context.Context, subscription *entity.Subscription) error
	GetByID(ctx context.Context, subscriptionID valueobject.SubscriptionID) (*entity.Subscription, error)
	DeleteByID(ctx context.Context, subscriptionID valueobject.SubscriptionID) error
	List(
		ctx context.Context,
		userID *valueobject.UserID,
		serviceName *string,
		fromDate *time.Time,
		toDate *time.Time,
		limit *int,
		offset *int,
	) ([]*entity.Subscription, error)
	CalculateTotalCost(
		ctx context.Context,
		userID *valueobject.UserID,
		serviceName *string,
		fromDate time.Time,
		toDate time.Time,
	) (int64, error)
}

package usecase

import (
	"context"
	"go-subscription-service/internal/domain/entity"
	"go-subscription-service/internal/domain/port"
	"go-subscription-service/internal/domain/valueobject"
	"time"
)

type ListSubscriptionsCommand struct {
	UserID      *valueobject.UserID
	ServiceName *string
	FromDate    *time.Time
	ToDate      *time.Time
	Limit       *int
	Offset      *int
}

type ListSubscriptionsUseCase struct {
	subscriptions port.SubscriptionRepo
}

func NewListSubscriptionsUseCase(
	subscriptions port.SubscriptionRepo,
) *ListSubscriptionsUseCase {
	return &ListSubscriptionsUseCase{
		subscriptions: subscriptions,
	}
}

func (uc *ListSubscriptionsUseCase) Execute(ctx context.Context, cmd ListSubscriptionsCommand) ([]*entity.Subscription, error) {
	subs, err := uc.subscriptions.List(
		ctx,
		cmd.UserID,
		cmd.ServiceName,
		cmd.FromDate,
		cmd.ToDate,
		cmd.Limit,
		cmd.Offset,
	)
	if err != nil {
		return nil, err
	}

	return subs, err
}

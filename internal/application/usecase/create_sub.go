package usecase

import (
	"context"
	"go-subscription-service/internal/application/dto"
	"go-subscription-service/internal/domain/entity"
	"go-subscription-service/internal/domain/port"
	"go-subscription-service/internal/domain/valueobject"
	"time"
)

type CreateSubscriptionCommand struct {
	UserID      valueobject.UserID
	ServiceName string
	Price       int64
	StartDate   time.Time
	EndDate     *time.Time
}

type CreateSubscriptionResult struct {
	Subscription dto.Subscription
}

type CreateSubscriptionUseCase struct {
	subscriptions port.SubscriptionRepo
}

func NewCreateSubscriptionUseCase(
	subscriptions port.SubscriptionRepo,
) *CreateSubscriptionUseCase {
	return &CreateSubscriptionUseCase{
		subscriptions: subscriptions,
	}
}

func (uc *CreateSubscriptionUseCase) Execute(
	ctx context.Context,
	cmd CreateSubscriptionCommand,
) (*CreateSubscriptionResult, error) {
	subscription := entity.NewSubscription(
		cmd.UserID,
		cmd.ServiceName,
		cmd.Price,
		cmd.StartDate,
		cmd.EndDate,
	)

	err := uc.subscriptions.Save(ctx, subscription)
	if err != nil {
		return nil, err
	}

	return &CreateSubscriptionResult{
		Subscription: dto.Subscription{
			ID: subscription.ID().String(),
		},
	}, nil
}

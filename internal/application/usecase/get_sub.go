package usecase

import (
	"context"
	"go-subscription-service/internal/domain/entity"
	"go-subscription-service/internal/domain/port"
	"go-subscription-service/internal/domain/valueobject"
)

type GetSubscriptionCommand struct {
	SubscriptionID valueobject.SubscriptionID
}

type GetSubscriptionUseCase struct {
	subscriptions port.SubscriptionRepo
}

func NewGetSubscriptionUseCase(
	subscriptions port.SubscriptionRepo,
) *GetSubscriptionUseCase {
	return &GetSubscriptionUseCase{
		subscriptions: subscriptions,
	}
}

func (uc *GetSubscriptionUseCase) Execute(
	ctx context.Context,
	cmd GetSubscriptionCommand,
) (*entity.Subscription, error) {
	Subscription, err := uc.subscriptions.GetByID(ctx, cmd.SubscriptionID)
	if err != nil {
		return nil, err
	}

	return Subscription, nil
}

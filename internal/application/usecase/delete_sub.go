package usecase

import (
	"context"
	"go-subscription-service/internal/domain/port"
	"go-subscription-service/internal/domain/valueobject"
)

type DeleteSubscriptionCommand struct {
	SubscriptionID valueobject.SubscriptionID
}

type DeleteSubscriptionUseCase struct {
	subscriptions port.SubscriptionRepo
}

func NewDeleteSubscriptionUseCase(
	subscriptions port.SubscriptionRepo,
) *DeleteSubscriptionUseCase {
	return &DeleteSubscriptionUseCase{
		subscriptions: subscriptions,
	}
}

func (uc *DeleteSubscriptionUseCase) Execute(
	ctx context.Context,
	cmd DeleteSubscriptionCommand,
) error {
	if err := uc.subscriptions.DeleteByID(ctx, cmd.SubscriptionID); err != nil {
		return err
	}

	return nil
}

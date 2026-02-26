package usecase

import (
	"context"
	aport "go-subscription-service/internal/application/port"
	"go-subscription-service/internal/domain/port"
	"go-subscription-service/internal/domain/valueobject"
)

type DeleteSubscriptionCommand struct {
	SubscriptionID valueobject.SubscriptionID
}

type DeleteSubscriptionUseCase struct {
	subscriptions port.SubscriptionRepo
	logger        aport.Logger
}

func NewDeleteSubscriptionUseCase(
	subscriptions port.SubscriptionRepo,
	logger aport.Logger,
) *DeleteSubscriptionUseCase {
	return &DeleteSubscriptionUseCase{
		subscriptions: subscriptions,
		logger:        logger,
	}
}

func (uc *DeleteSubscriptionUseCase) Execute(
	ctx context.Context,
	cmd DeleteSubscriptionCommand,
) error {
	uc.logger.Debug(ctx, "starting delete subscription",
		aport.Field{Key: "subscription_id", Value: cmd.SubscriptionID.String()},
	)

	if err := uc.subscriptions.DeleteByID(ctx, cmd.SubscriptionID); err != nil {
		uc.logger.Error(ctx, "failed to delete subscription",
			aport.Field{Key: "subscription_id", Value: cmd.SubscriptionID.String()},
			aport.Field{Key: "error", Value: err.Error()},
		)

		return err
	}

	uc.logger.Info(ctx, "subscription deleted successfully",
		aport.Field{Key: "subscription_id", Value: cmd.SubscriptionID.String()},
	)

	return nil
}

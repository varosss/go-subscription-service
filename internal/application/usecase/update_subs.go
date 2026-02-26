package usecase

import (
	"context"
	aport "go-subscription-service/internal/application/port"
	"go-subscription-service/internal/domain/port"
	"go-subscription-service/internal/domain/valueobject"
	"time"
)

type UpdateSubscriptionCommand struct {
	SubscriptionID valueobject.SubscriptionID
	ServiceName    *string
	Price          *int64
	StartDate      *time.Time
	EndDate        *time.Time
}

type UpdateSubscriptionUseCase struct {
	subscriptions port.SubscriptionRepo
	logger        aport.Logger
}

func NewUpdateSubscriptionUseCase(
	subscriptions port.SubscriptionRepo,
	logger aport.Logger,
) *UpdateSubscriptionUseCase {
	return &UpdateSubscriptionUseCase{
		subscriptions: subscriptions,
		logger:        logger,
	}
}

func (uc *UpdateSubscriptionUseCase) Execute(
	ctx context.Context,
	cmd UpdateSubscriptionCommand,
) error {
	uc.logger.Debug(ctx, "starting update subscription",
		aport.Field{Key: "subscription_id", Value: cmd.SubscriptionID.String()},
	)

	subscribtion, err := uc.subscriptions.GetByID(ctx, cmd.SubscriptionID)
	if err != nil {
		uc.logger.Error(ctx, "failed to get a subscription",
			aport.Field{Key: "subscription_id", Value: cmd.SubscriptionID.String()},
			aport.Field{Key: "error", Value: err.Error()},
		)

		return err
	}

	if cmd.Price != nil {
		subscribtion.SetPrice(*cmd.Price)
	}

	if cmd.EndDate != nil {
		subscribtion.SetEndDate(*cmd.EndDate)
	}

	if cmd.StartDate != nil {
		subscribtion.SetStartDate(*cmd.StartDate)
	}

	if cmd.ServiceName != nil {
		subscribtion.SetServiceName(*cmd.ServiceName)
	}

	if err := uc.subscriptions.Save(ctx, subscribtion); err != nil {
		uc.logger.Error(ctx, "failed to save subscription",
			aport.Field{Key: "subscription_id", Value: cmd.SubscriptionID.String()},
			aport.Field{Key: "error", Value: err.Error()},
		)

		return err
	}

	uc.logger.Info(ctx, "subscription updated successfully",
		aport.Field{Key: "subscription_id", Value: cmd.SubscriptionID.String()},
	)

	return nil
}

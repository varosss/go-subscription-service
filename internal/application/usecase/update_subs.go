package usecase

import (
	"context"
	"go-subscription-service/internal/application"
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

	subscription, err := uc.subscriptions.GetByID(ctx, cmd.SubscriptionID)
	if err != nil {
		uc.logger.Error(ctx, "failed to get a subscription",
			aport.Field{Key: "subscription_id", Value: cmd.SubscriptionID.String()},
			aport.Field{Key: "error", Value: err.Error()},
		)

		return application.ErrSubscriptionNotFound
	}

	if cmd.Price != nil {
		if err := subscription.ChangePrice(*cmd.Price); err != nil {
			uc.logger.Error(ctx, "failed to change subscriptions price",
				aport.Field{Key: "subscription_id", Value: cmd.SubscriptionID.String()},
				aport.Field{Key: "price", Value: *cmd.Price},
				aport.Field{Key: "error", Value: err.Error()},
			)

			return application.ErrInvalidSubscriptionData
		}
	}

	if cmd.EndDate != nil {
		if err := subscription.ChangeEndDate(cmd.EndDate); err != nil {
			uc.logger.Error(ctx, "failed to change subscriptions end date",
				aport.Field{Key: "subscription_id", Value: cmd.SubscriptionID.String()},
				aport.Field{Key: "end_date", Value: cmd.EndDate},
				aport.Field{Key: "error", Value: err.Error()},
			)

			return application.ErrInvalidSubscriptionData
		}
	}

	if cmd.StartDate != nil {
		if err := subscription.ChangeStartDate(*cmd.StartDate); err != nil {
			uc.logger.Error(ctx, "failed to change subscriptions start date",
				aport.Field{Key: "subscription_id", Value: cmd.SubscriptionID.String()},
				aport.Field{Key: "start_date", Value: cmd.StartDate},
				aport.Field{Key: "error", Value: err.Error()},
			)

			return application.ErrInvalidSubscriptionData
		}
	}

	if cmd.ServiceName != nil {
		if err := subscription.ChangeServiceName(*cmd.ServiceName); err != nil {
			uc.logger.Error(ctx, "failed to change subscriptions service name",
				aport.Field{Key: "subscription_id", Value: cmd.SubscriptionID.String()},
				aport.Field{Key: "service_name", Value: *cmd.ServiceName},
				aport.Field{Key: "error", Value: err.Error()},
			)

			return application.ErrInvalidSubscriptionData
		}
	}

	if err := uc.subscriptions.Save(ctx, subscription); err != nil {
		uc.logger.Error(ctx, "failed to save subscription",
			aport.Field{Key: "subscription_id", Value: cmd.SubscriptionID.String()},
			aport.Field{Key: "error", Value: err.Error()},
		)

		return application.ErrSubscriptionSaveFailed
	}

	uc.logger.Info(ctx, "subscription updated successfully",
		aport.Field{Key: "subscription_id", Value: cmd.SubscriptionID.String()},
	)

	return nil
}

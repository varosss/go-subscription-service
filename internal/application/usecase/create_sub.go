package usecase

import (
	"context"
	"go-subscription-service/internal/application/dto"
	aport "go-subscription-service/internal/application/port"
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
	logger        aport.Logger
}

func NewCreateSubscriptionUseCase(
	subscriptions port.SubscriptionRepo,
	logger aport.Logger,
) *CreateSubscriptionUseCase {
	return &CreateSubscriptionUseCase{
		subscriptions: subscriptions,
		logger:        logger,
	}
}

func (uc *CreateSubscriptionUseCase) Execute(
	ctx context.Context,
	cmd CreateSubscriptionCommand,
) (*CreateSubscriptionResult, error) {
	uc.logger.Debug(ctx, "starting create subscription",
		aport.Field{Key: "user_id", Value: cmd.UserID.String()},
		aport.Field{Key: "service_name", Value: cmd.ServiceName},
	)

	subscription := entity.NewSubscription(
		cmd.UserID,
		cmd.ServiceName,
		cmd.Price,
		cmd.StartDate,
		cmd.EndDate,
	)

	err := uc.subscriptions.Save(ctx, subscription)
	if err != nil {
		uc.logger.Error(ctx, "failed to save subscription",
			aport.Field{Key: "user_id", Value: cmd.UserID.String()},
			aport.Field{Key: "service_name", Value: cmd.ServiceName},
			aport.Field{Key: "error", Value: err.Error()},
		)

		return nil, err
	}

	uc.logger.Info(ctx, "subscription created successfully",
		aport.Field{Key: "subscription_id", Value: subscription.ID().String()},
		aport.Field{Key: "user_id", Value: cmd.UserID.String()},
	)

	return &CreateSubscriptionResult{
		Subscription: dto.Subscription{
			ID: subscription.ID().String(),
		},
	}, nil
}

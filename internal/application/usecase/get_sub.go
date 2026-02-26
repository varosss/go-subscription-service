package usecase

import (
	"context"
	"go-subscription-service/internal/application/dto"
	aport "go-subscription-service/internal/application/port"
	"go-subscription-service/internal/domain/port"
	"go-subscription-service/internal/domain/valueobject"
)

type GetSubscriptionCommand struct {
	SubscriptionID valueobject.SubscriptionID
}

type GetSubscriptionResult struct {
	Subscription dto.Subscription
}

type GetSubscriptionUseCase struct {
	subscriptions port.SubscriptionRepo
	logger        aport.Logger
}

func NewGetSubscriptionUseCase(
	subscriptions port.SubscriptionRepo,
	logger aport.Logger,
) *GetSubscriptionUseCase {
	return &GetSubscriptionUseCase{
		subscriptions: subscriptions,
		logger:        logger,
	}
}

func (uc *GetSubscriptionUseCase) Execute(
	ctx context.Context,
	cmd GetSubscriptionCommand,
) (*GetSubscriptionResult, error) {
	uc.logger.Debug(ctx, "starting get subscription",
		aport.Field{Key: "subscription_id", Value: cmd.SubscriptionID.String()},
	)

	subscription, err := uc.subscriptions.GetByID(ctx, cmd.SubscriptionID)
	if err != nil {
		uc.logger.Error(ctx, "failed to get a subscription",
			aport.Field{Key: "subscription_id", Value: cmd.SubscriptionID.String()},
			aport.Field{Key: "error", Value: err.Error()},
		)

		return nil, err
	}

	var endDate *string
	if subscription.EndDate() != nil {
		formattedDate := subscription.EndDate().Format("01-2006")
		endDate = &formattedDate
	}

	subResult := dto.Subscription{
		ID:          subscription.ID().String(),
		UserID:      subscription.UserID().String(),
		ServiceName: subscription.ServiceName(),
		Price:       subscription.Price(),
		StartDate:   subscription.StartDate().Format("01-2006"),
		EndDate:     endDate,
	}

	uc.logger.Info(ctx, "got subscription successfully",
		aport.Field{Key: "subscription", Value: subResult},
	)

	return &GetSubscriptionResult{
		Subscription: subResult,
	}, nil
}

package usecase

import (
	"context"
	"go-subscription-service/internal/application/dto"
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
) (*GetSubscriptionResult, error) {
	subscription, err := uc.subscriptions.GetByID(ctx, cmd.SubscriptionID)
	if err != nil {
		return nil, err
	}

	var endDate *string
	if subscription.EndDate() != nil {
		formattedDate := subscription.EndDate().Format("01-2006")
		endDate = &formattedDate
	}

	return &GetSubscriptionResult{
		Subscription: dto.Subscription{
			ID:          subscription.ID().String(),
			UserID:      subscription.UserID().String(),
			ServiceName: subscription.ServiceName(),
			Price:       subscription.Price(),
			StartDate:   subscription.StartDate().Format("01-2006"),
			EndDate:     endDate,
		},
	}, nil
}

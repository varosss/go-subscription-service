package builder

import (
	aport "go-subscription-service/internal/application/port"
	"go-subscription-service/internal/application/usecase"
	"go-subscription-service/internal/domain/port"
)

type UseCases struct {
	CreateSub          *usecase.CreateSubscriptionUseCase
	UpdateSub          *usecase.UpdateSubscriptionUseCase
	GetSub             *usecase.GetSubscriptionUseCase
	DeleteSub          *usecase.DeleteSubscriptionUseCase
	ListSubs           *usecase.ListSubscriptionsUseCase
	CalculateTotalCost *usecase.CalculateTotalCostUseCase
}

func BuildUseCases(
	subscriptions port.SubscriptionRepo,
	logger aport.Logger,
) *UseCases {
	return &UseCases{
		CreateSub:          usecase.NewCreateSubscriptionUseCase(subscriptions, logger),
		UpdateSub:          usecase.NewUpdateSubscriptionUseCase(subscriptions, logger),
		GetSub:             usecase.NewGetSubscriptionUseCase(subscriptions, logger),
		DeleteSub:          usecase.NewDeleteSubscriptionUseCase(subscriptions, logger),
		ListSubs:           usecase.NewListSubscriptionsUseCase(subscriptions, logger),
		CalculateTotalCost: usecase.NewCalculateTotalCostUseCase(subscriptions, logger),
	}
}

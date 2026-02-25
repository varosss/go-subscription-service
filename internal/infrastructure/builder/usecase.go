package builder

import (
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
) *UseCases {
	return &UseCases{
		CreateSub:          usecase.NewCreateSubscriptionUseCase(subscriptions),
		UpdateSub:          usecase.NewUpdateSubscriptionUseCase(subscriptions),
		GetSub:             usecase.NewGetSubscriptionUseCase(subscriptions),
		DeleteSub:          usecase.NewDeleteSubscriptionUseCase(subscriptions),
		ListSubs:           usecase.NewListSubscriptionsUseCase(subscriptions),
		CalculateTotalCost: usecase.NewCalculateTotalCostUseCase(subscriptions),
	}
}

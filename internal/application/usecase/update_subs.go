package usecase

import (
	"context"
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
}

func NewUpdateSubscriptionUseCase(
	subscriptions port.SubscriptionRepo,
) *UpdateSubscriptionUseCase {
	return &UpdateSubscriptionUseCase{
		subscriptions: subscriptions,
	}
}

func (uc *UpdateSubscriptionUseCase) Execute(
	ctx context.Context,
	cmd UpdateSubscriptionCommand,
) error {
	subscribtion, err := uc.subscriptions.GetByID(ctx, cmd.SubscriptionID)
	if err != nil {
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
		return err
	}

	return nil
}

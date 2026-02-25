package usecase

import (
	"context"
	"go-subscription-service/internal/domain/port"
	"go-subscription-service/internal/domain/valueobject"
	"time"
)

type CalculateTotalCostCommand struct {
	UserID      *valueobject.UserID
	ServiceName *string
	FromDate    time.Time
	ToDate      time.Time
}

type CalculateTotalCostResult struct {
	Total int64
}

type CalculateTotalCostUseCase struct {
	subscriptions port.SubscriptionRepo
}

func NewCalculateTotalCostUseCase(subscriptions port.SubscriptionRepo) *CalculateTotalCostUseCase {
	return &CalculateTotalCostUseCase{
		subscriptions: subscriptions,
	}
}

func (uc *CalculateTotalCostUseCase) Execute(
	ctx context.Context,
	cmd CalculateTotalCostCommand,
) (*CalculateTotalCostResult, error) {
	total, err := uc.subscriptions.CalculateTotalCost(
		ctx,
		cmd.UserID,
		cmd.ServiceName,
		cmd.FromDate,
		cmd.ToDate,
	)
	if err != nil {
		return nil, err
	}

	return &CalculateTotalCostResult{
		Total: total,
	}, nil
}

package usecase

import (
	"context"
	aport "go-subscription-service/internal/application/port"
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
	logger        aport.Logger
}

func NewCalculateTotalCostUseCase(
	subscriptions port.SubscriptionRepo,
	logger aport.Logger,
) *CalculateTotalCostUseCase {
	return &CalculateTotalCostUseCase{
		subscriptions: subscriptions,
		logger:        logger,
	}
}

func (uc *CalculateTotalCostUseCase) Execute(
	ctx context.Context,
	cmd CalculateTotalCostCommand,
) (*CalculateTotalCostResult, error) {
	uc.logger.Debug(ctx, "starting calculate total cost",
		aport.Field{Key: "from_date", Value: cmd.FromDate.Format("01-2006")},
		aport.Field{Key: "to_date", Value: cmd.ToDate.Format("01-2006")},
	)

	total, err := uc.subscriptions.CalculateTotalCost(
		ctx,
		cmd.UserID,
		cmd.ServiceName,
		cmd.FromDate,
		cmd.ToDate,
	)
	if err != nil {
		uc.logger.Error(ctx, "failed to calculate total cost",
			aport.Field{Key: "error", Value: err.Error()},
		)

		return nil, err
	}

	uc.logger.Info(ctx, "total cost calculated successfully",
		aport.Field{Key: "total", Value: total},
		aport.Field{Key: "from_date", Value: cmd.FromDate.Format("01-2006")},
		aport.Field{Key: "to_date", Value: cmd.ToDate.Format("01-2006")},
	)

	return &CalculateTotalCostResult{
		Total: total,
	}, nil
}

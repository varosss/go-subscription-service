package usecase

import (
	"context"
	"go-subscription-service/internal/application"
	"go-subscription-service/internal/application/dto"
	aport "go-subscription-service/internal/application/port"
	"go-subscription-service/internal/domain/port"
	"go-subscription-service/internal/domain/valueobject"
	"time"
)

type ListSubscriptionsCommand struct {
	UserID      *valueobject.UserID
	ServiceName *string
	FromDate    *time.Time
	ToDate      *time.Time
	Limit       *int
	Offset      *int
}

type ListSubscriptionsResult struct {
	Subs []*dto.Subscription
}

type ListSubscriptionsUseCase struct {
	subscriptions port.SubscriptionRepo
	logger        aport.Logger
}

func NewListSubscriptionsUseCase(
	subscriptions port.SubscriptionRepo,
	logger aport.Logger,
) *ListSubscriptionsUseCase {
	return &ListSubscriptionsUseCase{
		subscriptions: subscriptions,
		logger:        logger,
	}
}

func (uc *ListSubscriptionsUseCase) Execute(ctx context.Context, cmd ListSubscriptionsCommand) (*ListSubscriptionsResult, error) {
	uc.logger.Debug(ctx, "starting list subscriptions")

	subs, err := uc.subscriptions.List(
		ctx,
		cmd.UserID,
		cmd.ServiceName,
		cmd.FromDate,
		cmd.ToDate,
		cmd.Limit,
		cmd.Offset,
	)
	if err != nil {
		uc.logger.Error(ctx, "failed to get list of subscriptions",
			aport.Field{Key: "error", Value: err.Error()},
		)

		return nil, application.ErrSubscriptionListFailed
	}

	subResults := make([]*dto.Subscription, len(subs))
	for i, sub := range subs {
		var endDate *string
		if sub.EndDate() != nil {
			formattedDate := sub.EndDate().Format("01-2006")
			endDate = &formattedDate
		}

		subResults[i] = &dto.Subscription{
			ID:          sub.ID().String(),
			UserID:      sub.UserID().String(),
			ServiceName: sub.ServiceName(),
			Price:       sub.Price(),
			StartDate:   sub.StartDate().Format("01-2006"),
			EndDate:     endDate,
		}
	}

	uc.logger.Info(ctx, "list subscriptions successfully",
		aport.Field{Key: "subscriptions", Value: subResults},
	)

	return &ListSubscriptionsResult{
		Subs: subResults,
	}, err
}

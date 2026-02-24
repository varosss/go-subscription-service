package usecase

import (
	"context"
	"go-subscription-service/internal/application/dto"
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
}

func NewListSubscriptionsUseCase(
	subscriptions port.SubscriptionRepo,
) *ListSubscriptionsUseCase {
	return &ListSubscriptionsUseCase{
		subscriptions: subscriptions,
	}
}

func (uc *ListSubscriptionsUseCase) Execute(ctx context.Context, cmd ListSubscriptionsCommand) (*ListSubscriptionsResult, error) {
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
		return nil, err
	}

	subsResult := make([]*dto.Subscription, len(subs))
	for i, sub := range subs {
		var endDate *string
		if sub.EndDate() != nil {
			formattedDate := sub.EndDate().Format("01-2006")
			endDate = &formattedDate
		}

		subsResult[i] = &dto.Subscription{
			ID:          sub.ID().String(),
			UserID:      sub.UserID().String(),
			ServiceName: sub.ServiceName(),
			Price:       sub.Price(),
			StartDate:   sub.StartDate().Format("01-2006"),
			EndDate:     endDate,
		}
	}

	return &ListSubscriptionsResult{
		Subs: subsResult,
	}, err
}

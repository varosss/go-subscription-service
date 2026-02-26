package mapper

import (
	"go-subscription-service/internal/domain/entity"
	"go-subscription-service/internal/domain/valueobject"
	"go-subscription-service/internal/infrastructure/adapter/gorm/model"
)

func ToSubscriptionModel(d *entity.Subscription) *model.Subscription {
	return &model.Subscription{
		ID:          d.ID().String(),
		UserID:      d.UserID().String(),
		ServiceName: d.ServiceName(),
		Price:       d.Price(),
		StartDate:   d.StartDate(),
		EndDate:     d.EndDate(),
	}
}

func ToSubscriptionDomain(m model.Subscription) *entity.Subscription {
	return entity.SubscriptionFromPrimitives(
		valueobject.SubscriptionID(m.ID),
		valueobject.UserID(m.UserID),
		m.ServiceName,
		m.Price,
		m.StartDate,
		m.EndDate,
	)
}

package repository

import (
	"context"
	"time"

	"go-subscription-service/internal/domain/entity"
	"go-subscription-service/internal/domain/valueobject"
	"go-subscription-service/internal/infrastructure/adapter/gorm/mapper"
	"go-subscription-service/internal/infrastructure/adapter/gorm/model"

	"gorm.io/gorm"
)

type gormSubscriptionRepo struct {
	db *gorm.DB
}

func NewGormSubscriptionRepo(db *gorm.DB) *gormSubscriptionRepo {
	return &gormSubscriptionRepo{db: db}
}

func (r *gormSubscriptionRepo) Save(ctx context.Context, subscription *entity.Subscription) error {
	model := mapper.ToSubscribtionModel(subscription)

	now := time.Now()

	model.CreatedAt = now
	model.UpdatedAt = now

	return r.db.WithContext(ctx).Save(model).Error
}

func (r *gormSubscriptionRepo) GetByID(ctx context.Context, subscriptionID valueobject.SubscriptionID) (*entity.Subscription, error) {
	var model model.Subscription
	err := r.db.WithContext(ctx).First(&model, "id = ?", subscriptionID.String()).Error
	if err != nil {
		return nil, err
	}

	return mapper.ToSubscribtionDomain(model), nil
}

func (r *gormSubscriptionRepo) DeleteByID(ctx context.Context, subscriptionID valueobject.SubscriptionID) error {
	var model model.Subscription
	return r.db.WithContext(ctx).Delete(&model, "id = ?", subscriptionID.String()).Error
}

func (r *gormSubscriptionRepo) List(
	ctx context.Context,
	userID *valueobject.UserID,
	serviceName *string,
	fromDate *time.Time,
	toDate *time.Time,
	limit *int,
	offset *int,
) ([]*entity.Subscription, error) {
	var subModels []model.Subscription
	query := r.db.WithContext(ctx).Model(&model.Subscription{})

	if userID != nil {
		query = query.Where("user_id = ?", userID.String())
	}

	if serviceName != nil {
		query = query.Where("service_name ILIKE ?", "%"+*serviceName+"%")
	}

	if fromDate != nil && toDate != nil {
		query = query.Where("start_date <= ? AND (end_date IS NULL OR end_date >= ?)", *toDate, *fromDate)
	} else if fromDate != nil {
		query = query.Where("(end_date IS NULL OR end_date >= ?)", *fromDate)
	} else if toDate != nil {
		query = query.Where("start_date <= ?", *toDate)
	}

	if limit != nil {
		query = query.Limit(*limit)
	}

	if offset != nil {
		query = query.Offset(*offset)
	}

	err := query.Order("start_date desc").Find(&subModels).Error
	if err != nil {
		return nil, err
	}

	subs := make([]*entity.Subscription, len(subModels))
	for i, model := range subModels {
		subs[i] = mapper.ToSubscribtionDomain(model)
	}

	return subs, nil
}

func (r *gormSubscriptionRepo) CalculateTotalCost(
	ctx context.Context,
	userID *valueobject.UserID,
	serviceName *string,
	fromDate time.Time,
	toDate time.Time,
) (int64, error) {

	query := r.db.WithContext(ctx).
		Model(&model.Subscription{}).
		Where("start_date <= ?", toDate).
		Where("(end_date IS NULL OR end_date >= ?)", fromDate)

	if userID != nil {
		query = query.Where("user_id = ?", userID.String())
	}

	if serviceName != nil {
		query = query.Where("service_name ILIKE ?", "%"+*serviceName+"%")
	}

	var total int64

	err := query.Select(`
		COALESCE(SUM(
			(
				(
					EXTRACT(YEAR FROM age(
						LEAST(COALESCE(end_date, ?), ?),
						GREATEST(start_date, ?)
					)) * 12
				)
				+
				EXTRACT(MONTH FROM age(
					LEAST(COALESCE(end_date, ?), ?),
					GREATEST(start_date, ?)
				))
				+ 1
			) * price
		), 0)
	`,
		toDate, toDate, fromDate,
		toDate, toDate, fromDate,
	).Scan(&total).Error

	if err != nil {
		return 0, err
	}

	return total, nil
}

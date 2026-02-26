package application

import "errors"

var (
	ErrSubscriptionCreationFailed = errors.New("failed to create new subscription")
	ErrSubscriptionNotFound       = errors.New("subscription not found")
	ErrSubscriptionListFailed     = errors.New("subscription list failed")
	ErrInvalidSubscriptionData    = errors.New("invalid subscription data")
	ErrSubscriptionSaveFailed     = errors.New("could not save subscription")
	ErrSubscriptionDeleteFailed   = errors.New("could not delete subscription")
	ErrTotalCostCalculationFailed = errors.New("failed to calculate total cost")
)

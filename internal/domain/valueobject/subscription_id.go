package valueobject

import "github.com/google/uuid"

type SubscriptionID string

func NewSubscriptionID() SubscriptionID {
	return SubscriptionID(uuid.NewString())
}

func ParseSubscriptionID(id string) (SubscriptionID, error) {
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		return "", err
	}

	return SubscriptionID(parsedUUID.String()), nil
}

func (id SubscriptionID) String() string {
	return string(id)
}

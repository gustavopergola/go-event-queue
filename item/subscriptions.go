package item

import "go-orquestrator/observer"

type PriceChangedEvent struct {
	Item     Item
	NewPrice int
}

type PriceChangedSubscription struct {
	observer.Subscription[PriceChangedEvent]
}

package cart

import (
	"go-orquestrator/item"
	"go-orquestrator/observer"
)

type ItemAddedEvent struct {
	Cart Cart
	Item item.Item
	Qtd  int
}

type ItemAddedSubscription struct {
	observer.Subscription[ItemAddedEvent]
}

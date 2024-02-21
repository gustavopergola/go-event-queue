package item

import (
	"fmt"
	"go-orquestrator/cart"
)

type Service struct {
	PriceChangedSubscription PriceChangedSubscription
}

func (s Service) ChangePrice(i *Item, newPrice int) {
	i.Price = newPrice

	s.PriceChangedSubscription.NotifyAll(PriceChangedEvent{*i, newPrice})
}

func (s Service) OnNotify(event cart.ItemAddedEvent) {
	fmt.Println("update total items left in stock when item was added to some cart")
}

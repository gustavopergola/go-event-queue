package cart

import (
	"fmt"
	"go-orquestrator/item"
)

type Service struct {
	ItemAddedSubscription ItemAddedSubscription
}

func (s Service) AddItem(cart *Cart, item item.Item, qtd int) {
	s.ItemAddedSubscription.NotifyAll(ItemAddedEvent{
		*cart,
		item,
		0,
	})
}

// TODO how to bypass OnNotify same method name problem
func (s Service) OnNotify(event any) {
	fmt.Printf("update cart prices when item prices changed")
}

func (s Service) OnNotify(event ItemAddedEvent) {
	fmt.Printf("update cart prices when item prices changed")
}

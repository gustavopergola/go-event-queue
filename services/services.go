package services

import (
	"go-orquestrator/cart"
	"go-orquestrator/item"
)

type ServiceContainer struct {
	CartService cart.Service
	ItemService item.Service
}

func NewServiceContainer() ServiceContainer {
	return ServiceContainer{
		CartService: cart.Service{},
		ItemService: item.Service{},
	}
}

func (sc *ServiceContainer) SubscribeAll() {
	sc.ItemService.PriceChangedSubscription.Register(sc.CartService)
	sc.CartService.ItemAddedSubscription.Register(sc.ItemService)
}

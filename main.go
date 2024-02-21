package main

import (
	"go-orquestrator/cart"
	"go-orquestrator/item"
	"go-orquestrator/services"
)

func main() {
	container := services.NewServiceContainer()
	container.SubscribeAll()

	item1 := item.Item{Name: "Item1", Price: 10}
	item2 := item.Item{Name: "Item2", Price: 20}

	cart1 := cart.Cart{Name: "Cart1"}
	cart2 := cart.Cart{Name: "Cart2"}

	container.ItemService.ChangePrice(&item1, 15)
}

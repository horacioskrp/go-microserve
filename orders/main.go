package main

import "context"

func main() {

	store := NewStore()
	srvice := NewService(store)
	srvice.CreateOrder(context.Background())
}

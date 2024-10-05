package main

import "context"

type store struct {
	//ad mongo instance
}

func NewStore() *store {
	return &store{}
}

func (s *store) Create(context.Context) error {
	return nil
}

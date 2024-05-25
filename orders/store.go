package main

import "context"

type store struct {
	// Add mongodb instance here
}

func NewStore() *store {
	return &store{}
}

func (s *store) Create(context.Context) error {
	return nil
}

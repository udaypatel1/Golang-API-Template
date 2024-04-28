package repository

import (
	"sync"

	"github.com/udaypatel3/model"
)

type ItemRepository interface {
	GetAll() []model.Item
	Add(item model.Item) error
}

type InMemoryItemRepository struct {
	store map[int]model.Item
	mu    sync.Mutex
}

func NewInMemoryItemRepository() *InMemoryItemRepository {
	return &InMemoryItemRepository{
		store: make(map[int]model.Item),
	}
}

func (r *InMemoryItemRepository) GetAll() []model.Item {
	r.mu.Lock()
	defer r.mu.Unlock()

	items := make([]model.Item, 0, len(r.store))
	for _, item := range r.store {
		items = append(items, item)
	}

	return items
}

func (r *InMemoryItemRepository) Add(item model.Item) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.store[item.ID] = item
	return nil
}

package service

import (
	"github.com/udaypatel3/model"
	"github.com/udaypatel3/repository"
)

type ItemService interface {
	ListItems() []model.Item
	AddItem(item model.Item) error
}

type DefaultItemService struct {
	repo repository.ItemRepository
}

// AddItem implements ItemService.
func (s *DefaultItemService) AddItem(item model.Item) error {
	return s.repo.Add(item)
}

func NewDefaultItemService(repo repository.ItemRepository) *DefaultItemService {
	return &DefaultItemService{
		repo: repo,
	}
}

func (s *DefaultItemService) ListItems() []model.Item {
	return s.repo.GetAll()
}

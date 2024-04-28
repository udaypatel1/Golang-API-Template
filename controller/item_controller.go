package controller

import (
	"encoding/json"
	"net/http"

	"github.com/udaypatel3/model"
	"github.com/udaypatel3/service"
)

type ItemController struct {
	service service.ItemService
}

func NewItemController(service service.ItemService) *ItemController {
	return &ItemController{
		service: service,
	}
}

func (c *ItemController) ListItemsHandler(w http.ResponseWriter, r *http.Request) {

	items := c.service.ListItems()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode((items))
}

func (c *ItemController) AddItemHandler(w http.ResponseWriter, r *http.Request) {

	var item model.Item

	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := c.service.AddItem(item); err != nil {
		http.Error(w, "Error adding item", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)

}

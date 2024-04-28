package main

import (
	"log"
	"net/http"

	"github.com/udaypatel3/controller"
	"github.com/udaypatel3/repository"
	"github.com/udaypatel3/service"
)

func main() {

	repo := repository.NewInMemoryItemRepository()
	itemService := service.NewDefaultItemService(repo)
	itemController := controller.NewItemController(itemService)

	http.HandleFunc("/items", itemController.ListItemsHandler)
	http.HandleFunc("/items/add", itemController.AddItemHandler)

	log.Println("Starting server on :8080...")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}

}

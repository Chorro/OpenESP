package main

import (
	"fmt"
	"net/http"

	"github.com/chorro/openesp/api"
	"github.com/chorro/openesp/repository"
	"github.com/chorro/openesp/service"
)

func main() {
	memoryRepository := repository.NewMemoryDeviceRepository()

	deviceService := service.NewDeviceRule(memoryRepository)

	api := api.NewRestAPI(deviceService)
	fmt.Println("up and running: 0.0.0.0:8000")
	http.ListenAndServe("0.0.0.0:8000", api.Router())
}

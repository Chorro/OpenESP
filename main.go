package main

import (
	"fmt"

	"github.com/chorro/openesp/repository"
	"github.com/chorro/openesp/service"
)

func main() {
	memoryRepository := repository.NewMemoryDeviceRepository()

	deviceService := service.NewDeviceRule(memoryRepository)

	deviceService.CreateDevice("my-device")
	if err := deviceService.CreateDevice("too-long-device-name."); err != nil {
		fmt.Println("error: ", err)
	}

	devices, _ := deviceService.GetDevices()
	fmt.Println(devices)
}

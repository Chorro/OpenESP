package service

import (
	"fmt"

	"github.com/chorro/openesp/model"
	"github.com/chorro/openesp/repository"
)

type DeviceRule interface {
	CreateDevice(name string) error
	GetDevices() (model.DevicesList, error)
	GetDeviceByID(ID string) (*model.Device, error)
	UpdateDevice(device model.Device) error
	DeleteDeviceByID(ID string) error
}

type memoryRuleAdapter struct {
	Repository repository.DeviceRepository
}

func NewDeviceRule(repository repository.DeviceRepository) DeviceRule {
	return &memoryRuleAdapter{
		Repository: repository,
	}
}

func (m *memoryRuleAdapter) CreateDevice(name string) error {
	if name == "" || len(name) > 20 {
		return fmt.Errorf("incorrect device's name")
	}

	return m.Repository.CreateDevice(name)
}

func (m *memoryRuleAdapter) GetDevices() (model.DevicesList, error) {
	return m.Repository.GetDevices()
}

func (m *memoryRuleAdapter) GetDeviceByID(ID string) (*model.Device, error) {
	return m.Repository.GetDeviceByID(ID)
}

func (m *memoryRuleAdapter) UpdateDevice(device model.Device) error {
	return m.Repository.UpdateDevice(device)
}

func (m *memoryRuleAdapter) DeleteDeviceByID(ID string) error {
	return m.Repository.DeleteDeviceByID(ID)
}

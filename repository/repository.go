package repository

import "github.com/chorro/openesp/model"

// CRUD
// Create
// Read
// Update
// Delete

type DeviceRepository interface {
	CreateDevice(name string) error
	GetDevices() (model.DevicesList, error)
	GetDeviceByID(ID string) (*model.Device, error)
	UpdateDevice(device model.Device) error
	DeleteDeviceByID(ID string) error
}

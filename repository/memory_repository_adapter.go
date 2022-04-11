package repository

import "github.com/chorro/openesp/model"

type memoryRepositoryAdapter struct {
	devices map[string]model.Device
}

func NewMemoryDeviceRepository() DeviceRepository {
	return &memoryRepositoryAdapter{
		make(map[string]model.Device),
	}
}

func (m *memoryRepositoryAdapter) CreateDevice(name string) error {
	m.devices["1"] = model.Device{
		ID:   "1",
		Name: name,
	}

	return nil
}

func (m *memoryRepositoryAdapter) GetDevices() (model.DevicesList, error) {
	var list model.DevicesList
	for _, v := range m.devices {
		list = append(list, v)
	}

	return list, nil
}

func (m *memoryRepositoryAdapter) GetDeviceByID(ID string) (*model.Device, error) {
	if dev, ok := m.devices[ID]; ok {
		return &dev, nil
	}

	return nil, nil
}

func (m *memoryRepositoryAdapter) UpdateDevice(device model.Device) error {
	m.devices[device.ID] = device
	return nil
}

func (m *memoryRepositoryAdapter) DeleteDeviceByID(ID string) error {
	delete(m.devices, ID)
	return nil
}

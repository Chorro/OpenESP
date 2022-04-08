package model

type DevicesList []Device

type Device struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

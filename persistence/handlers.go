package persistence

import (
	"errors"
	"github.com/kubeIoT/api-docs/models"
)

func GetApplications() []*models.Application {
	res := make([]*models.Application, 0, len(applications))
	for _, k := range applications {
		res = append(res, k.Transform())
	}
	return res
}

func GetImages() []*models.Image {
	res := make([]*models.Image, 0, len(images))
	for _, k := range images {
		res = append(res, k.Transorm())
	}
	return res
}

func GetImageById(id int64) (*models.Image, error) {
	if v, ok := images[id]; ok {
		return v.Transorm(), nil
	}
	return nil, errors.New("Image not found")
}

func GetApplicationById(id int64) (*models.Application, error) {
	if v, ok := applications[id]; ok {
		return v.Transform(), nil
	}
	return nil, errors.New("Not found")
}

func GetCapabilities() []*models.Capability {
	res := make([]*models.Capability, 0, len(images))
	for _, k := range capabilities {
		res = append(res, k.Transform())
	}
	return res
}

func GetCapabilityById(id int64) (*models.Capability, error) {
	if v, ok := capabilities[id]; ok {
		return v.Transform(), nil
	}
	return nil, errors.New("Not found")
}

func GetUsers() []*models.User {
	res := make([]*models.User, 0, len(images))
	for _, k := range users {
		res = append(res, k.Transform())
	}
	return res
}

func GetUsersById(id int64) (*models.User, error) {
	if v, ok := users[id]; ok {
		return v.Transform(), nil
	}
	return nil, errors.New("Not found")
}

func GetDevices() []*models.Device {
	res := make([]*models.Device, 0, len(images))
	for _, k := range devices {
		res = append(res, k.Transform())
	}
	return res
}

func GetDeviceById(id int64) (*models.Device, error) {
	if v, ok := devices[id]; ok {
		return v.Transform(), nil
	}
	return nil, errors.New("Not found")
}

func GetDeviceCapabilities(id int64) (res []*models.DeviceCapability, err error) {
	res = nil
	if v, ok := devices[id]; ok {
		res = make([]*models.DeviceCapability, 0, len(v.InstalledCapabilities))

		for _, k := range v.InstalledCapabilities {
			res = append(res, k.Transform())
		}
		return
	}
	err = errors.New("Device not found")
	return
}

func GetDeviceCapabilitiesById(id, capId int64) (*models.DeviceCapability, error) {
	if v, ok := devices[id]; ok {
		if k, ok := v.InstalledCapabilities[capId]; ok {
			return k.Transform(), nil
		}
		return nil, errors.New("Capability not found")
	}
	return nil, errors.New("Device not found")
}

func GetDeviceEvents(id int64) ([]*models.Event, error) {
	if v, ok := devices[id]; ok {
		res := make([]*models.Event, len(v.events))
		for _, k := range v.events {
			res = append(res, k.Transform())
		}
		return res, nil
	}
	return nil, errors.New("Device not found")
}

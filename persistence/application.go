package persistence

import (
	"sync/atomic"

	"github.com/kubeIoT/api-docs/models"
)

//Application database entry
type Application struct {
	// base image
	BaseImage *Image

	// Used capabilities on the device
	Capabilities map[int64]*DeviceCapability

	//The Owning device
	OwningDevice *Device

	// id
	// Read Only: true
	ID int64

	// name
	Name string

	// ports
	Ports []string

	// Ip of kubernetes service
	ServiceIP string

	// status
	Status int64

	// status message
	StatusMessage string
}

func (app Application) Transform() (v *models.Application) {
	v = &models.Application{Ports: make([]string, len(app.Ports))}
	v.BaseImage = app.BaseImage.ID
	v.Capabilities = getIds(app.Capabilities)
	v.DeviceID = app.OwningDevice.ID
	v.ID = app.ID
	v.Name = app.Name
	copy(v.Ports, app.Ports)
	v.ServiceIP = app.ServiceIP
	v.Status = app.Status
	v.StatusMessage = app.StatusMessage

	return
}

func NewApplication(name string, image *Image, ports []string, sIp string) (app *Application) {
	atomic.AddInt64(&lastApplicationID, 1)
	app = &Application{
		ID:            lastApplicationID,
		Name:          name,
		BaseImage:     image,
		ServiceIP:     sIp,
		Status:        0,
		StatusMessage: "Ok",
		Capabilities:  make(map[int64]*DeviceCapability),
		Ports:         make([]string, len(ports))}
	copy(app.Ports, ports)
	return
}

package persistence

import (
	"sync/atomic"

	"github.com/kubeIoT/api-docs/models"
)

//Capability database entry
type Capability struct {
	// id
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// Name of the peripheral device
	PeripherialDevice string `json:"peripherial_device,omitempty"`

	// Communication protocol of capability
	Protocol string `json:"protocol,omitempty"`
}

func (cap Capability) Transform() (c *models.Capability) {
	c = &models.Capability{
		ID:                cap.ID,
		Name:              cap.Name,
		PeripherialDevice: cap.PeripherialDevice,
		Protocol:          cap.Protocol}
	return
}

func NewCapability(name, peripheralDevice, protocol string) (cap *Capability) {
	atomic.AddInt64(&lastCapabilityID, 1)

	cap = &Capability{
		ID:                lastCapabilityID,
		Name:              name,
		PeripherialDevice: peripheralDevice,
		Protocol:          protocol}

	return
}

package persistence

import (
	"errors"
	"sync/atomic"

	"github.com/kubeIoT/api-docs/models"
)

//Device database entry
type Device struct {
	// address
	Address string

	// applications
	Applications map[int64]*Application

	// device vendor
	DeviceVendor string

	// device version
	DeviceVersion string

	// id
	// Read Only: true
	ID int64

	// installed capabilities
	InstalledCapabilities map[int64]*DeviceCapability

	// kernel version
	KernelVersion string

	// number of applications
	NumberOfApplications int64

	// os distribution
	OsDistribution string

	// System id of user owning device
	Owner *User

	// system info
	SystemInfo string

	// used capabilities
	UsedCapabilities map[int64]int64

	events map[int64]*Event

	lastCapID int64

	lastEventID int64
}

//DeviceCapability database entry
type DeviceCapability struct {
	// app id
	App *Application

	// The bus on device where the capability is connected
	// Format cannot be currently specified, we don't know yet
	// how the connection to peripheral will look like
	//
	BusConnection string

	// Id to capability type
	OccupiedCapability *Capability

	// id
	// Read Only: true
	ID int64

	// Is the capability used by application
	Used bool
}

type Event struct {
	// event message
	EventMessage string

	// event timestamp
	EventTimestamp string

	// event type
	EventType string

	// id
	// Read Only: true
	ID int64

	// Id of owner of event
	//
	parent *Device
}

func (dev Device) Transform() (d *models.Device) {
	d = &models.Device{
		Address:       dev.Address,
		Applications:  getIds(dev.Applications),
		DeviceVendor:  dev.DeviceVendor,
		DeviceVersion: dev.DeviceVersion,
		ID:            dev.ID,
		InstalledCapabilities: getIds(dev.InstalledCapabilities),
		KernelVersion:         dev.KernelVersion,
		NumberOfApplications:  dev.NumberOfApplications,
		OsDistribution:        dev.OsDistribution,
		Owner:                 dev.Owner.ID,
		SystemInfo:            dev.SystemInfo,
		UsedCapabilities:      make([]*models.DeviceUsedCapabilitiesItems0, 0, len(dev.UsedCapabilities)),
	}
	for k, v := range dev.UsedCapabilities {
		d.UsedCapabilities = append(d.UsedCapabilities, &models.DeviceUsedCapabilitiesItems0{ApplicationID: k, CapabilityID: v})
	}

	return
}

func (ev Event) Transform() (event *models.Event) {
	return &models.Event{
		EventMessage:   ev.EventMessage,
		EventTimestamp: ev.EventTimestamp,
		EventType:      ev.EventType,
		ID:             ev.ID,
		ParentID:       ev.parent.ID}

}

func (cap DeviceCapability) Transform() (c *models.DeviceCapability) {
	c = &models.DeviceCapability{
		AppID:         cap.ID,
		BusConnection: cap.BusConnection,
		CapID:         cap.OccupiedCapability.ID,
		ID:            cap.ID,
		Used:          cap.Used}
	return
}

func NewDevice(address, deviceVendor, deviceVersion, kernelVersion, OsDistribution string) (dev *Device) {
	atomic.AddInt64(&lastDeviceID, 1)
	dev = &Device{
		ID:                    lastDeviceID,
		DeviceVendor:          deviceVendor,
		KernelVersion:         kernelVersion,
		OsDistribution:        OsDistribution,
		SystemInfo:            "Running",
		NumberOfApplications:  0,
		InstalledCapabilities: make(map[int64]*DeviceCapability),
		Applications:          make(map[int64]*Application),
		UsedCapabilities:      make(map[int64]int64),
		events:                make(map[int64]*Event)}
	return
}

func (dev *Device) addCapability(busConnection string, capability *Capability) (cap *DeviceCapability) {
	atomic.AddInt64(&dev.lastCapID, 1)
	cap = &DeviceCapability{
		ID:                 dev.lastCapID,
		BusConnection:      busConnection,
		OccupiedCapability: capability,
		Used:               false}
	dev.InstalledCapabilities[capability.ID] = cap
	return
}

func (dev *Device) addApplication(app *Application) (err error) {
	err = nil
	capsToChange := make([]*DeviceCapability, 0, 10)
	for id := range app.BaseImage.RequiredCapabilities {
		if v, ok := dev.InstalledCapabilities[id]; ok {
			capsToChange = append(capsToChange, v)
		} else {
			err = errors.New("Capability not found")
		}
	}
	for _, v := range capsToChange {
		app.Capabilities[v.ID] = v
		v.App = app
		v.Used = true
		dev.UsedCapabilities[v.ID] = app.ID
	}
	app.OwningDevice = dev
	dev.Applications[app.ID] = app
	return
}

func (dev *Device) addEvent(message, timestamp, eventType string) {
	atomic.AddInt64(&dev.lastEventID, 1)
	r := &Event{
		EventMessage:   message,
		EventTimestamp: timestamp,
		EventType:      eventType,
		ID:             dev.lastEventID,
		parent:         dev}
	dev.events[r.ID] = r
}

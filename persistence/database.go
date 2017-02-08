package persistence

import (
	"fmt"
	"reflect"
	"sync"
)

func getIds(v interface{}) []int64 {
	vVal := reflect.ValueOf(v)
	if vVal.Kind() != reflect.Map {
		return nil
	}
	vType := vVal.Type()
	if vType.Key().Kind() != reflect.Int64 {
		return nil
	}
	result := make([]int64, 0, len(vVal.MapKeys()))
	for _, key := range vVal.MapKeys() {
		result = append(result, key.Int())
	}
	return result
}

type Transformable interface {
	Transform() interface{}
}

var accMutt sync.Mutex
var images = make(map[int64]*Image)
var lastImageID int64

var applications = make(map[int64]*Application)
var lastApplicationID int64

var capabilities = make(map[int64]*Capability)
var lastCapabilityID int64

var devices = make(map[int64]*Device)
var lastDeviceID int64

var users = make(map[int64]*User)
var lastUserID int64

func Bootstrap() {
	l_users := []*User{
		NewUser("Test Tst", "test@example.com", "2.2.2002", "3.2.2017", true),
		NewUser("Test User2", "test2@example.com", "12.12.2012", "12.12.2012", true)}

	l_capabilities := []*Capability{
		NewCapability("Servo Motor 1", "servo_motor", "PWM"),
		NewCapability("Led diode blue", "led_diode", "GPIO"),
		NewCapability("Led diode Green", "led_diode", "GPIO"),
		NewCapability("Thermometer scientific", "thermometer", "I2C")}

	l_images := []*Image{
		NewImage("Thermometer publisher", "Pulishes data from thermometer via http", "Ubuntu 10.04", []string{"8080", "8443"},
			[]*Capability{l_capabilities[3]}),
		NewImage("Led blinker", "Blinkes blue and green diode", "Debian wheezy", []string{"8080"},
			[]*Capability{l_capabilities[1], l_capabilities[2]}),
		NewImage("Door opener", "Opens door while blinking blue diode", "Debian wheezy", []string{"1234"},
			[]*Capability{l_capabilities[1], l_capabilities[3]})}

	l_devices := []*Device{
		NewDevice("192.168.0.12", "BeagleBone", "BeagleBoneBlack", "4.8.2", "Debian jessie"),
		NewDevice("192.168.0.13", "RaspberryPi", "Zero", "4.9.2", "Ubuntu core"),
		NewDevice("192.168.0.14", "RaspberryPi", "Rpi 3", "4.7.2", "Fedora 25")}

	l_devices[0].addCapability("GPIO/2", l_capabilities[1])
	l_devices[0].addCapability("GPIO/3", l_capabilities[2])
	l_devices[0].addCapability("PWM/1", l_capabilities[0])
	l_devices[1].addCapability("GPIO/1", l_capabilities[1])
	l_devices[1].addCapability("GPIO/2", l_capabilities[1])
	l_devices[1].addCapability("GPIO/3", l_capabilities[2])
	l_devices[1].addCapability("I2C/1", l_capabilities[3])
	l_devices[2].addCapability("GPIO/1", l_capabilities[2])

	fmt.Println("The users", l_users, "The devices", l_devices)
	l_users[0].addDevice(l_devices[0])
	l_users[1].addDevice(l_devices[1])
	l_users[1].addDevice(l_devices[2])

	l_application := []*Application{
		NewApplication("My app 1", l_images[1], []string{"8080:8080"}, "172.17.0.10"),
		NewApplication("My app 2", l_images[2], []string{}, "192.168.0.3"),
		NewApplication("My app 3", l_images[2], []string{"9090:9090", "22:22", "443:443"}, "10.0.0.5"),
		NewApplication("My app 4", l_images[0], []string{"21:21"}, "10.0.0.6")}

	l_devices[1].addApplication(l_application[0])
	l_devices[1].addApplication(l_application[1])
	l_devices[2].addApplication(l_application[2])
	l_devices[0].addApplication(l_application[3])

	l_devices[0].addEvent("Device started", "00102023", "System")
	l_devices[0].addEvent("Shutdown scheduled", "00102028", "System")
	l_devices[1].addEvent("SysV init segmentation fault", "---", "System")
	l_devices[2].addEvent("Device on fire", "29218310", "System")

	for _, k := range l_users {
		users[k.ID] = k
	}

	for _, k := range l_capabilities {
		capabilities[k.ID] = k
	}

	for _, k := range l_images {
		images[k.ID] = k
	}

	for _, k := range l_devices {
		devices[k.ID] = k
	}

	for _, k := range l_application {
		applications[k.ID] = k
	}
}

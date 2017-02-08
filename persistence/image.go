package persistence

import (
	"sync/atomic"

	"fmt"
	"github.com/kubeIoT/api-docs/models"
)

// Image database entry
type Image struct {
	// base image
	BaseImage string

	Name string

	Description string

	// Ports exposed by application
	ExposedPorts []string

	// id
	// Read Only: true
	ID int64

	// required capabilities
	RequiredCapabilities map[int64]*Capability
}

func (im Image) Transorm() (res *models.Image) {
	fmt.Println(im.BaseImage, im.ExposedPorts, im.Description)
	res = &models.Image{
		BaseImage:            im.BaseImage,
		Description:          im.Description,
		ExposedPorts:         make([]string, len(im.ExposedPorts)),
		ID:                   im.ID,
		Name:                 im.Name,
		RequiredCapabilities: getIds(im.RequiredCapabilities)}
	copy(res.ExposedPorts, im.ExposedPorts)
	fmt.Println("Transforing image", res.ExposedPorts, res.RequiredCapabilities)
	return
}

func NewImage(name, description, baseImage string, exposedPorts []string, cap []*Capability) (image *Image) {
	atomic.AddInt64(&lastImageID, 1)
	image = &Image{
		BaseImage:            baseImage,
		ID:                   lastImageID,
		Description:          description,
		Name:                 name,
		RequiredCapabilities: make(map[int64]*Capability),
		ExposedPorts:         make([]string, len(exposedPorts))}
	for _, k := range cap {
		fmt.Println("Saving this", k)
		image.RequiredCapabilities[k.ID] = k
	}
	copy(image.ExposedPorts, exposedPorts)
	return
}

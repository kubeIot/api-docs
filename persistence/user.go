package persistence

import (
	"errors"
	"sync/atomic"

	"github.com/kubeIoT/api-docs/models"
)

//User database entry
type User struct {
	// account created
	AccountCreated string `json:"account_created,omitempty"`

	// devices
	Devices map[int64]*Device

	// email
	Email string `json:"email,omitempty"`

	// email verified
	EmailVerified bool `json:"email_verified,omitempty"`

	// id
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// last logged in
	LastLoggedIn string `json:"last_logged_in,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

func (us User) Transform() (res *models.User) {

	res = &models.User{
		AccountCreated: us.AccountCreated,
		Devices:        getIds(us.Devices),
		Email:          us.Email,
		EmailVerified:  us.EmailVerified,
		ID:             us.ID,
		LastLoggedIn:   us.LastLoggedIn,
		Username:       us.Username}
	return
}

func NewUser(username, email, accountCreated, lastLoggedIn string, emailVerified bool) (user *User) {
	atomic.AddInt64(&lastUserID, 1)

	user = &User{
		ID:             lastUserID,
		Email:          email,
		AccountCreated: accountCreated,
		LastLoggedIn:   lastLoggedIn,
		EmailVerified:  emailVerified,
		Devices:        make(map[int64]*Device)}
	return
}

func (user *User) addDevice(device *Device) (err error) {
	if device.Owner != nil {
		err = errors.New("Device is already owned by user")
	}
	device.Owner = user
	user.Devices[device.ID] = device
	return
}

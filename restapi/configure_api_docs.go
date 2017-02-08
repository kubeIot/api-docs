package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/tylerb/graceful"

	"github.com/kubeIoT/api-docs/models"
	"github.com/kubeIoT/api-docs/persistence"
	"github.com/kubeIoT/api-docs/restapi/operations"
	"github.com/kubeIoT/api-docs/restapi/operations/application"
	"github.com/kubeIoT/api-docs/restapi/operations/capability"
	"github.com/kubeIoT/api-docs/restapi/operations/device"
	"github.com/kubeIoT/api-docs/restapi/operations/image"
	"github.com/kubeIoT/api-docs/restapi/operations/user"
)

// This file is safe to edit. Once it exists it will not be overwritten

//go:generate swagger generate server --target .. --name apiDocs --spec ../swagger.yml

func configureFlags(api *operations.APIDocsAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.APIDocsAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// s.api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.ApplicationDeleteApplicationIDHandler = application.DeleteApplicationIDHandlerFunc(func(params application.DeleteApplicationIDParams) middleware.Responder {
		return middleware.NotImplemented("operation application.DeleteApplicationID has not yet been implemented")
	})
	api.CapabilityDeleteCapabilityIDHandler = capability.DeleteCapabilityIDHandlerFunc(func(params capability.DeleteCapabilityIDParams) middleware.Responder {
		return middleware.NotImplemented("operation capability.DeleteCapabilityID has not yet been implemented")
	})
	api.DeviceDeleteDeviceIDHandler = device.DeleteDeviceIDHandlerFunc(func(params device.DeleteDeviceIDParams) middleware.Responder {
		return middleware.NotImplemented("operation device.DeleteDeviceID has not yet been implemented")
	})
	api.DeviceDeleteDeviceIDCapabilitiesHandler = device.DeleteDeviceIDCapabilitiesHandlerFunc(func(params device.DeleteDeviceIDCapabilitiesParams) middleware.Responder {
		return middleware.NotImplemented("operation device.DeleteDeviceIDCapabilities has not yet been implemented")
	})
	api.ImageDeleteImageIDHandler = image.DeleteImageIDHandlerFunc(func(params image.DeleteImageIDParams) middleware.Responder {
		return middleware.NotImplemented("operation image.DeleteImageID has not yet been implemented")
	})
	api.UserDeleteUserIDHandler = user.DeleteUserIDHandlerFunc(func(params user.DeleteUserIDParams) middleware.Responder {
		return middleware.NotImplemented("operation user.DeleteUserID has not yet been implemented")
	})
	api.ApplicationGetApplicationHandler = application.GetApplicationHandlerFunc(func(params application.GetApplicationParams) middleware.Responder {
		return application.NewGetApplicationOK().WithPayload(persistence.GetApplications())
	})
	api.ApplicationGetApplicationIDHandler = application.GetApplicationIDHandlerFunc(func(params application.GetApplicationIDParams) middleware.Responder {
		data, err := persistence.GetApplicationById(params.ID)
		if err != nil {
			return application.NewGetApplicationIDNotFound().WithPayload(&models.Error{Code: 404, Message: swag.String(err.Error())})
		}
		return application.NewGetApplicationIDOK().WithPayload(data)
	})
	api.CapabilityGetCapabilityHandler = capability.GetCapabilityHandlerFunc(func(params capability.GetCapabilityParams) middleware.Responder {
		return capability.NewGetCapabilityOK().WithPayload(persistence.GetCapabilities())
	})
	api.CapabilityGetCapabilityIDHandler = capability.GetCapabilityIDHandlerFunc(func(params capability.GetCapabilityIDParams) middleware.Responder {
		data, err := persistence.GetCapabilityById(params.ID)
		if err != nil {
			return capability.NewGetCapabilityIDNotFound().WithPayload(&models.Error{Code: 404, Message: swag.String(err.Error())})
		}
		return capability.NewGetCapabilityIDOK().WithPayload(data)
	})
	api.DeviceGetDeviceHandler = device.GetDeviceHandlerFunc(func(params device.GetDeviceParams) middleware.Responder {
		return device.NewGetDeviceOK().WithPayload(persistence.GetDevices())
	})
	api.DeviceGetDeviceIDHandler = device.GetDeviceIDHandlerFunc(func(params device.GetDeviceIDParams) middleware.Responder {
		data, err := persistence.GetDeviceById(params.ID)
		if err != nil {
			return device.NewGetDeviceIDNotFound().WithPayload(&models.Error{Code: 404, Message: swag.String(err.Error())})
		}
		return device.NewGetDeviceIDOK().WithPayload(data)
	})
	api.DeviceGetDeviceIDCapabilitiesHandler = device.GetDeviceIDCapabilitiesHandlerFunc(func(params device.GetDeviceIDCapabilitiesParams) middleware.Responder {
		data, err := persistence.GetDeviceCapabilities(params.ID)
		if err != nil {
			return device.NewGetDeviceIDCapabilitiesNotFound().WithPayload(&models.Error{Code: 404, Message: swag.String(err.Error())})
		}
		return device.NewGetDeviceIDCapabilitiesOK().WithPayload(data)
	})
	api.DeviceGetDeviceIDEventsHandler = device.GetDeviceIDEventsHandlerFunc(func(params device.GetDeviceIDEventsParams) middleware.Responder {
		data, err := persistence.GetDeviceEvents(params.ID)
		if err != nil {
			return device.NewGetDeviceIDEventsNotFound().WithPayload(&models.Error{Code: 404, Message: swag.String(err.Error())})
		}
		return device.NewGetDeviceIDEventsOK().WithPayload(data)
	})
	api.ImageGetImageHandler = image.GetImageHandlerFunc(func(params image.GetImageParams) middleware.Responder {
		return image.NewGetImageOK().WithPayload(persistence.GetImages())
	})
	api.ImageGetImageIDHandler = image.GetImageIDHandlerFunc(func(params image.GetImageIDParams) middleware.Responder {
		data, err := persistence.GetImageById(params.ID)
		if err != nil {
			return image.NewGetImageIDNotFound().WithPayload(&models.Error{Code: 404, Message: swag.String(err.Error())})
		}
		return image.NewGetImageIDOK().WithPayload(data)

	})
	api.UserGetUserHandler = user.GetUserHandlerFunc(func(params user.GetUserParams) middleware.Responder {
		return user.NewGetUserOK().WithPayload(persistence.GetUsers())
	})
	api.UserGetUserIDHandler = user.GetUserIDHandlerFunc(func(params user.GetUserIDParams) middleware.Responder {
		data, err := persistence.GetUsersById(params.ID)
		if err != nil {
			return user.NewGetUserIDNotFound().WithPayload(&models.Error{Code: 404, Message: swag.String(err.Error())})
		}
		return user.NewGetUserIDOK().WithPayload(data)
	})
	api.ApplicationPostApplicationHandler = application.PostApplicationHandlerFunc(func(params application.PostApplicationParams) middleware.Responder {
		return middleware.NotImplemented("operation application.PostApplication has not yet been implemented")
	})
	api.CapabilityPostCapabilityHandler = capability.PostCapabilityHandlerFunc(func(params capability.PostCapabilityParams) middleware.Responder {
		return middleware.NotImplemented("operation capability.PostCapability has not yet been implemented")
	})
	api.DevicePostDeviceHandler = device.PostDeviceHandlerFunc(func(params device.PostDeviceParams) middleware.Responder {
		return middleware.NotImplemented("operation device.PostDevice has not yet been implemented")
	})
	api.DevicePostDeviceIDCapabilitiesHandler = device.PostDeviceIDCapabilitiesHandlerFunc(func(params device.PostDeviceIDCapabilitiesParams) middleware.Responder {
		return middleware.NotImplemented("operation device.PostDeviceIDCapabilities has not yet been implemented")
	})
	api.ImagePostImageHandler = image.PostImageHandlerFunc(func(params image.PostImageParams) middleware.Responder {
		return middleware.NotImplemented("operation image.PostImage has not yet been implemented")
	})
	api.UserPostUserHandler = user.PostUserHandlerFunc(func(params user.PostUserParams) middleware.Responder {
		return middleware.NotImplemented("operation user.PostUser has not yet been implemented")
	})
	api.ApplicationPutApplicationIDHandler = application.PutApplicationIDHandlerFunc(func(params application.PutApplicationIDParams) middleware.Responder {
		return middleware.NotImplemented("operation application.PutApplicationID has not yet been implemented")
	})
	api.CapabilityPutCapabilityIDHandler = capability.PutCapabilityIDHandlerFunc(func(params capability.PutCapabilityIDParams) middleware.Responder {
		return middleware.NotImplemented("operation capability.PutCapabilityID has not yet been implemented")
	})
	api.DevicePutDeviceIDHandler = device.PutDeviceIDHandlerFunc(func(params device.PutDeviceIDParams) middleware.Responder {
		return middleware.NotImplemented("operation device.PutDeviceID has not yet been implemented")
	})
	api.DevicePutDeviceIDCapabilitiesHandler = device.PutDeviceIDCapabilitiesHandlerFunc(func(params device.PutDeviceIDCapabilitiesParams) middleware.Responder {
		return middleware.NotImplemented("operation device.PutDeviceIDCapabilities has not yet been implemented")
	})
	api.ImagePutImageIDHandler = image.PutImageIDHandlerFunc(func(params image.PutImageIDParams) middleware.Responder {
		return middleware.NotImplemented("operation image.PutImageID has not yet been implemented")
	})
	api.UserPutUserIDHandler = user.PutUserIDHandlerFunc(func(params user.PutUserIDParams) middleware.Responder {
		return middleware.NotImplemented("operation user.PutUserID has not yet been implemented")
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *graceful.Server, scheme string) {
	persistence.Bootstrap()
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}

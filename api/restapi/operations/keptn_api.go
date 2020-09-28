// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/keptn/keptn/api/models"
	"github.com/keptn/keptn/api/restapi/operations/auth"
	"github.com/keptn/keptn/api/restapi/operations/configuration"
	"github.com/keptn/keptn/api/restapi/operations/evaluation"
	"github.com/keptn/keptn/api/restapi/operations/event"
	"github.com/keptn/keptn/api/restapi/operations/metadata"
	"github.com/keptn/keptn/api/restapi/operations/project"
	"github.com/keptn/keptn/api/restapi/operations/service"
)

// NewKeptnAPI creates a new Keptn instance
func NewKeptnAPI(spec *loads.Document) *KeptnAPI {
	return &KeptnAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		ProjectDeleteProjectProjectNameHandler: project.DeleteProjectProjectNameHandlerFunc(func(params project.DeleteProjectProjectNameParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation project.DeleteProjectProjectName has not yet been implemented")
		}),
		ServiceDeleteProjectProjectNameServiceServiceNameHandler: service.DeleteProjectProjectNameServiceServiceNameHandlerFunc(func(params service.DeleteProjectProjectNameServiceServiceNameParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation service.DeleteProjectProjectNameServiceServiceName has not yet been implemented")
		}),
		ConfigurationGetConfigBridgeHandler: configuration.GetConfigBridgeHandlerFunc(func(params configuration.GetConfigBridgeParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation configuration.GetConfigBridge has not yet been implemented")
		}),
		EventGetEventHandler: event.GetEventHandlerFunc(func(params event.GetEventParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation event.GetEvent has not yet been implemented")
		}),
		ConfigurationPostConfigBridgeHandler: configuration.PostConfigBridgeHandlerFunc(func(params configuration.PostConfigBridgeParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation configuration.PostConfigBridge has not yet been implemented")
		}),
		EventPostEventHandler: event.PostEventHandlerFunc(func(params event.PostEventParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation event.PostEvent has not yet been implemented")
		}),
		ProjectPostProjectHandler: project.PostProjectHandlerFunc(func(params project.PostProjectParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation project.PostProject has not yet been implemented")
		}),
		ServicePostProjectProjectNameServiceHandler: service.PostProjectProjectNameServiceHandlerFunc(func(params service.PostProjectProjectNameServiceParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation service.PostProjectProjectNameService has not yet been implemented")
		}),
		AuthAuthHandler: auth.AuthHandlerFunc(func(params auth.AuthParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation auth.Auth has not yet been implemented")
		}),
		MetadataMetadataHandler: metadata.MetadataHandlerFunc(func(params metadata.MetadataParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation metadata.Metadata has not yet been implemented")
		}),
		EvaluationTriggerEvaluationHandler: evaluation.TriggerEvaluationHandlerFunc(func(params evaluation.TriggerEvaluationParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation evaluation.TriggerEvaluation has not yet been implemented")
		}),

		// Applies when the "x-token" header is set
		KeyAuth: func(token string) (*models.Principal, error) {
			return nil, errors.NotImplemented("api key auth (key) x-token from header param [x-token] has not yet been implemented")
		},
		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*KeptnAPI the keptn API */
type KeptnAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/cloudevents+json
	//   - application/json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// KeyAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key x-token provided in the header
	KeyAuth func(string) (*models.Principal, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// ProjectDeleteProjectProjectNameHandler sets the operation handler for the delete project project name operation
	ProjectDeleteProjectProjectNameHandler project.DeleteProjectProjectNameHandler
	// ServiceDeleteProjectProjectNameServiceServiceNameHandler sets the operation handler for the delete project project name service service name operation
	ServiceDeleteProjectProjectNameServiceServiceNameHandler service.DeleteProjectProjectNameServiceServiceNameHandler
	// ConfigurationGetConfigBridgeHandler sets the operation handler for the get config bridge operation
	ConfigurationGetConfigBridgeHandler configuration.GetConfigBridgeHandler
	// EventGetEventHandler sets the operation handler for the get event operation
	EventGetEventHandler event.GetEventHandler
	// ConfigurationPostConfigBridgeHandler sets the operation handler for the post config bridge operation
	ConfigurationPostConfigBridgeHandler configuration.PostConfigBridgeHandler
	// EventPostEventHandler sets the operation handler for the post event operation
	EventPostEventHandler event.PostEventHandler
	// ProjectPostProjectHandler sets the operation handler for the post project operation
	ProjectPostProjectHandler project.PostProjectHandler
	// ServicePostProjectProjectNameServiceHandler sets the operation handler for the post project project name service operation
	ServicePostProjectProjectNameServiceHandler service.PostProjectProjectNameServiceHandler
	// AuthAuthHandler sets the operation handler for the auth operation
	AuthAuthHandler auth.AuthHandler
	// MetadataMetadataHandler sets the operation handler for the metadata operation
	MetadataMetadataHandler metadata.MetadataHandler
	// EvaluationTriggerEvaluationHandler sets the operation handler for the trigger evaluation operation
	EvaluationTriggerEvaluationHandler evaluation.TriggerEvaluationHandler
	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *KeptnAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *KeptnAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *KeptnAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *KeptnAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *KeptnAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *KeptnAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *KeptnAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *KeptnAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *KeptnAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the KeptnAPI
func (o *KeptnAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.KeyAuth == nil {
		unregistered = append(unregistered, "XTokenAuth")
	}

	if o.ProjectDeleteProjectProjectNameHandler == nil {
		unregistered = append(unregistered, "project.DeleteProjectProjectNameHandler")
	}
	if o.ServiceDeleteProjectProjectNameServiceServiceNameHandler == nil {
		unregistered = append(unregistered, "service.DeleteProjectProjectNameServiceServiceNameHandler")
	}
	if o.ConfigurationGetConfigBridgeHandler == nil {
		unregistered = append(unregistered, "configuration.GetConfigBridgeHandler")
	}
	if o.EventGetEventHandler == nil {
		unregistered = append(unregistered, "event.GetEventHandler")
	}
	if o.ConfigurationPostConfigBridgeHandler == nil {
		unregistered = append(unregistered, "configuration.PostConfigBridgeHandler")
	}
	if o.EventPostEventHandler == nil {
		unregistered = append(unregistered, "event.PostEventHandler")
	}
	if o.ProjectPostProjectHandler == nil {
		unregistered = append(unregistered, "project.PostProjectHandler")
	}
	if o.ServicePostProjectProjectNameServiceHandler == nil {
		unregistered = append(unregistered, "service.PostProjectProjectNameServiceHandler")
	}
	if o.AuthAuthHandler == nil {
		unregistered = append(unregistered, "auth.AuthHandler")
	}
	if o.MetadataMetadataHandler == nil {
		unregistered = append(unregistered, "metadata.MetadataHandler")
	}
	if o.EvaluationTriggerEvaluationHandler == nil {
		unregistered = append(unregistered, "evaluation.TriggerEvaluationHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *KeptnAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *KeptnAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	result := make(map[string]runtime.Authenticator)
	for name := range schemes {
		switch name {
		case "key":
			scheme := schemes[name]
			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, func(token string) (interface{}, error) {
				return o.KeyAuth(token)
			})

		}
	}
	return result
}

// Authorizer returns the registered authorizer
func (o *KeptnAPI) Authorizer() runtime.Authorizer {
	return o.APIAuthorizer
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *KeptnAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/cloudevents+json":
			result["application/cloudevents+json"] = o.JSONConsumer
		case "application/json":
			result["application/json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *KeptnAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *KeptnAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the keptn API
func (o *KeptnAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *KeptnAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/project/{projectName}"] = project.NewDeleteProjectProjectName(o.context, o.ProjectDeleteProjectProjectNameHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/project/{projectName}/service/{serviceName}"] = service.NewDeleteProjectProjectNameServiceServiceName(o.context, o.ServiceDeleteProjectProjectNameServiceServiceNameHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/config/bridge"] = configuration.NewGetConfigBridge(o.context, o.ConfigurationGetConfigBridgeHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/event"] = event.NewGetEvent(o.context, o.EventGetEventHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/config/bridge"] = configuration.NewPostConfigBridge(o.context, o.ConfigurationPostConfigBridgeHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/event"] = event.NewPostEvent(o.context, o.EventPostEventHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/project"] = project.NewPostProject(o.context, o.ProjectPostProjectHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/project/{projectName}/service"] = service.NewPostProjectProjectNameService(o.context, o.ServicePostProjectProjectNameServiceHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/auth"] = auth.NewAuth(o.context, o.AuthAuthHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/metadata"] = metadata.NewMetadata(o.context, o.MetadataMetadataHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/project/{projectName}/stage/{stageName}/service/{serviceName}/evaluation"] = evaluation.NewTriggerEvaluation(o.context, o.EvaluationTriggerEvaluationHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *KeptnAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *KeptnAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *KeptnAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *KeptnAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *KeptnAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}

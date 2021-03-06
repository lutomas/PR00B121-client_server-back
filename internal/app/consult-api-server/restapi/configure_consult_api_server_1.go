// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"github.com/Sirupsen/logrus"
	"github.com/lutomas/PR00B121-client_server-back/internal/app/consult-api-server/handler"
	"github.com/rs/cors"
	"github.com/urfave/negroni"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/lutomas/PR00B121-client_server-back/internal/app/consult-api-server/restapi/operations"
	"github.com/lutomas/PR00B121-client_server-back/internal/app/consult-api-server/restapi/operations/consultation"
)

func configureFlags(api *operations.ConsultAPIServerAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.ConsultAPIServerAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	//TODO: Uncomment 4
	dbHandler := handler.NewDBHandler("localhost", "PR00B121")

	api.ConsultationAddConsultationHandler = consultation.AddConsultationHandlerFunc(func(params consultation.AddConsultationParams) middleware.Responder {
		return middleware.NotImplemented("operation consultation.AddConsultation has not yet been implemented")
		////TODO: Uncomment 3
		//return handler.ConsultationAddConsultationHandler(&params)
		//TODO: Uncomment 4 (comment-out 3)
		return dbHandler.ConsultationAddConsultationHandler(&params)
	})

	//TODO: Uncomment 5
	api.ConsultationGetConsultationsHandler = consultation.GetConsultationsHandlerFunc(func(params consultation.GetConsultationsParams) middleware.Responder {
		return dbHandler.ConsultationGetConsultationsHandler(&params)
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
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	//return handler

	// TODO: Uncomment 1
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logrus.SetLevel(logrus.DebugLevel)

	handleLogging := negroni.New()
	negroni.LoggerDefaultFormat = "negroni | {{.Status}} | \t {{.Duration}} | {{.Hostname}} | {{.Method}} {{.Path}}"
	logger := negroni.NewLogger()
	logger.ALogger = logrus.StandardLogger() // Use standard logrus logger
	handleLogging.Use(logger)

	handleLogging.UseHandler(handler)

	//return handleLogging

	// TODO: Uncomment 2
	handleCORS := cors.Default().Handler

	return handleCORS(handleLogging)
}

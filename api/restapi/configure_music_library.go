// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"main/api/restapi/operations"
)

//go:generate swagger generate server --target ..\..\..\music-lib --name MusicLibrary --spec ..\..\swagger.yaml --model-package ./internal/models --server-package ./api/restapi --principal interface{} --exclude-main

func configureFlags(api *operations.MusicLibraryAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.MusicLibraryAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.DeleteSongsHandler == nil {
		api.DeleteSongsHandler = operations.DeleteSongsHandlerFunc(func(params operations.DeleteSongsParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.DeleteSongs has not yet been implemented")
		})
	}
	if api.GetSongsHandler == nil {
		api.GetSongsHandler = operations.GetSongsHandlerFunc(func(params operations.GetSongsParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetSongs has not yet been implemented")
		})
	}
	if api.GetSongsLyricsHandler == nil {
		api.GetSongsLyricsHandler = operations.GetSongsLyricsHandlerFunc(func(params operations.GetSongsLyricsParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetSongsLyrics has not yet been implemented")
		})
	}
	if api.PatchSongsHandler == nil {
		api.PatchSongsHandler = operations.PatchSongsHandlerFunc(func(params operations.PatchSongsParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PatchSongs has not yet been implemented")
		})
	}
	if api.PostSongsHandler == nil {
		api.PostSongsHandler = operations.PostSongsHandlerFunc(func(params operations.PostSongsParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostSongs has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

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
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}

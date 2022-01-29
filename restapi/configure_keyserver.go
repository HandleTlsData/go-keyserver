// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"encoding/hex"
	"fmt"
	"hash"
	gDB "keyserver/database"
	"keyserver/models"
	"net/http"
	"time"

	"crypto/sha256"
	"keyserver/restapi/operations"
	"keyserver/restapi/operations/single_key_operation"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

func stringHasher(algorithm hash.Hash, text string) string {
	algorithm.Write([]byte(text))
	return hex.EncodeToString(algorithm.Sum(nil))
}

func SHA256(text string) string {
	algorithm := sha256.New()
	return stringHasher(algorithm, text)
}

func CreateKeyImpl(params single_key_operation.CreateKeyParams) middleware.Responder {
	var newEntity models.KeyData
	requestData := *params.Body
	newEntity.CreatedBy = requestData.CreatedBy
	newEntity.Expire = requestData.Expire
	dt := time.Now()
	unencStr := fmt.Sprintf("%s%s%s", newEntity.CreatedBy, requestData.Expire, dt.Format("01-02-2006 15:04:05"))
	encStr := SHA256(unencStr)
	newEntity.Data = encStr
	newEntity.Valid = true
	err := gDB.Create(newEntity)
	if err != nil {
		return single_key_operation.NewCreateKeyInternalServerError()
	}
	var responsePayload []*models.KeyData
	responsePayload = append(responsePayload, &newEntity)
	return single_key_operation.NewCreateKeyOK().WithPayload(responsePayload)
}

func GetKeyImpl(params single_key_operation.KeyGetParams) middleware.Responder {
	entity, err := gDB.Get(params.StrKey)
	if err != nil {
		if err.Error() == gDB.StrNotFound {
			single_key_operation.NewKeyGetNotFound()
		}
		single_key_operation.NewKeyGetInternalServerError()
	}
	var responsePayload []*models.KeyData
	responsePayload = append(responsePayload, &entity)
	return single_key_operation.NewKeyGetOK().WithPayload(responsePayload)
}

//go:generate swagger generate server --target ..\..\test-server --name Keyserver --spec ..\swagger\swagger.yml --principal interface{}

func configureFlags(api *operations.KeyserverAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.KeyserverAPI) http.Handler {
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

	api.SingleKeyOperationCreateKeyHandler = single_key_operation.CreateKeyHandlerFunc(CreateKeyImpl)
	api.SingleKeyOperationKeyGetHandler = single_key_operation.KeyGetHandlerFunc(GetKeyImpl)

	if api.SingleKeyOperationCreateKeyHandler == nil {
		api.SingleKeyOperationCreateKeyHandler = single_key_operation.CreateKeyHandlerFunc(func(params single_key_operation.CreateKeyParams) middleware.Responder {
			return middleware.NotImplemented("operation single_key_operation.CreateKey has not yet been implemented")
		})
	}
	if api.SingleKeyOperationKeyGetHandler == nil {
		api.SingleKeyOperationKeyGetHandler = single_key_operation.KeyGetHandlerFunc(func(params single_key_operation.KeyGetParams) middleware.Responder {
			return middleware.NotImplemented("operation single_key_operation.KeyGet has not yet been implemented")
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

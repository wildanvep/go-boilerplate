// Package oapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package oapi

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
	strictecho "github.com/oapi-codegen/runtime/strictmiddleware/echo"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// CreateProfile defines model for CreateProfile.
type CreateProfile struct {
	Dob   ZeroableTime   `json:"dob,omitempty"`
	Email ZeroableString `json:"email,omitempty"`
	Name  ZeroableString `json:"name,omitempty"`
	Nin   ZeroableString `json:"nin,omitempty"`
	Phone ZeroableString `json:"phone,omitempty"`
}

// Profile defines model for Profile.
type Profile struct {
	Dob      ZeroableTime   `json:"dob,omitempty"`
	Email    ZeroableString `json:"email,omitempty"`
	Id       UUID           `json:"id,omitempty"`
	Name     ZeroableString `json:"name,omitempty"`
	Nin      ZeroableString `json:"nin,omitempty"`
	Phone    ZeroableString `json:"phone,omitempty"`
	TenantId UUID           `json:"tenant_id,omitempty"`
}

// UUID defines model for UUID.
type UUID = openapi_types.UUID

// ZeroableString defines model for ZeroableString.
type ZeroableString = string

// ZeroableTime defines model for ZeroableTime.
type ZeroableTime = time.Time

// PostProfileParams defines parameters for PostProfile.
type PostProfileParams struct {
	Validate *bool `form:"validate,omitempty" json:"validate,omitempty"`
}

// GetProfileByNameAndEmailParams defines parameters for GetProfileByNameAndEmail.
type GetProfileByNameAndEmailParams struct {
	Name  string `form:"name" json:"name"`
	Email string `form:"email" json:"email"`
}

// PostProfileJSONRequestBody defines body for PostProfile for application/json ContentType.
type PostProfileJSONRequestBody = CreateProfile

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// create profile
	// (POST /tenants/{tenant-id}/profiles)
	PostProfile(ctx echo.Context, tenantId UUID, params PostProfileParams) error
	// get profile by name and email
	// (GET /tenants/{tenant-id}/profiles/search)
	GetProfileByNameAndEmail(ctx echo.Context, tenantId UUID, params GetProfileByNameAndEmailParams) error
	// get profile
	// (GET /tenants/{tenant-id}/profiles/{profile-id})
	GetProfile(ctx echo.Context, tenantId UUID, profileId UUID) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// PostProfile converts echo context to params.
func (w *ServerInterfaceWrapper) PostProfile(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "tenant-id" -------------
	var tenantId UUID

	err = runtime.BindStyledParameterWithOptions("simple", "tenant-id", ctx.Param("tenant-id"), &tenantId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter tenant-id: %s", err))
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params PostProfileParams
	// ------------- Optional query parameter "validate" -------------

	err = runtime.BindQueryParameter("form", true, false, "validate", ctx.QueryParams(), &params.Validate)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter validate: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostProfile(ctx, tenantId, params)
	return err
}

// GetProfileByNameAndEmail converts echo context to params.
func (w *ServerInterfaceWrapper) GetProfileByNameAndEmail(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "tenant-id" -------------
	var tenantId UUID

	err = runtime.BindStyledParameterWithOptions("simple", "tenant-id", ctx.Param("tenant-id"), &tenantId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter tenant-id: %s", err))
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params GetProfileByNameAndEmailParams
	// ------------- Required query parameter "name" -------------

	err = runtime.BindQueryParameter("form", true, true, "name", ctx.QueryParams(), &params.Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter name: %s", err))
	}

	// ------------- Required query parameter "email" -------------

	err = runtime.BindQueryParameter("form", true, true, "email", ctx.QueryParams(), &params.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter email: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetProfileByNameAndEmail(ctx, tenantId, params)
	return err
}

// GetProfile converts echo context to params.
func (w *ServerInterfaceWrapper) GetProfile(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "tenant-id" -------------
	var tenantId UUID

	err = runtime.BindStyledParameterWithOptions("simple", "tenant-id", ctx.Param("tenant-id"), &tenantId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter tenant-id: %s", err))
	}

	// ------------- Path parameter "profile-id" -------------
	var profileId UUID

	err = runtime.BindStyledParameterWithOptions("simple", "profile-id", ctx.Param("profile-id"), &profileId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter profile-id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetProfile(ctx, tenantId, profileId)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST(baseURL+"/tenants/:tenant-id/profiles", wrapper.PostProfile)
	router.GET(baseURL+"/tenants/:tenant-id/profiles/search", wrapper.GetProfileByNameAndEmail)
	router.GET(baseURL+"/tenants/:tenant-id/profiles/:profile-id", wrapper.GetProfile)

}

type PostProfileRequestObject struct {
	TenantId UUID `json:"tenant-id"`
	Params   PostProfileParams
	Body     *PostProfileJSONRequestBody
}

type PostProfileResponseObject interface {
	VisitPostProfileResponse(w http.ResponseWriter) error
}

type PostProfile201JSONResponse Profile

func (response PostProfile201JSONResponse) VisitPostProfileResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

type PostProfile400Response struct {
}

func (response PostProfile400Response) VisitPostProfileResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type GetProfileByNameAndEmailRequestObject struct {
	TenantId UUID `json:"tenant-id"`
	Params   GetProfileByNameAndEmailParams
}

type GetProfileByNameAndEmailResponseObject interface {
	VisitGetProfileByNameAndEmailResponse(w http.ResponseWriter) error
}

type GetProfileByNameAndEmail200JSONResponse Profile

func (response GetProfileByNameAndEmail200JSONResponse) VisitGetProfileByNameAndEmailResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetProfileByNameAndEmail404Response struct {
}

func (response GetProfileByNameAndEmail404Response) VisitGetProfileByNameAndEmailResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type GetProfileRequestObject struct {
	TenantId  UUID `json:"tenant-id"`
	ProfileId UUID `json:"profile-id"`
}

type GetProfileResponseObject interface {
	VisitGetProfileResponse(w http.ResponseWriter) error
}

type GetProfile200JSONResponse Profile

func (response GetProfile200JSONResponse) VisitGetProfileResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetProfile404Response struct {
}

func (response GetProfile404Response) VisitGetProfileResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// create profile
	// (POST /tenants/{tenant-id}/profiles)
	PostProfile(ctx context.Context, request PostProfileRequestObject) (PostProfileResponseObject, error)
	// get profile by name and email
	// (GET /tenants/{tenant-id}/profiles/search)
	GetProfileByNameAndEmail(ctx context.Context, request GetProfileByNameAndEmailRequestObject) (GetProfileByNameAndEmailResponseObject, error)
	// get profile
	// (GET /tenants/{tenant-id}/profiles/{profile-id})
	GetProfile(ctx context.Context, request GetProfileRequestObject) (GetProfileResponseObject, error)
}

type StrictHandlerFunc = strictecho.StrictEchoHandlerFunc
type StrictMiddlewareFunc = strictecho.StrictEchoMiddlewareFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// PostProfile operation middleware
func (sh *strictHandler) PostProfile(ctx echo.Context, tenantId UUID, params PostProfileParams) error {
	var request PostProfileRequestObject

	request.TenantId = tenantId
	request.Params = params

	var body PostProfileJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostProfile(ctx.Request().Context(), request.(PostProfileRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostProfile")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostProfileResponseObject); ok {
		return validResponse.VisitPostProfileResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetProfileByNameAndEmail operation middleware
func (sh *strictHandler) GetProfileByNameAndEmail(ctx echo.Context, tenantId UUID, params GetProfileByNameAndEmailParams) error {
	var request GetProfileByNameAndEmailRequestObject

	request.TenantId = tenantId
	request.Params = params

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetProfileByNameAndEmail(ctx.Request().Context(), request.(GetProfileByNameAndEmailRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetProfileByNameAndEmail")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetProfileByNameAndEmailResponseObject); ok {
		return validResponse.VisitGetProfileByNameAndEmailResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetProfile operation middleware
func (sh *strictHandler) GetProfile(ctx echo.Context, tenantId UUID, profileId UUID) error {
	var request GetProfileRequestObject

	request.TenantId = tenantId
	request.ProfileId = profileId

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetProfile(ctx.Request().Context(), request.(GetProfileRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetProfile")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetProfileResponseObject); ok {
		return validResponse.VisitGetProfileResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/9xVwW7UMBD9lWjgmDRbugfkWwsI7QVVgl5AFfLGs7uGxOPaTkUU5d/ROMluQ6G7bREU",
	"TrEm9nue914mLRRUWTJoggfRgi82WMm4fOVQBjx3tNIlcsE6suiCxvha0ZIfzx2uQMCzfIeTDyD5R3Qk",
	"lyV+0BVClwJWUpeHHnofnDZrPmZkhQ84pc39D9kNmXtzdV0KT0AmrfadubhYvP4XBE0hoJEmfD60JXYg",
	"rkQLK3KVDCCgrrWCFEJjEQT4HjuFb9maMi5m/qu2Gdmgycgys6RNQAciuBq7FH64lGgfjRT9vXlDJQNm",
	"gasPvSY3rs2KGLbUBRofGXqDYcE7jSwhhdqVIGATghV5XlIhyw35EKXWgXM7Jjg5PV9ACtfovCYDAo6P",
	"Zkcz3kgWjbQaBJzEUgpWhk0Med775fO2X2RadbntAeMGy2SiZQwnuZOFYkryYfxwGM3JCgM6D+JTC5y2",
	"yABjYGELDik4vKq1Q9ULkQ6D68C4pAP8VY2u2eFfy1KzJ3ATbjBmSVSiNNB1lz05+nBGquEtBZmAJjYo",
	"rS11EVvMv3gyu5m672rTcRudnfYYC96S8b2mL2bHv418QqvQF07HwHEi66JA7zkB89ksjrTJ+6VUyaAH",
	"Kwsei9rp0LCJrJWvq0q6BgQUscPEbrnSu4OTe5Su2DDlGn8Sn7c4pueseScrPDXqTRydTyFL8XEX9PSD",
	"/yUODh0dDnR5KyezP56T+e2cGArJimqj9qRkjWGMSLJsEpYhkUYlvRD7Q9MOK64eEJ2/E5Yp/O7Gj8X/",
	"f7zvaTy669GX3S/Mi3z0W7ycz0+AdZ2+3v7ihg2X3fcAAAD//9HC2d7qCgAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}

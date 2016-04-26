package middleware

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type (
	// BasicAuthConfig defines the config for HTTP basic auth middleware.
	BasicAuthConfig struct {
		// AuthFunc is the function to validate basic auth credentials.
		AuthFunc BasicAuthFunc
	}

	// BasicAuthFunc defines a function to validate basic auth credentials.
	BasicAuthFunc func(string, string) bool

	// JWTAuthConfig defines the config for JWT auth middleware.
	JWTAuthConfig struct {
		// SigningKey is the key to validate token.
		// Required.
		SigningKey []byte

		// SigningMethod is used to check token signing method.
		// Optional, with default value as `HS256`.
		SigningMethod string

		// ContextKey is the key to be used for storing user information from the
		// token into context.
		// Optional, with default value as `user`.
		ContextKey string

		// Extractor is a function that extracts token from the request
		// Optional, with default values as `JWTFromHeader`.
		Extractor JWTExtractor
	}

	// JWTExtractor defines a function that takes `echo.Context` and returns either
	// a token or an error.
	JWTExtractor func(echo.Context) (string, error)
)

const (
	basic  = "Basic"
	bearer = "Bearer"
)

// Algorithims
const (
	AlgorithmHS256 = "HS256"
)

var (
	// DefaultBasicAuthConfig is the default basic auth middleware config.
	DefaultBasicAuthConfig = BasicAuthConfig{}

	// DefaultJWTAuthConfig is the default JWT auth middleware config.
	DefaultJWTAuthConfig = JWTAuthConfig{
		SigningMethod: AlgorithmHS256,
		ContextKey:    "user",
		Extractor:     JWTFromHeader,
	}
)

// BasicAuth returns an HTTP basic auth middleware.
//
// For valid credentials it calls the next handler.
// For invalid credentials, it sends "401 - Unauthorized" response.
// For empty or invalid `Authorization` header, it sends "400 - Bad Request" response.
func BasicAuth(fn BasicAuthFunc) echo.MiddlewareFunc {
	c := DefaultBasicAuthConfig
	c.AuthFunc = fn
	return BasicAuthWithConfig(c)
}

// BasicAuthWithConfig returns an HTTP basic auth middleware from config.
// See `BasicAuth()`.
func BasicAuthWithConfig(config BasicAuthConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			auth := c.Request().Header().Get(echo.HeaderAuthorization)
			l := len(basic)

			if len(auth) > l+1 && auth[:l] == basic {
				b, err := base64.StdEncoding.DecodeString(auth[l+1:])
				if err != nil {
					return err
				}
				cred := string(b)
				for i := 0; i < len(cred); i++ {
					if cred[i] == ':' {
						// Verify credentials
						if config.AuthFunc(cred[:i], cred[i+1:]) {
							return next(c)
						}
						c.Response().Header().Set(echo.HeaderWWWAuthenticate, basic+" realm=Restricted")
						return echo.ErrUnauthorized
					}
				}
			}
			return echo.NewHTTPError(http.StatusBadRequest, "invalid basic-auth authorization header="+auth)
		}
	}
}

// JWTAuth returns a JSON Web Token (JWT) auth middleware.
//
// For valid token, it sets the user in context and calls next handler.
// For invalid token, it sends "401 - Unauthorized" response.
// For empty or invalid `Authorization` header, it sends "400 - Bad Request".
//
// See https://jwt.io/introduction
func JWTAuth(key []byte) echo.MiddlewareFunc {
	c := DefaultJWTAuthConfig
	c.SigningKey = key
	return JWTAuthWithConfig(c)
}

// JWTAuthWithConfig returns a JWT auth middleware from config.
// See `JWTAuth()`.
func JWTAuthWithConfig(config JWTAuthConfig) echo.MiddlewareFunc {
	// Defaults
	if config.SigningKey == nil {
		panic("jwt middleware requires signing key")
	}
	if config.SigningMethod == "" {
		config.SigningMethod = DefaultJWTAuthConfig.SigningMethod
	}
	if config.ContextKey == "" {
		config.ContextKey = DefaultJWTAuthConfig.ContextKey
	}
	if config.Extractor == nil {
		config.Extractor = DefaultJWTAuthConfig.Extractor
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			auth, err := config.Extractor(c)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			token, err := jwt.Parse(auth, func(t *jwt.Token) (interface{}, error) {
				// Check the signing method
				if t.Method.Alg() != config.SigningMethod {
					return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
				}
				return config.SigningKey, nil

			})
			if err == nil && token.Valid {
				// Store user information from token into context.
				c.Set(config.ContextKey, token)
				return next(c)
			}
			return echo.ErrUnauthorized
		}
	}
}

// JWTFromHeader is a `JWTExtractor` that extracts token from the `Authorization` request
// header.
func JWTFromHeader(c echo.Context) (string, error) {
	auth := c.Request().Header().Get(echo.HeaderAuthorization)
	l := len(bearer)
	if len(auth) > l+1 && auth[:l] == bearer {
		return auth[l+1:], nil
	}
	return "", echo.NewHTTPError(http.StatusBadRequest, "invalid jwt authorization header="+auth)
}

// JWTFromQuery returns a `JWTExtractor` that extracts token from the provided query
// parameter.
func JWTFromQuery(param string) JWTExtractor {
	return func(c echo.Context) (string, error) {
		return c.QueryParam(param), nil
	}
}

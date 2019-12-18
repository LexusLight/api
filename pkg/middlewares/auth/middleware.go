package auth

import (
	"errors"
	"time"
	"github.com/deissh/osu-api-server/pkg"
	"github.com/deissh/osu-api-server/pkg/services/oauth"
	"github.com/labstack/echo/v4"
)

// Middleware check access_token
func Middleware(requireScopes []string, requireRoles []string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			extractor := keyFromHeader("Authorization")

			key, err := extractor(c)
			if err != nil {
				return pkg.NewHTTPError(401, "auth_token_required", "Need auth token in headers")
			}

			token, err := oauth.ValidateOAuthToken(key)
			if err != nil {
				return err
			}

			pkg.Rb.SAdd("online_users", token.UserID)

			c.Set("uset_token", token)
			return next(c)
		}
	}
}

// keyFromHeader returns a `keyExtractor` that extracts key from the request header.
func keyFromHeader(header string) func(echo.Context) (string, error) {
	return func(c echo.Context) (string, error) {
		auth := c.Request().Header.Get(header)
		if auth == "" {
			return "", errors.New("missing key in request header")
		}
		if header == echo.HeaderAuthorization {
			l := len("Bearer")
			if len(auth) > l+1 && auth[:l] == "Bearer" {
				return auth[l+1:], nil
			}
			return "", errors.New("invalid key in the request header")
		}
		return auth, nil
	}
}
package middleware

import (
	"log"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/vincentkdeli/vinance-backend/entity"
	auth "github.com/vincentkdeli/vinance-backend/pkg/authenticator"
	"github.com/vincentkdeli/vinance-backend/pkg/errors"
)

const authorizationHeaderKey = "Authorization"

type AuthConfig struct {
	AllowAuthorizationHeader bool
}

type authenticateFn func(c echo.Context, authenticator *auth.Authenticator) (tokenClaims *entity.TokenClaims, err error)

func Authentication(authenticator *auth.Authenticator, config AuthConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var tokenClaims *entity.TokenClaims
			var err error

			var fns []authenticateFn
			if config.AllowAuthorizationHeader {
				fns = append(fns, authenticateAuthorizationHeader)
			}

			for _, fn := range fns {
				if tokenClaims == nil {
					tokenClaims, err = fn(c, authenticator)
					if err != nil {
						return err
					}
				}
			}

			if tokenClaims == nil {
				return errors.ErrUnauthorized(nil, "Missing Authorization header.")
			}

			log.Println("Authentication successful for: ", tokenClaims.UserEmail)
			c.Set("member_id", tokenClaims.UserID)
			c.Set("member_email", tokenClaims.UserEmail)

			return next(c)
		}
	}
}

func authenticateAuthorizationHeader(c echo.Context, authenticator *auth.Authenticator) (*entity.TokenClaims, error) {
	authHeader := c.Request().Header.Get(authorizationHeaderKey)
	if authHeader == "" {
		return nil, nil
	}

	splitToken := strings.Split(authHeader, "Bearer")
	if len(splitToken) != 2 {
		return nil, errors.ErrUnauthorized(nil, "Invalid Authorization header format.")
	}

	accessToken := strings.TrimSpace(splitToken[1])

	tokenClaims, err := authenticator.ParseToken(accessToken)
	if err != nil {
		return nil, errors.ErrUnauthorized(err, "Invalid Authorization header.")
	}

	return tokenClaims, nil
}

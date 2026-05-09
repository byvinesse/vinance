package authenticator

import (
	"fmt"
	"time"

	"github.com/byvinesse/vinance-backend/config"
	"github.com/byvinesse/vinance-backend/entity"
	"github.com/byvinesse/vinance-backend/pkg/errors"
	"github.com/golang-jwt/jwt"
)

type Authenticator struct {
	config config.Auth
}

func NewAuthenticator(config config.Auth) *Authenticator {
	return &Authenticator{
		config: config,
	}
}

func (s *Authenticator) GenerateJwtToken(user entity.User) (string, time.Time, error) {
	tokenExpiresAt := time.Now().Add(s.config.AccessTokenDuration).UTC()

	tokenClaims := &entity.TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: tokenExpiresAt.Unix(),
		},
		UserID:    user.ID,
		UserEmail: user.Email,
	}

	unsignedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
	token, err := unsignedToken.SignedString([]byte(s.config.JwtKey))
	if err != nil {
		return "", time.Time{}, fmt.Errorf("#generateJwtToken token.SignedString: %w", err)
	}

	return token, tokenExpiresAt, nil
}

func (s *Authenticator) ParseToken(token string) (*entity.TokenClaims, error) {
	tokenClaims := new(entity.TokenClaims)

	jwtTokenObj, err := jwt.ParseWithClaims(token, tokenClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.config.JwtKey), nil
	})

	if err != nil {
		return nil, errors.ErrUnauthorized(err, "Invalid access token.")
	}

	if !jwtTokenObj.Valid {
		return nil, errors.ErrUnauthorized(fmt.Errorf("invalid JWT token"), "Invalid access token.")
	}

	return tokenClaims, nil
}

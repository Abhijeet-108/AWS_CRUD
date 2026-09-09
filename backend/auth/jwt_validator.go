package auth

import (
	"errors"

	"github.com/MicahParks/keyfunc/v3"
	"github.com/golang-jwt/jwt/v5"
)

type CognitoClaims struct {
	TokenUse string `json:"token_use"`
	ClientID string `json:"client_id"`
	Username string `json:"username"`
	Name     string `json:"name"`

	jwt.RegisteredClaims
}

type JWTValidator struct {
	JWKS     keyfunc.Keyfunc
	Issuer   string
	ClientID string
}

func NewJWTValidator(jwksURL, issuer, clientID string) (*JWTValidator, error) {
	jwks, err := keyfunc.NewDefault([]string{jwksURL})

	if err != nil {
		return nil, err
	}

	return &JWTValidator{
		JWKS:     jwks,
		Issuer:   issuer,
		ClientID: clientID,
	}, nil
}

func (v *JWTValidator) ValidateToken(tokenString string) (*CognitoClaims, error) {
	claims := &CognitoClaims{}

	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		v.JWKS.Keyfunc,
		jwt.WithValidMethods([]string{"RS256"}),
		jwt.WithIssuer(v.Issuer),
		jwt.WithExpirationRequired(),
	)

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	if claims.TokenUse != "access" {
		return nil, errors.New("token is not an access token")
	}

	if claims.ClientID != v.ClientID {
		return nil, errors.New("invalid client_id")
	}

	return claims, nil

}

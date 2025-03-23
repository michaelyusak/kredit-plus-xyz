package utils

import (
	"encoding/json"
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/michaelyusak/kredit-plus-xyz/config"
)

type JwtCustomClaims struct {
	AccountId int64  `json:"user_id"`
	Role      string `json:"role"`
}

type JWTHelper interface {
	CreateAndSign(customClaims JwtCustomClaims, expiredAt int64) (*string, error)
	ParseAndVerify(signed string) (*JwtCustomClaims, error)
}

type jwtHelperImpl struct {
	config config.JwtConfig
	Method *jwt.SigningMethodHMAC
}

func NewJWTHelper(config config.JwtConfig) *jwtHelperImpl {
	return &jwtHelperImpl{
		config: config,
	}
}

func (h *jwtHelperImpl) CreateAndSign(customClaims JwtCustomClaims, expiredAt int64) (*string, error) {
	customClaimsJsonBytes, err := json.Marshal(customClaims)
	if err != nil {
		return nil, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":  h.config.Issuer,
		"exp":  expiredAt,
		"data": string(customClaimsJsonBytes),
	})

	signed, err := token.SignedString([]byte(h.config.Key))
	if err != nil {
		return nil, err
	}

	return &signed, nil
}

func (h *jwtHelperImpl) ParseAndVerify(signed string) (*JwtCustomClaims, error) {
	token, err := jwt.Parse(signed, func(token *jwt.Token) (interface{}, error) {
		return []byte(h.config.Key), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}),
		jwt.WithIssuer(h.config.Issuer),
		jwt.WithExpirationRequired(),
	)
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenInvalidIssuer) {
			return nil, nil
		}

		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}

	customClaims := JwtCustomClaims{}

	data := claims["data"].(string)

	err = json.Unmarshal([]byte(data), &customClaims)
	if err != nil {
		return nil, err
	}

	return &customClaims, nil
}

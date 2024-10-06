package authservice

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	OwnerEnum = iota + 1
	VetEnum
)

const (
	jwtSalt    = "solyanovo"
	signingKey = "entering"
	tokenTTL   = time.Hour
)

type Payload struct {
	jwt.StandardClaims
	UserId   uint
	FullName string
	Role     uint
}

func (s *AuthService) CreateToken(id uint, fullame string, isVet bool) (string, error) {
	op := "AuthService.CreateToken"
	log := s.log.WithField("op", op)

	payload := Payload{
		UserId:   id,
		FullName: fullame,
	}

	if isVet {
		payload.Role = VetEnum
	} else {
		payload.Role = OwnerEnum
	}

	payload.ExpiresAt = time.Now().Add(tokenTTL).Unix()
	payload.IssuedAt = time.Now().Unix()

	log.Debug("creating token")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(token string) (Payload, error) {
	tok, err := jwt.ParseWithClaims(token, &Payload{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return Payload{}, err
	}

	claims, ok := tok.Claims.(*Payload)
	if !ok {
		return Payload{}, errors.New("token claims are not of type *tokenClaims")
	}

	return *claims, nil
}

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

func (s *AuthService) CreateToken(email, passwordHash string, isVet bool) (string, error) {
	op := "AuthService.CreateToken"
	log := s.log.WithField("op", op)

	log.Debug("getting user")

	var payload Payload
	var err error // wtf))

	// Из-за ублюдской бд приходится делать разделение... Пользователи увы, в разных таблицах
	if isVet {
		payload, err = s.createVetPayload(email, passwordHash)
	} else {
		payload, err = s.createOwnerPayload(email, passwordHash)
	}
	if err != nil {
		return "", err
	}

	payload.ExpiresAt = time.Now().Add(tokenTTL).Unix()
	payload.IssuedAt = time.Now().Unix()

	log.Debug("creating token")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) createVetPayload(email, passwordHash string) (Payload, error) {
	vet, err := s.storage.GetVet(email, passwordHash)
	if err != nil {
		return Payload{}, err
	}
	return Payload{
		UserId:   vet.ID,
		FullName: vet.User.FullName,
		Role:     VetEnum,
	}, nil
}

func (s *AuthService) createOwnerPayload(email, passwordHash string) (Payload, error) {
	owner, err := s.storage.GetOwner(email, passwordHash)
	if err != nil {
		return Payload{}, err
	}
	return Payload{
		UserId:   owner.ID,
		FullName: owner.User.FullName,
		Role:     OwnerEnum,
	}, nil
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

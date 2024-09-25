package authservice

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/vet-clinic-back/sso-service/internal/logging"
	"github.com/vet-clinic-back/sso-service/internal/models"
	"github.com/vet-clinic-back/sso-service/internal/storage"
)

const (
	salt       = "solyanovo"
	signingKey = "entering"
	tokenTTL   = time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int
}

type AuthService struct {
	log     *logging.Logger
	storage storage.Auth
}

func New(log *logging.Logger, storage storage.Auth) *AuthService {
	return &AuthService{log: log, storage: storage}
}

func (s *AuthService) CreateToken(email, passwordHash string) (string, error) {
	op := "AuthService.CreateToken"
	log := s.log.WithField("op", op)

	log.Debug("getting user")
	user, err := s.storage.GetUser(email, passwordHash)
	if err != nil {
		return "", err
	}

	log.Debug("creating token")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) CreateUser(user models.User) (int, error) {
	return s.storage.CreateUser(user)
}

func (s *AuthService) ParseToken(token string) (int, error) {
	tok, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := tok.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

// func generatePasswordHash(password string) string {
// 	hash := sha1.New()
// 	hash.Write([]byte(password))

// 	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
// }

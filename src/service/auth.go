package service

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// jwt service
type JWTService interface {
	GenerateToken(username, password string, uid int64, role Role) string
	ValidateToken(token string) (*jwt.Token, error)
}
type authCustomClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     Role   `json:"role"`
	ID       int64  `json:"uid"`
	jwt.StandardClaims
}

type jwtServices struct {
	secretKey string
	issure    string
}

// auth-jwt
func JWTAuthService() JWTService {
	return &jwtServices{
		secretKey: getSecretKey(),
		issure:    "jiawei",
	}
}

func getSecretKey() string {
	secret := os.Getenv("SECRET")
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func (service *jwtServices) GenerateToken(username, password string, uid int64, role Role) string {
	claims := &authCustomClaims{
		username,
		password,
		role,
		uid,
		jwt.StandardClaims{
			// ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:   service.issure,
			IssuedAt: time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (service *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("invalid token %+v", token.Header["alg"])

		}
		return []byte(service.secretKey), nil
	})

}

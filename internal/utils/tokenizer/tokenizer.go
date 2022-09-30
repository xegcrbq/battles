package tokenizer

import (
	"battles/internal/utils/env"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"os"
	"sync"
	"time"
)

var once sync.Once
var tknz *Tokenizer = nil

type Tokenizer struct {
	jwtKey []byte
}
type DataClaims struct {
	Data string
	jwt.StandardClaims
}

func (c DataClaims) Valid() error {
	if c.ExpiresAt < time.Now().Unix() {
		return jwt.ValidationError{Errors: jwt.ValidationErrorExpired}
	} else {
		return nil
	}
}
func Get() *Tokenizer {
	env.InitEnv()
	once.Do(func() {
		tknz = &Tokenizer{
			jwtKey: []byte(os.Getenv("jwtKey")),
		}
	})
	return tknz
}

func (t *Tokenizer) NewJWTCookie(name, data string, expirationTime time.Time) *fiber.Cookie {
	dc := &DataClaims{
		Data: data,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	signedToken, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, dc).SignedString(t.jwtKey)
	return &fiber.Cookie{
		Name:    name,
		Value:   signedToken,
		Expires: expirationTime,
	}
}
func (t *Tokenizer) NewJWTCookieHTTPOnly(name, data, path string, expirationTime time.Time) (*fiber.Cookie, error) {
	dc := &DataClaims{
		Data: data,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	signedToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, dc).SignedString(t.jwtKey)
	return &fiber.Cookie{
		Name:     name,
		Value:    signedToken,
		Expires:  expirationTime,
		HTTPOnly: true,
		Path:     path,
	}, err
}
func (t *Tokenizer) ParseDataClaims(data string) (*DataClaims, *jwt.Token, error) {
	dc := &DataClaims{}
	tkn, err := jwt.ParseWithClaims(data, dc, func(token *jwt.Token) (interface{}, error) {
		return t.jwtKey, nil
	})
	return dc, tkn, err
}

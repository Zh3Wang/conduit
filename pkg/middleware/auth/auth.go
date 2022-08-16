package auth

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/golang-jwt/jwt"
	"time"
)

type Claims struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

const TokenExpire = time.Hour * 24

func GenerateJwtToken(secret, username, email string) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		UserName: username,
		Email:    email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpire).Unix(),
		},
	}).SignedString([]byte(secret))
}

func ParseJwtToken(secret, tokenString string) (Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return Claims{}, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return *claims, nil
	}
	return Claims{}, nil
}

func JWTAuthorization() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			//if tr, ok := transport.FromServerContext(ctx); ok {
			//	// Do something on entering
			//	header := tr.RequestHeader()
			//	defer func() {
			//		// Do something on exiting
			//	}()
			//}
			return handler(ctx, req)
		}
	}
}

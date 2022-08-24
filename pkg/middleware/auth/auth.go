package auth

import (
	errorPb "conduit/api/interface/v1"
	"context"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/golang-jwt/jwt"
	errorsv2 "github.com/pkg/errors"
	"strings"
	"time"
)

type Claims struct {
	UserName string `json:"username"`
	Id       int64  `json:"id"`
	jwt.StandardClaims
}

const TokenExpire = time.Hour * 24

// GenerateJwtToken 生成jwt token
func GenerateJwtToken(secret, username string, userId int64) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		UserName: username,
		Id:       userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpire).Unix(),
		},
	}).SignedString([]byte(secret))
}

// ParseJwtToken 解析jwt Token
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
	return Claims{}, errorsv2.New("token invalid")
}

type GlobalUserInfo struct {
	UserId   int64
	UserName string
	Email    string
}

// JWTAuthorization 中间件 - jwt鉴权
func JWTAuthorization(secret string) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				// Do something on entering
				header := tr.RequestHeader()
				defer func() {
					// Do something on exiting
				}()
				authHeader := header.Get("Authorization")
				authHeaders := strings.Split(authHeader, " ")
				if authHeader == "" || len(authHeaders) < 2 {
					return reply, errorPb.ErrorTokenInvalid("Authorization Header missed")
				}
				tokenString := authHeaders[1]
				claims, err := ParseJwtToken(secret, tokenString)
				if err != nil {
					return nil, errorPb.ErrorTokenInvalid("%s", err.Error())
				}
				// 元数据放入到ctx中
				ctx = WithContext(ctx, &GlobalUserInfo{
					UserId:   claims.Id,
					UserName: claims.UserName,
				})
			}
			return handler(ctx, req)
		}
	}
}

var GlobalUserInfoKey struct{}

func FromContext(ctx context.Context) *GlobalUserInfo {
	return ctx.Value(GlobalUserInfoKey).(*GlobalUserInfo)
}

func WithContext(ctx context.Context, user *GlobalUserInfo) context.Context {
	return context.WithValue(ctx, GlobalUserInfoKey, user)
}

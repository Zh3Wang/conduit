package auth

import (
	"conduit/api/interface/v1"
	errorPb "conduit/api/interface/v1"
	"context"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/golang-jwt/jwt"
	errorsv2 "github.com/pkg/errors"
	"strconv"
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
				if (authHeader == "" || len(authHeaders) < 2) && !OptionalAuthenticationOperation(tr.Operation()) {
					return reply, errorPb.ErrorTokenInvalid("Authorization Header missed")
				}
				if len(authHeaders) > 1 {
					tokenString := authHeaders[1]
					claims, err := ParseJwtToken(secret, tokenString)
					if err != nil {
						return nil, errorPb.ErrorTokenInvalid("%s", err.Error())
					}
					// 元数据放入到ctx中
					ctx = WithContext(ctx, claims.Id)
				}
			}
			return handler(ctx, req)
		}
	}
}

func OptionalAuthenticationOperation(op string) bool {
	var skip = map[string]struct{}{
		// optional authentication
		interfacePb.OperationConduitInterfaceGetProfile:   {},
		interfacePb.OperationConduitInterfaceListArticles: {},
	}
	_, ok := skip[op]
	return ok
}

const Uid = "x-md-global-uid"

func GetUserIdFromContext(ctx context.Context) int64 {
	if md, ok := metadata.FromServerContext(ctx); ok {
		extra := md.Get(Uid)
		if extra == "" {
			return getUserIdFromClientContext(ctx)
		}
		uid, _ := strconv.Atoi(extra)
		return int64(uid)
	}
	return 0
}

func getUserIdFromClientContext(ctx context.Context) int64 {
	if md, ok := metadata.FromClientContext(ctx); ok {
		extra := md.Get(Uid)
		uid, _ := strconv.Atoi(extra)
		return int64(uid)
	}
	return 0
}

func WithContext(ctx context.Context, user int64) context.Context {
	return metadata.AppendToClientContext(ctx, Uid, strconv.Itoa(int(user)))
}

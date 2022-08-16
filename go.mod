module conduit

go 1.16

require (
	github.com/envoyproxy/protoc-gen-validate v0.6.7
	github.com/go-kratos/kratos/contrib/registry/etcd/v2 v2.0.0-20220727050715-86eba9464615
	github.com/go-kratos/kratos/v2 v2.4.0
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/golang/protobuf v1.5.2
	github.com/google/wire v0.5.0
	github.com/pkg/errors v0.9.1
	go.etcd.io/etcd/client/v3 v3.5.4
	golang.org/x/crypto v0.0.0-20220722155217-630584e8d5aa
	golang.org/x/sys v0.0.0-20220811171246-fbc7d0a398ab // indirect
	google.golang.org/genproto v0.0.0-20220519153652-3a47de7e79bd
	google.golang.org/grpc v1.48.0
	google.golang.org/protobuf v1.28.0
	gorm.io/driver/mysql v1.3.5
	gorm.io/gorm v1.23.8
)

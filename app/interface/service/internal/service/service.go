package service

import (
	interfacePb "conduit/api/interface/v1"
	"conduit/app/interface/service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewConduitInterface)

type ConduitInterface struct {
	interfacePb.UnimplementedConduitInterfaceServer

	log *log.Helper
	uc  *biz.UserUsecase
	ac  *biz.ArticleUsecase
}

func NewConduitInterface(
	uc *biz.UserUsecase,
	ac *biz.ArticleUsecase,
	l log.Logger) *ConduitInterface {
	return &ConduitInterface{
		log: log.NewHelper(l),
		uc:  uc,
		ac:  ac,
	}
}

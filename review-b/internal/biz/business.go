package biz

import (
	"context"


	"github.com/go-kratos/kratos/v2/log"
)

type ReplyParam struct {
	ReviewID      int64                 
	StoreID       int64                             
	Content       string              
	PicInfo       string              
	VideoInfo     string
	Reason		  string
}


// GreeterRepo is a Greater repo.
type BusinessRepo interface {
	Reply(context.Context, *ReplyParam) (int64, error)
}

// GreeterUsecase is a Greeter usecase.
type BusinessUsecase struct {
	repo BusinessRepo
	log  *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewBusinessUsecase(repo BusinessRepo, logger log.Logger) *BusinessUsecase {
	return &BusinessUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *BusinessUsecase) CreateReply(ctx context.Context, param *ReplyParam) (int64, error) {
	uc.log.WithContext(ctx).Info("create reply", param)
	return uc.repo.Reply(ctx, param)
}

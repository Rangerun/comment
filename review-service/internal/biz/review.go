package biz

import (
	"context"
	v1 "review-service/api/review/v1"
	"review-service/internal/data/model"
	"review-service/pkg/snowflake"

	"github.com/go-kratos/kratos/v2/log"
)

// GreeterRepo is a Greater repo.
type ReviewRepo interface {
	SaveReview(context.Context, *model.ReviewInfo) (*model.ReviewInfo, error)
	//Update(context.Context, *Greeter) (*Greeter, error)
	GetReviewByOrderID(context.Context, int64) ([]*model.ReviewInfo, error)
	SaveReply(context.Context, *model.ReviewReplyInfo) (*model.ReviewReplyInfo, error)
	//ListByHello(context.Context, string) ([]*Greeter, error)
	//ListAll(context.Context) ([]*Greeter, error)
}

// GreeterUsecase is a Greeter usecase.
type ReviewUsecase struct {
	repo ReviewRepo
	log  *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewReviewUsecase(repo ReviewRepo, logger log.Logger) *ReviewUsecase {
	return &ReviewUsecase{repo: repo, log: log.NewHelper(logger)}
}


func (uc *ReviewUsecase) CreateReview(ctx context.Context, review *model.ReviewInfo) (*model.ReviewInfo, error) {
	uc.log.WithContext(ctx).Debugf("biz create req%v", review)
	// 数据校验
	reviews, err := uc.repo.GetReviewByOrderID(ctx, review.OrderID)
	if err != nil {
		return nil, v1.ErrorDbFailed("查询数据库数失败")
	}
	if (len(reviews) > 0) {
		return nil, v1.ErrorOrderHaveing("订单已存在%d", review.OrderID)
	}
	// 生成reviewId
	review.ReviewID = snowflake.GenID()
	// 查询订单和商品

	// 拼装数据入库

	return uc.repo.SaveReview(ctx, review)

}

func (uc *ReviewUsecase) CreateReply(ctx context.Context, param *ReplyParam) (*model.ReviewReplyInfo, error) {
	// 调用data 层创建回复

	reply := &model.ReviewReplyInfo{
		ReviewID: param.ReviewID,
		StoreID: param.StoreID,
		Content: param.Content,
		PicInfo: param.VideoInfo,
		VideoInfo: param.VideoInfo,
		ReplyID: snowflake.GenID(),
	}
	return uc.repo.SaveReply(ctx, reply)

}
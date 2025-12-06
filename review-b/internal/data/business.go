package data

import (
	"context"
	v1 "review-b/api/review/v1"
	"review-b/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type businessRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewBusinessRepo(data *Data, logger log.Logger) biz.BusinessRepo {
	return &businessRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *businessRepo) Reply(ctx context.Context, param *biz.ReplyParam) (int64, error) {
	// 之前都是写操作数据库 现在是走rpc调用
	r.log.WithContext(ctx).Info("param ", param)
	ret, err := r.data.rc.ReplyReview(ctx, &v1.ReplyReviewRequest{
		ReviewID: param.ReviewID,
		StoreID: param.StoreID,
		Content: param.Content,
		Reason: param.Reason,
		PicInfo: param.PicInfo,
		VideoInfo: param.VideoInfo,
	})
	r.log.WithContext(ctx).Debug("ret----", ret)
	if err != nil {
		return 0, err
	}
	return ret.GetReplyID(), nil
}


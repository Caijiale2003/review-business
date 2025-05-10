package data

import (
	"context"
	v1 "review-business/api/review/v1"
	"review-business/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type BusinessRepo struct {
	data *Data
	log  *log.Helper
}

func NewBusinessRepo(data *Data, logger log.Logger) biz.BusinessRepo {
	return &BusinessRepo{
		data: data,
		log:  log.NewHelper(logger),
	}

}

func (r *BusinessRepo) Reply(ctx context.Context, param *biz.ReplyParam) (int64, error) {
	res, err := r.data.rc.ReplyReview(ctx, &v1.ReplyReviewRequest{
		ReviewID:  param.ReviewID,
		StoreID:   param.StoreID,
		Content:   param.Content,
		PicInfo:   param.PicInfo,
		VideoInfo: param.VideoInfo,
	})
	if err != nil {
		return 0, err
	}
	return res.GetReplyID(), nil
}

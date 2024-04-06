package postreviewhandler

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
	postreviewiomodel "paradise-booking/modules/post_review/iomodel"
)

type PostReviewUseCase interface {
	CreatePostReview(ctx context.Context, data *postreviewiomodel.CreatePostReviewReq) error
	UpdatePostReview(ctx context.Context, data *postreviewiomodel.UpdatePostReviewReq) error
	ListPostReviewByAccountID(ctx context.Context, accountID int, paging *common.Paging) ([]*entities.PostReview, error)
	DeletePostReviewByID(ctx context.Context, postReviewID int) error
	GetPostReviewByID(ctx context.Context, postReviewID int) (*entities.PostReview, error)
}

type postReviewHandler struct {
	postReviewUC PostReviewUseCase
}

func NewPostReviewHandler(postReviewUC PostReviewUseCase) *postReviewHandler {
	return &postReviewHandler{postReviewUC: postReviewUC}
}

package reviews

import "context"

type ReviewService struct {
	reviewRepo *ReviewRepo
}

func NewReviewService(reviewRepo *ReviewRepo) *ReviewService {
	return &ReviewService{
		reviewRepo: reviewRepo,
	}
}

func (h *ReviewService) CreateReview(ctx context.Context, review Review) (Review, error) {
	tx, err := h.reviewRepo.db.BeginTx(ctx, nil)
	if err != nil {
		return Review{}, err
	}

	defer tx.Rollback()

	res, err := h.reviewRepo.CreateReview(ctx, review)
	if err != nil {
		return Review{}, err
	}

	createdReview := Review{
		ID:             res.ID,
		BidID:          *res.BidID,
		AuthorUsername: res.Username,
		Comment:        *res.Comment,
		CreatedAt:      *res.CreatedAt,
	}

	return createdReview, nil
}

func (h *ReviewService) GetReviews(ctx context.Context, bidId int32) ([]Review, error) {
	tx, err := h.reviewRepo.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	res, err := h.reviewRepo.GetReviews(ctx, bidId)
	if err != nil {
		return nil, err
	}

	reviews := make([]Review, len(res))

	for i, review := range res {
		reviews[i] = Review{
			ID:             review.ID,
			BidID:          *review.BidID,
			AuthorUsername: review.Username,
			Comment:        *review.Comment,
			CreatedAt:      *review.CreatedAt,
		}
	}

	return reviews, nil
}

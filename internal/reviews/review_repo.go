package reviews

import (
	"context"
	"database/sql"
	repoModel "tender-service/internal/repository/tables/tenders/public/model"
	"tender-service/internal/repository/tables/tenders/public/table"

	"github.com/go-jet/jet/v2/postgres"
)

type ReviewRepo struct {
	db *sql.DB
}

func NewReviewRepo(db *sql.DB) *ReviewRepo {
	return &ReviewRepo{db: db}
}

func (h *ReviewRepo) CreateReview(ctx context.Context, review Review) (repoModel.Review, error) {
	stmt := table.Review.INSERT(
		table.Review.BidID,
		table.Review.Username,
		table.Review.Comment,
	).VALUES(
		review.BidID,
		review.AuthorUsername,
		review.Comment,
	).RETURNING(table.Review.AllColumns)

	dest := repoModel.Review{}

	err := stmt.QueryContext(ctx, h.db, &dest)
	if err != nil {
		return repoModel.Review{}, err
	}

	return dest, nil
}

func (h *ReviewRepo) GetReviews(ctx context.Context, bidId int32) ([]repoModel.Review, error) {
	stmt := table.Review.SELECT(table.Review.AllColumns).WHERE(table.Review.BidID.EQ(postgres.Int32(bidId)))

	dest := []repoModel.Review{}

	err := stmt.QueryContext(ctx, h.db, &dest)
	if err != nil {
		return nil, err
	}

	return dest, nil
}

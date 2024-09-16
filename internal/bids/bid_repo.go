package bids

import (
	"context"
	"database/sql"
	repoModel "tender-service/internal/repository/tables/tenders/public/model"
	"tender-service/internal/repository/tables/tenders/public/table"
	"tender-service/internal/utils"
	"time"

	"github.com/go-jet/jet/v2/postgres"
)

type BidRepo struct {
	db *sql.DB
}

func NewBidRepo(db *sql.DB) *BidRepo {
	return &BidRepo{db: db}
}

func (h *BidRepo) GetBidsByTenderId(ctx context.Context, tenderId int32) ([]repoModel.Bid, error) {
	stmt := table.Bid.SELECT(table.Bid.AllColumns).
		WHERE(table.Bid.TenderID.EQ(postgres.Int32(tenderId)))

	dest := []repoModel.Bid{}

	err := stmt.QueryContext(ctx, h.db, &dest)
	if err != nil {
		return nil, err
	}

	return dest, nil
}

func (h *BidRepo) GetBidsByCreatorUsername(ctx context.Context, username string) ([]repoModel.Bid, error) {
	stmt := table.Bid.SELECT(table.Bid.AllColumns).
		WHERE(table.Bid.CreatorUsername.EQ(postgres.String(username)))

	dest := []repoModel.Bid{}

	err := stmt.QueryContext(ctx, h.db, &dest)
	if err != nil {
		return nil, err
	}

	return dest, nil
}

func (h *BidRepo) GetMyBids(ctx context.Context, organizationId string) ([]repoModel.Bid, error) {
	stmt := table.Bid.SELECT(table.Bid.AllColumns).
		WHERE(table.Bid.OrganizationID.EQ(postgres.UUID(utils.StringStringer(organizationId))))

	dest := []repoModel.Bid{}

	err := stmt.QueryContext(ctx, h.db, &dest)
	if err != nil {
		return nil, err
	}

	return dest, nil
}

func (h *BidRepo) GetBidById(ctx context.Context, bidId int32) (repoModel.Bid, error) {
	stmt := table.Bid.SELECT(table.Bid.AllColumns).WHERE(table.Bid.ID.EQ(postgres.Int32(bidId)))
	dest := repoModel.Bid{}

	err := stmt.QueryContext(ctx, h.db, &dest)
	if err != nil {
		return repoModel.Bid{}, err
	}

	return dest, nil
}

func (h *BidRepo) CreateBidVersion(ctx context.Context, bid repoModel.Bid) error {
	stmt := table.BidVersions.INSERT(
		table.BidVersions.BidID,
		table.BidVersions.Name,
		table.BidVersions.Description,
		table.BidVersions.Status,
		table.BidVersions.TenderID,
		table.BidVersions.OrganizationID,
		table.BidVersions.CreatorUsername,
		table.BidVersions.Version,
		table.BidVersions.CreatedAt,
		table.BidVersions.UpdatedAt,
	).VALUES(
		bid.ID,
		bid.Name,
		bid.Description,
		bid.Status,
		bid.TenderID,
		bid.OrganizationID,
		bid.CreatorUsername,
		bid.Version,
		bid.CreatedAt,
		bid.UpdatedAt,
	)

	_, err := stmt.ExecContext(ctx, h.db)
	if err != nil {
		return err
	}

	return nil
}

func (h *BidRepo) GetBidVersionById(ctx context.Context, bidId, version int32) (repoModel.BidVersions, error) {
	stmt := table.BidVersions.SELECT(
		table.BidVersions.AllColumns,
	).WHERE(
		table.BidVersions.BidID.EQ(postgres.Int32(bidId)).
			AND(table.BidVersions.Version.EQ(postgres.Int32(version))),
	)
	dest := repoModel.BidVersions{}

	err := stmt.QueryContext(ctx, h.db, &dest)
	if err != nil {
		return repoModel.BidVersions{}, err
	}

	return dest, nil
}

func (h *BidRepo) CreateBid(ctx context.Context, bid Bid) (repoModel.Bid, error) {
	stmt := table.Bid.INSERT(
		table.Bid.Name,
		table.Bid.Description,
		table.Bid.TenderID,
		table.Bid.OrganizationID,
		table.Bid.CreatorUsername,
	).VALUES(
		bid.Name,
		bid.Description,
		bid.TenderId,
		bid.OrganizationId,
		bid.CreatorUsername,
	).RETURNING(table.Bid.AllColumns)

	dest := repoModel.Bid{}

	err := stmt.QueryContext(ctx, h.db, &dest)
	if err != nil {
		return repoModel.Bid{}, err
	}

	return dest, nil
}

func (h *BidRepo) EditBid(ctx context.Context, bid Bid, oldBid repoModel.Bid) (repoModel.Bid, error) {
	if bid.Name == "" {
		bid.Name = oldBid.Name
	}

	if bid.Description == "" {
		bid.Description = *oldBid.Description
	}

	stmt := table.Bid.UPDATE(
		table.Bid.Name,
		table.Bid.Description,
		table.Bid.Version,
		table.Bid.UpdatedAt,
	).SET(
		bid.Name,
		bid.Description,
		*oldBid.Version+1,
		time.Now(),
	).WHERE(
		table.Bid.ID.EQ(postgres.Int32(bid.ID)),
	).RETURNING(table.Bid.AllColumns)

	dest := repoModel.Bid{}

	err := stmt.QueryContext(ctx, h.db, &dest)
	if err != nil {
		return repoModel.Bid{}, err
	}

	return dest, nil
}

func (h *BidRepo) SetBidStatus(ctx context.Context, bidId int32, status string) error {
	stmt := table.Bid.UPDATE(table.Bid.Status).
		SET(status).WHERE(table.Bid.ID.EQ(postgres.Int32(bidId)))

	_, err := stmt.ExecContext(ctx, h.db)
	if err != nil {
		return err
	}

	return nil
}

func (h *BidRepo) RollbackBid(ctx context.Context, oldBid repoModel.BidVersions, version int32) (repoModel.Bid, error) {
	stmt := table.Bid.UPDATE(
		table.Bid.Name,
		table.Bid.Description,
		table.Bid.Status,
		table.Bid.TenderID,
		table.Bid.OrganizationID,
		table.Bid.CreatorUsername,
		table.Bid.Version,
		table.Bid.UpdatedAt,
	).SET(
		oldBid.Name,
		oldBid.Description,
		oldBid.Status,
		oldBid.TenderID,
		oldBid.OrganizationID,
		oldBid.CreatorUsername,
		version+1,
		time.Now(),
	).WHERE(
		table.Bid.ID.EQ(postgres.Int32(*oldBid.BidID)),
	).RETURNING(table.Bid.AllColumns)

	dest := repoModel.Bid{}

	err := stmt.QueryContext(ctx, h.db, &dest)
	if err != nil {
		return repoModel.Bid{}, err
	}

	return dest, nil
}

func (h *BidRepo) VoteOnBid(ctx context.Context, vote BidVote) (repoModel.BidVotes, error) {
	stmt := table.BidVotes.INSERT(
		table.BidVotes.BidID,
		table.BidVotes.Username,
		table.BidVotes.Decision,
	).VALUES(
		vote.BidId,
		vote.CreatorUsername,
		vote.Decision,
	).RETURNING(table.BidVotes.AllColumns)

	dest := repoModel.BidVotes{}

	err := stmt.QueryContext(ctx, h.db, &dest)
	if err != nil {
		return repoModel.BidVotes{}, err
	}

	return dest, nil
}

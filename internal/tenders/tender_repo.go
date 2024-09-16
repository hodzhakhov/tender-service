package tenders

import (
	"context"
	"database/sql"
	repoModel "tender-service/internal/repository/tables/tenders/public/model"
	"tender-service/internal/repository/tables/tenders/public/table"
	"tender-service/internal/utils"
	"time"

	"github.com/go-jet/jet/v2/postgres"
)

type TenderRepo struct {
	db *sql.DB
}

func NewTenderRepo(db *sql.DB) *TenderRepo {
	return &TenderRepo{db: db}
}

func (h *TenderRepo) GetTenders(ctx context.Context, serviceType string) ([]repoModel.Tender, error) {
	stmt := table.Tender.SELECT(table.Tender.AllColumns)

	if serviceType != "" {
		stmt = stmt.WHERE(table.Tender.ServiceType.EQ(postgres.String(serviceType)))
	}

	dest := []repoModel.Tender{}

	err := stmt.QueryContext(ctx, h.db, &dest)
	if err != nil {
		return nil, err
	}

	return dest, nil
}

func (h *TenderRepo) GetMyTenders(ctx context.Context, organizationId string) ([]repoModel.Tender, error) {
	stmt := table.Tender.SELECT(table.Tender.AllColumns).
		WHERE(table.Tender.OrganizationID.EQ(postgres.UUID(utils.StringStringer(organizationId))))

	dest := []repoModel.Tender{}

	err := stmt.QueryContext(ctx, h.db, &dest)
	if err != nil {
		return nil, err
	}

	return dest, nil
}

func (h *TenderRepo) GetTenderById(ctx context.Context, tenderId int32) (repoModel.Tender, error) {
	stmt := table.Tender.SELECT(table.Tender.AllColumns).WHERE(table.Tender.ID.EQ(postgres.Int32(tenderId)))
	dest := repoModel.Tender{}

	err := stmt.QueryContext(ctx, h.db, &dest)
	if err != nil {
		return repoModel.Tender{}, err
	}

	return dest, nil
}

func (h *TenderRepo) CreateTenderVersion(ctx context.Context, tender repoModel.Tender) error {
	stmt := table.TenderVersions.INSERT(
		table.TenderVersions.TenderID,
		table.TenderVersions.Name,
		table.TenderVersions.Description,
		table.TenderVersions.ServiceType,
		table.TenderVersions.Status,
		table.TenderVersions.OrganizationID,
		table.TenderVersions.CreatorUsername,
		table.TenderVersions.Version,
		table.TenderVersions.CreatedAt,
		table.TenderVersions.UpdatedAt,
	).VALUES(
		tender.ID,
		tender.Name,
		tender.Description,
		tender.ServiceType,
		tender.Status,
		tender.OrganizationID,
		tender.CreatorUsername,
		tender.Version,
		tender.CreatedAt,
		tender.UpdatedAt,
	)

	_, err := stmt.ExecContext(ctx, h.db)
	if err != nil {
		return err
	}

	return nil
}

func (h *TenderRepo) GetTenderVersionById(ctx context.Context, tenderId, version int32) (repoModel.TenderVersions, error) {
	oldStmt := table.TenderVersions.SELECT(
		table.TenderVersions.AllColumns,
	).WHERE(
		table.TenderVersions.TenderID.EQ(postgres.Int32(tenderId)).
			AND(table.TenderVersions.Version.EQ(postgres.Int32(version))),
	)
	dest := repoModel.TenderVersions{}

	err := oldStmt.QueryContext(ctx, h.db, &dest)
	if err != nil {
		return repoModel.TenderVersions{}, err
	}

	return dest, nil
}

func (h *TenderRepo) CreateTender(ctx context.Context, tender Tender) (repoModel.Tender, error) {
	insStmt := table.Tender.INSERT(
		table.Tender.Name,
		table.Tender.Description,
		table.Tender.ServiceType,
		table.Tender.OrganizationID,
		table.Tender.CreatorUsername,
	).VALUES(
		tender.Name,
		tender.Description,
		tender.ServiceType,
		tender.OrganizationId,
		tender.CreatorUsername,
	).RETURNING(table.Tender.AllColumns)

	dest := repoModel.Tender{}

	err := insStmt.QueryContext(ctx, h.db, &dest)
	if err != nil {
		return repoModel.Tender{}, err
	}

	return dest, nil
}

func (h *TenderRepo) EditTender(ctx context.Context, tender Tender, oldTender repoModel.Tender) (repoModel.Tender, error) {
	if tender.Name == "" {
		tender.Name = oldTender.Name
	}

	if tender.Description == "" {
		tender.Description = *oldTender.Description
	}

	if tender.ServiceType == "" {
		tender.ServiceType = oldTender.ServiceType
	}

	updStmt := table.Tender.UPDATE(
		table.Tender.Name,
		table.Tender.Description,
		table.Tender.ServiceType,
		table.Tender.Version,
		table.Tender.UpdatedAt,
	).SET(
		tender.Name,
		tender.Description,
		tender.ServiceType,
		*oldTender.Version+1,
		time.Now(),
	).WHERE(
		table.Tender.ID.EQ(postgres.Int32(tender.ID)),
	).RETURNING(table.Tender.AllColumns)

	dest := repoModel.Tender{}

	err := updStmt.QueryContext(ctx, h.db, &dest)
	if err != nil {
		return repoModel.Tender{}, err
	}

	return dest, nil
}

func (h *TenderRepo) SetTenderStatus(ctx context.Context, tenderId int32, status string) error {
	stmt := table.Tender.UPDATE(table.Tender.Status).
		SET(status).WHERE(table.Tender.ID.EQ(postgres.Int32(tenderId)))

	_, err := stmt.ExecContext(ctx, h.db)
	if err != nil {
		return err
	}

	return nil
}

func (h *TenderRepo) RollbackTender(ctx context.Context, oldTender repoModel.TenderVersions, version int32) (repoModel.Tender, error) {
	stmt := table.Tender.UPDATE(
		table.Tender.Name,
		table.Tender.Description,
		table.Tender.ServiceType,
		table.Tender.OrganizationID,
		table.Tender.CreatorUsername,
		table.Tender.Version,
		table.Tender.UpdatedAt,
	).SET(
		oldTender.Name,
		oldTender.Description,
		oldTender.ServiceType,
		oldTender.OrganizationID,
		oldTender.CreatorUsername,
		version+1,
		time.Now(),
	).WHERE(
		table.Tender.ID.EQ(postgres.Int32(*oldTender.TenderID)),
	).RETURNING(table.Tender.AllColumns)

	dest := repoModel.Tender{}

	err := stmt.QueryContext(ctx, h.db, &dest)
	if err != nil {
		return repoModel.Tender{}, err
	}

	return dest, nil
}

func (h *TenderRepo) CheckDecisionCount(ctx context.Context, bidId int32, decision bool) ([]repoModel.BidVotes, error) {
	stmt := table.BidVotes.SELECT(table.BidVotes.AllColumns).
		WHERE(table.BidVotes.BidID.EQ(postgres.Int32(bidId)).AND(table.BidVotes.Decision.EQ(postgres.Bool(decision))))

	var dest []repoModel.BidVotes

	err := stmt.QueryContext(ctx, h.db, &dest)
	if err != nil {
		return nil, err
	}

	return dest, nil
}

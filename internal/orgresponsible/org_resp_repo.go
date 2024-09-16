package orgresponsible

import (
	"context"
	"database/sql"
	repoModel "tender-service/internal/repository/tables/tenders/public/model"
	"tender-service/internal/repository/tables/tenders/public/table"
	"tender-service/internal/utils"

	"github.com/go-jet/jet/v2/postgres"
)

type OrgRespRepo struct {
	db *sql.DB
}

func NewOrgRespRepo(db *sql.DB) *OrgRespRepo {
	return &OrgRespRepo{db: db}
}

func (h *OrgRespRepo) GetResponsible(ctx context.Context, userId string) (repoModel.OrganizationResponsible, error) {
	stmt := table.OrganizationResponsible.SELECT(table.OrganizationResponsible.AllColumns).
		WHERE(table.OrganizationResponsible.UserID.EQ(postgres.UUID(utils.StringStringer(userId))))

	dest := repoModel.OrganizationResponsible{}

	err := stmt.QueryContext(ctx, h.db, &dest)
	if err != nil {
		return repoModel.OrganizationResponsible{}, err
	}

	return dest, nil
}

func (h *OrgRespRepo) GetResponsiblesByOrgId(ctx context.Context, orgId string) ([]repoModel.OrganizationResponsible, error) {
	stmt := table.OrganizationResponsible.SELECT(table.OrganizationResponsible.AllColumns).
		WHERE(table.OrganizationResponsible.OrganizationID.EQ(postgres.UUID(utils.StringStringer(orgId))))

	dest := []repoModel.OrganizationResponsible{}

	err := stmt.QueryContext(ctx, h.db, &dest)
	if err != nil {
		return nil, err
	}

	return dest, nil
}

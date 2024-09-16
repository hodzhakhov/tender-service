package orgresponsible

import "context"

type OrgRespService struct {
	orgRespRepo *OrgRespRepo
}

func NewOrgRespService(orgRespRepo *OrgRespRepo) *OrgRespService {
	return &OrgRespService{
		orgRespRepo: orgRespRepo,
	}
}

func (h *OrgRespService) GetResponsible(ctx context.Context, userId string) (OrganizationResponsible, error) {
	tx, err := h.orgRespRepo.db.BeginTx(ctx, nil)
	if err != nil {
		return OrganizationResponsible{}, err
	}

	defer tx.Rollback()

	res, err := h.orgRespRepo.GetResponsible(ctx, userId)
	if err != nil {
		return OrganizationResponsible{}, err
	}

	responsible := OrganizationResponsible{
		ID:             res.ID.String(),
		OrganizationId: res.OrganizationID.String(),
		UserId:         res.UserID.String(),
	}

	return responsible, nil
}

func (h *OrgRespService) GetResponsiblesByOrgId(ctx context.Context, orgId string) ([]OrganizationResponsible, error) {
	tx, err := h.orgRespRepo.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	res, err := h.orgRespRepo.GetResponsiblesByOrgId(ctx, orgId)
	if err != nil {
		return nil, err
	}

	responsibles := make([]OrganizationResponsible, len(res))

	for i, resp := range res {
		responsibles[i] = OrganizationResponsible{
			ID:             resp.ID.String(),
			OrganizationId: resp.OrganizationID.String(),
			UserId:         resp.UserID.String(),
		}
	}

	return responsibles, nil
}

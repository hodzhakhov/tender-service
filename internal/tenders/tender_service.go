package tenders

import "context"

type TenderService struct {
	tenderRepo *TenderRepo
}

func NewTenderService(tenderRepo *TenderRepo) *TenderService {
	return &TenderService{
		tenderRepo: tenderRepo,
	}
}

func (h *TenderService) GetTenders(ctx context.Context, serviceType string) ([]Tender, error) {
	tx, err := h.tenderRepo.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	res, err := h.tenderRepo.GetTenders(ctx, serviceType)
	if err != nil {
		return nil, err
	}

	tenders := make([]Tender, 0)
	for _, tender := range res {
		if *tender.Status == "PUBLISHED" {
			tenders = append(tenders, Tender{
				ID:              tender.ID,
				Name:            tender.Name,
				Description:     *tender.Description,
				ServiceType:     tender.ServiceType,
				Status:          *tender.Status,
				OrganizationId:  tender.OrganizationID.String(),
				CreatorUsername: tender.CreatorUsername,
				Version:         *tender.Version,
				CreatedAt:       *tender.CreatedAt,
				UpdatedAt:       *tender.UpdatedAt,
			})
		}
	}

	return tenders, tx.Commit()
}

func (h *TenderService) GetMyTenders(ctx context.Context, organizationId string) ([]Tender, error) {
	tx, err := h.tenderRepo.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	res, err := h.tenderRepo.GetMyTenders(ctx, organizationId)
	if err != nil {
		return nil, err
	}

	tenders := make([]Tender, len(res))
	for i, tender := range res {
		tenders[i] = Tender{
			ID:              tender.ID,
			Name:            tender.Name,
			Description:     *tender.Description,
			ServiceType:     tender.ServiceType,
			Status:          *tender.Status,
			OrganizationId:  tender.OrganizationID.String(),
			CreatorUsername: tender.CreatorUsername,
			Version:         *tender.Version,
			CreatedAt:       *tender.CreatedAt,
			UpdatedAt:       *tender.UpdatedAt,
		}
	}

	return tenders, tx.Commit()
}

func (h *TenderService) CreateTender(ctx context.Context, tender Tender) (Tender, error) {
	tx, err := h.tenderRepo.db.BeginTx(ctx, nil)
	if err != nil {
		return Tender{}, err
	}

	defer tx.Rollback()

	res, err := h.tenderRepo.CreateTender(ctx, tender)
	if err != nil {
		return Tender{}, err
	}

	err = h.tenderRepo.CreateTenderVersion(ctx, res)
	if err != nil {
		return Tender{}, err
	}

	createdTender := Tender{
		ID:              res.ID,
		Name:            res.Name,
		Description:     *res.Description,
		ServiceType:     res.ServiceType,
		Status:          *res.Status,
		OrganizationId:  res.OrganizationID.String(),
		CreatorUsername: res.CreatorUsername,
		Version:         *res.Version,
		CreatedAt:       *res.CreatedAt,
		UpdatedAt:       *res.UpdatedAt,
	}

	return createdTender, tx.Commit()
}

func (h *TenderService) EditTender(ctx context.Context, tender Tender) (Tender, error) {
	tx, err := h.tenderRepo.db.BeginTx(ctx, nil)
	if err != nil {
		return Tender{}, err
	}

	defer tx.Rollback()

	oldTender, err := h.tenderRepo.GetTenderById(ctx, tender.ID)
	if err != nil {
		return Tender{}, err
	}

	res, err := h.tenderRepo.EditTender(ctx, tender, oldTender)
	if err != nil {
		return Tender{}, err
	}

	err = h.tenderRepo.CreateTenderVersion(ctx, res)
	if err != nil {
		return Tender{}, err
	}

	updTender := Tender{
		ID:              res.ID,
		Name:            res.Name,
		Description:     *res.Description,
		ServiceType:     res.ServiceType,
		Status:          *res.Status,
		OrganizationId:  res.OrganizationID.String(),
		CreatorUsername: res.CreatorUsername,
		Version:         *res.Version,
		CreatedAt:       *res.CreatedAt,
		UpdatedAt:       *res.UpdatedAt,
	}

	return updTender, tx.Commit()
}

func (h *TenderService) SetTenderStatus(ctx context.Context, tenderId int32, status string) error {
	tx, err := h.tenderRepo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	err = h.tenderRepo.SetTenderStatus(ctx, tenderId, status)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (h *TenderService) RollbackTender(ctx context.Context, tenderId, version int32) (Tender, error) {
	tx, err := h.tenderRepo.db.BeginTx(ctx, nil)
	if err != nil {
		return Tender{}, err
	}

	defer tx.Rollback()

	currentTender, err := h.tenderRepo.GetTenderById(ctx, tenderId)
	if err != nil {
		return Tender{}, err
	}

	oldTender, err := h.tenderRepo.GetTenderVersionById(ctx, tenderId, version)
	if err != nil {
		return Tender{}, err
	}

	res, err := h.tenderRepo.RollbackTender(ctx, oldTender, *currentTender.Version)
	if err != nil {
		return Tender{}, err
	}

	err = h.tenderRepo.CreateTenderVersion(ctx, res)
	if err != nil {
		return Tender{}, err
	}

	updTender := Tender{
		ID:              res.ID,
		Name:            res.Name,
		Description:     *res.Description,
		ServiceType:     res.ServiceType,
		Status:          *res.Status,
		OrganizationId:  res.OrganizationID.String(),
		CreatorUsername: res.CreatorUsername,
		Version:         *res.Version,
		CreatedAt:       *res.CreatedAt,
		UpdatedAt:       *res.UpdatedAt,
	}

	return updTender, tx.Commit()
}

func (h *TenderService) GetTenderById(ctx context.Context, tenderId int32) (Tender, error) {
	tx, err := h.tenderRepo.db.BeginTx(ctx, nil)
	if err != nil {
		return Tender{}, err
	}

	defer tx.Rollback()

	res, err := h.tenderRepo.GetTenderById(ctx, tenderId)
	if err != nil {
		return Tender{}, err
	}

	tender := Tender{
		ID:              res.ID,
		Name:            res.Name,
		Description:     *res.Description,
		ServiceType:     res.ServiceType,
		Status:          *res.Status,
		OrganizationId:  res.OrganizationID.String(),
		CreatorUsername: res.CreatorUsername,
		Version:         *res.Version,
		CreatedAt:       *res.CreatedAt,
		UpdatedAt:       *res.UpdatedAt,
	}

	return tender, tx.Commit()
}

func (h *TenderService) CheckDecisionCount(ctx context.Context, bidId int32, decision bool) (int32, error) {
	tx, err := h.tenderRepo.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	defer tx.Rollback()

	res, err := h.tenderRepo.CheckDecisionCount(ctx, bidId, decision)
	if err != nil {
		return 0, err
	}

	return int32(len(res)), tx.Commit()
}

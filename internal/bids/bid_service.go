package bids

import "context"

type BidService struct {
	bidRepo *BidRepo
}

func NewBidService(bidRepo *BidRepo) *BidService {
	return &BidService{
		bidRepo: bidRepo,
	}
}

func (h *BidService) GetBidsByTenderId(ctx context.Context, tenderId int32) ([]Bid, error) {
	tx, err := h.bidRepo.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	res, err := h.bidRepo.GetBidsByTenderId(ctx, tenderId)
	if err != nil {
		return nil, err
	}

	bids := make([]Bid, len(res))
	for i, bid := range res {
		bids[i] = Bid{
			ID:              bid.ID,
			Name:            bid.Name,
			Description:     *bid.Description,
			Status:          *bid.Status,
			TenderId:        *bid.TenderID,
			OrganizationId:  bid.OrganizationID.String(),
			CreatorUsername: bid.CreatorUsername,
			Version:         *bid.Version,
			CreatedAt:       *bid.CreatedAt,
			UpdatedAt:       *bid.UpdatedAt,
		}
	}

	return bids, tx.Commit()
}

func (h *BidService) GetBidById(ctx context.Context, bidId int32) (Bid, error) {
	tx, err := h.bidRepo.db.BeginTx(ctx, nil)
	if err != nil {
		return Bid{}, err
	}

	defer tx.Rollback()

	res, err := h.bidRepo.GetBidById(ctx, bidId)
	if err != nil {
		return Bid{}, err
	}

	bid := Bid{
		ID:              res.ID,
		Name:            res.Name,
		Description:     *res.Description,
		Status:          *res.Status,
		TenderId:        *res.TenderID,
		OrganizationId:  res.OrganizationID.String(),
		CreatorUsername: res.CreatorUsername,
		Version:         *res.Version,
		CreatedAt:       *res.CreatedAt,
		UpdatedAt:       *res.UpdatedAt,
	}

	return bid, tx.Commit()
}

func (h *BidService) GetBidsByCreatorUsername(ctx context.Context, username string) ([]Bid, error) {
	tx, err := h.bidRepo.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	res, err := h.bidRepo.GetBidsByCreatorUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	bids := make([]Bid, len(res))
	for i, bid := range res {
		bids[i] = Bid{
			ID:              bid.ID,
			Name:            bid.Name,
			Description:     *bid.Description,
			Status:          *bid.Status,
			TenderId:        *bid.TenderID,
			OrganizationId:  bid.OrganizationID.String(),
			CreatorUsername: bid.CreatorUsername,
			Version:         *bid.Version,
			CreatedAt:       *bid.CreatedAt,
			UpdatedAt:       *bid.UpdatedAt,
		}
	}

	return bids, tx.Commit()
}

func (h *BidService) GetMyBids(ctx context.Context, organizationId string) ([]Bid, error) {
	tx, err := h.bidRepo.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	res, err := h.bidRepo.GetMyBids(ctx, organizationId)
	if err != nil {
		return nil, err
	}

	bids := make([]Bid, len(res))
	for i, bid := range res {
		bids[i] = Bid{
			ID:              bid.ID,
			Name:            bid.Name,
			Description:     *bid.Description,
			Status:          *bid.Status,
			TenderId:        *bid.TenderID,
			OrganizationId:  bid.OrganizationID.String(),
			CreatorUsername: bid.CreatorUsername,
			Version:         *bid.Version,
			CreatedAt:       *bid.CreatedAt,
			UpdatedAt:       *bid.UpdatedAt,
		}
	}

	return bids, tx.Commit()
}

func (h *BidService) CreateBid(ctx context.Context, bid Bid) (Bid, error) {
	tx, err := h.bidRepo.db.BeginTx(ctx, nil)
	if err != nil {
		return Bid{}, err
	}

	defer tx.Rollback()

	res, err := h.bidRepo.CreateBid(ctx, bid)
	if err != nil {
		return Bid{}, err
	}

	err = h.bidRepo.CreateBidVersion(ctx, res)
	if err != nil {
		return Bid{}, err
	}

	createdBid := Bid{
		ID:              res.ID,
		Name:            res.Name,
		Description:     *res.Description,
		Status:          *res.Status,
		TenderId:        *res.TenderID,
		OrganizationId:  res.OrganizationID.String(),
		CreatorUsername: res.CreatorUsername,
		Version:         *res.Version,
		CreatedAt:       *res.CreatedAt,
		UpdatedAt:       *res.UpdatedAt,
	}

	return createdBid, tx.Commit()
}

func (h *BidService) EditBid(ctx context.Context, bid Bid) (Bid, error) {
	tx, err := h.bidRepo.db.BeginTx(ctx, nil)
	if err != nil {
		return Bid{}, err
	}

	defer tx.Rollback()

	oldBid, err := h.bidRepo.GetBidById(ctx, bid.ID)
	if err != nil {
		return Bid{}, err
	}

	res, err := h.bidRepo.EditBid(ctx, bid, oldBid)
	if err != nil {
		return Bid{}, err
	}

	err = h.bidRepo.CreateBidVersion(ctx, res)
	if err != nil {
		return Bid{}, err
	}

	updBid := Bid{
		ID:              res.ID,
		Name:            res.Name,
		Description:     *res.Description,
		Status:          *res.Status,
		TenderId:        *res.TenderID,
		OrganizationId:  res.OrganizationID.String(),
		CreatorUsername: res.CreatorUsername,
		Version:         *res.Version,
		CreatedAt:       *res.CreatedAt,
		UpdatedAt:       *res.UpdatedAt,
	}

	return updBid, tx.Commit()
}

func (h *BidService) SetBidStatus(ctx context.Context, bidId int32, status string) error {
	tx, err := h.bidRepo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	err = h.bidRepo.SetBidStatus(ctx, bidId, status)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (h *BidService) RollbackBid(ctx context.Context, bidId, version int32) (Bid, error) {
	tx, err := h.bidRepo.db.BeginTx(ctx, nil)
	if err != nil {
		return Bid{}, err
	}

	defer tx.Rollback()

	currentTender, err := h.bidRepo.GetBidById(ctx, bidId)
	if err != nil {
		return Bid{}, err
	}

	oldTender, err := h.bidRepo.GetBidVersionById(ctx, bidId, version)
	if err != nil {
		return Bid{}, err
	}

	res, err := h.bidRepo.RollbackBid(ctx, oldTender, *currentTender.Version)
	if err != nil {
		return Bid{}, err
	}

	err = h.bidRepo.CreateBidVersion(ctx, res)
	if err != nil {
		return Bid{}, err
	}

	updTender := Bid{
		ID:              res.ID,
		Name:            res.Name,
		Description:     *res.Description,
		Status:          *res.Status,
		TenderId:        *res.TenderID,
		OrganizationId:  res.OrganizationID.String(),
		CreatorUsername: res.CreatorUsername,
		Version:         *res.Version,
		CreatedAt:       *res.CreatedAt,
		UpdatedAt:       *res.UpdatedAt,
	}

	return updTender, tx.Commit()
}

func (h *BidService) VoteOnBid(ctx context.Context, vote BidVote) (BidVote, error) {
	tx, err := h.bidRepo.db.BeginTx(ctx, nil)
	if err != nil {
		return BidVote{}, err
	}

	defer tx.Rollback()

	bid, err := h.bidRepo.GetBidById(ctx, vote.BidId)
	if err != nil {
		return BidVote{}, err
	}

	if *bid.Status != "CREATED" {
		return BidVote{}, nil
	}

	res, err := h.bidRepo.VoteOnBid(ctx, vote)
	if err != nil {
		return BidVote{}, err
	}

	bidVote := BidVote{
		ID:        res.ID,
		BidId:     *res.BidID,
		Decision:  *res.Decision,
		CreatedAt: *res.CreatedAt,
	}

	return bidVote, tx.Commit()
}

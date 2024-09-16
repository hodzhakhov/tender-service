package bids

import "time"

type Bid struct {
	ID              int32     `json:"id"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	Status          string    `json:"status"`
	TenderId        int32     `json:"tenderId"`
	OrganizationId  string    `json:"organizationId"`
	CreatorUsername string    `json:"creatorUsername"`
	Version         int32     `json:"version"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

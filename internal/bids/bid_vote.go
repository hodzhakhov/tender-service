package bids

import "time"

type BidVote struct {
	ID              int32     `json:"id"`
	BidId           int32     `json:"bid_id"`
	CreatorUsername string    `json:"username"`
	Decision        bool      `json:"decision"`
	CreatedAt       time.Time `json:"created_at"`
}

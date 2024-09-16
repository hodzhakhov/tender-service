package reviews

import "time"

type Review struct {
	ID             int32     `json:"id"`
	BidID          int32     `json:"bid_id"`
	AuthorUsername string    `json:"author_username"`
	Comment        string    `json:"comment"`
	CreatedAt      time.Time `json:"created_at"`
}

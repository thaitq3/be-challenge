package dto

type Purchase struct {
	ID          uint    `json:"id"`
	WagerId     uint    `json:"wager_id"`
	BuyingPrice float64 `json:"buying_price"`
	BoughtAt    int64   `json:"bought_at"`
}

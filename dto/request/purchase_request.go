package request

type PurchaseRequest struct {
	WagerId     uint    `json:"wager_id"`
	BuyingPrice float64 `json:"buying_price" validate:"gt=0,is-monetary"`
}

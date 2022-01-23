package request

type CreateWagerRequest struct {
	TotalWagerValue   int64   `json:"total_wager_value" validate:"gt=0"`
	Odds              int64   `json:"odds" validate:"gt=0"`
	SellingPercentage int64   `json:"selling_percentage" validate:"gte=1,lte=100"`
	SellingPrice      float64 `json:"selling_price" validate:"gt=0,is-monetary"`
}

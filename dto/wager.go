package dto

type Wager struct {
	Id                  uint     `json:"id"`
	TotalWagerValue     int64    `json:"total_wager_value"`
	Odds                int64    `json:"odds"`
	SellingPercentage   int64    `json:"selling_percentage"`
	SellingPrice        float64  `json:"selling_[rice"`
	CurrentSellingPrice float64  `json:"current_selling_price"`
	PercentageSold      *int64   `json:"percentage_sold"`
	AmountSold          *float64 `json:"amount_sold"`
	PlacedAt            int64    `json:"placed_at"`
}

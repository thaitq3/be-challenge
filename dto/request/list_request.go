package request

type ListRequest struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

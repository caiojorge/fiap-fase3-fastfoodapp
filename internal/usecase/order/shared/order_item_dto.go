package usecase

type OrderItemDTO struct {
	ID        string  `json:"id"`
	ProductID string  `json:"productid"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
	Status    string  `json:"status"`
}

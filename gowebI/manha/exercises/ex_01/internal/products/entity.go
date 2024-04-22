package products

type Product struct {
	ID        uint64  `json:"id"`
	Name      string  `json:"name"`
	Color     string  `json:"color"`
	Price     float64 `json:"price"`
	Count     int     `json:"count"`
	Code      string  `json:"code"`
	Published bool    `json:"published"`
	CreatedAt string  `json:"created_at"`
}

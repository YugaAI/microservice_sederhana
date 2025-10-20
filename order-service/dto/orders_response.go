package dto

type OrderRequest struct {
	UserID int    `json:"user_id"`
	Item   string `json:"item"`
}
type OrdersResponse struct {
	OrderID int          `json:"order_id"`
	Item    string       `json:"item"`
	User    UserResponse `json:"user"`
}

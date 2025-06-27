package request

type CreateAddressRequest struct {
	Street     string `json:"street" binding:"required"`
	City       string `json:"city"`
	Province   string `json:"province" binding:"required"`
	Country    string `json:"country"`
	PostalCode string `json:"postal_code"`
}

type UpdateAddressRequest struct {
	Street     string `json:"street" binding:"required"`
	City       string `json:"city"`
	Province   string `json:"province" binding:"required"`
	Country    string `json:"country"`
	PostalCode string `json:"postal_code"`
}

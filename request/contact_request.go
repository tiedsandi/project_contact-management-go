package request

type CreateContactRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" binding:"required,email"`
	Phone     string `json:"phone"`
}

type UpdateContactRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" binding:"required,email"`
	Phone     string `json:"phone"`
}

package dtos

type CreateCustomerRequest struct {
	ID        int64  `json:"-"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
	Location  string `json:"location"`
}

type CreateCustomerResponse struct {
	ID int64 `json:"customer_id"`
}

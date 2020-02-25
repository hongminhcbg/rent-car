package dtos
// LocalCustomer customer have no account
type LocalCustomer struct {
	Phone     string `json:"customer_phone"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

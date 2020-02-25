package models

// Customer who want to rent car
type Customer struct {
	ID        int    `gorm:"column:id;primary_key"`
	FirstName string `gorm:"column:first_name;not null"`
	LastName  string `gorm:"column:last_name;not null"`
	Email     string `gorm:"column:email"`
	Phone     string `gorm:"column:phone;not null;unique"`
	Location  string `gorm:"column:location"`
}

// TableName set table for struct Customer
func (cli *Customer) TableName() string {
	return "customers"
}

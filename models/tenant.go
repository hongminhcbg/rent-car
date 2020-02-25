package models

// Tenant define, who have the cars and want customer rent them
type Tenant struct {
	ID              int    `gorm:"column:id;primary_key"`
	NameAgency      string `gorm:"column:name_agency;not null"`
	Email           string `gorm:"column:email;not null;unique"`
	Phone           string `gorm:"column:phone;not null;unique"`
	BankInformation string `gorm:"column:bank_information"`
	Location        string `gorm:"column:location;not null"`
}

// TableName set table for struct Tenant
func (c *Tenant) TableName() string {
	return "tenants"
}

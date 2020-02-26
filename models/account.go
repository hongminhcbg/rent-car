package models

// Account keep user and and password
type Account struct {
	ID       int    `gorm:"column:id;primary_key"`
	Phone    string `gorm:"column:phone;not null;unique"`
	Password string `gorm:"column:password;not null"`
	Types    string `gorm:"column:type;not null"`
}

func (Account) TableName() string {
	return "accounts"
}

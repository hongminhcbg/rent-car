package daos

import (
	"github.com/hongminhcbg/rent-car/models"
	"github.com/jinzhu/gorm"
)

type AccountDao interface {
	Login(acc models.Account) (*models.Account, error)
	Delete(acc models.Account) error
	Create(acc models.Account) error
}
func NewAccountDao(db *gorm.DB) AccountDao {
	return &accountDaoImpl{db:db,}
}
type accountDaoImpl struct {
	db *gorm.DB
}

func (c *accountDaoImpl) Login(acc models.Account) (*models.Account, error) {
	var accRes models.Account
	if err := c.db.Where("Phone = ? AND Passwords = ?", acc.Phone, acc.Password).First(&accRes).Error; err != nil {
		return nil, err
	}

	return &accRes, nil
}

func (c *accountDaoImpl) Delete(acc models.Account) error  {
	return c.db.Exec(`DELETE FROM accounts WHERE Phone = ?`, acc.Phone).Error
}

func (c *accountDaoImpl) Create(acc models.Account) error  {
	return c.db.Create(&acc).Error
}

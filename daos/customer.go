package daos

import (
	"github.com/hongminhcbg/rent-car/models"
	"github.com/jinzhu/gorm"
)

type CustomerDao interface {
	Create(Cus *models.Customer) (interface{}, error)
	Read(phone interface{}) (*models.Customer, error)
	ReadByID(cusID interface{}) (*models.Customer, error)
	Update(Cus *models.Customer) error
	Delete(Cus *models.Customer) error
	DeleteAcc(phone interface{}) error
	Login(acc *models.Account) (interface{}, error)
}

type customerDaoImpl struct {
	db *gorm.DB
}

func NewCustomerDao(db *gorm.DB) CustomerDao {
	return &customerDaoImpl{db}
}

func (c *customerDaoImpl) Create(Cus *models.Customer) (interface{}, error) {
	var Acc models.Account
	Acc.Phone = Cus.Phone
	Acc.Password = Cus.Password
	Acc.Types = model.CustomerTitle
	if err := c.db.Create(&Acc).Error; err != nil {
		return nil, err
	}

	if err := c.db.Create(Cus).Error; err != nil {
		return nil, err
	}

	return nil, nil
}

func (c *customerDaoImpl) Read(phone interface{}) (*models.Customer, error) {
	var customer models.Customer
	if err := c.db.Where("Phone = ?", phone).First(&customer).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}

func (c *customerDaoImpl) ReadByID(cusID interface{}) (*models.Customer, error) {
	var cus models.Customer
	err := c.db.Where("ID = ?", cusID).First(&cus).Error
	return &cus, err
}

func (c *customerDaoImpl) Update(Cus *models.Customer) error {
	return c.db.Save(Cus).Error
}

func (c *customerDaoImpl) Delete(Cus *models.Customer) error {
	c.db.Exec(`DELETE FROM customers WHERE ID = ? AND Phone = ?`, Cus.ID, Cus.Phone)
	return nil
}

func (c *customerDaoImpl) DeleteAcc(phone interface{}) error {
	return deleteAcc(c.db, phone)
}

func (c *customerDaoImpl) Login(acc *model.Account) (interface{}, error) {
	return login(c.db, acc)
}

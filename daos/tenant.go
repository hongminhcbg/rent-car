package daos

import (
	"github.com/hongminhcbg/rent-car/models"
	"github.com/jinzhu/gorm"
)

type TenantDao interface {
	Create(Ten *models.Tenant) (interface{}, error)
	Read(phone interface{}) (*models.Tenant, error)
	ReadByID(tenID interface{}) (*models.Tenant, error)
	Update(Ten *models.Tenant) (interface{}, error)
	Delete(Ten *models.Tenant) error
}

type tenantDaoImpl struct {
	db *gorm.DB
}

func NewTenantDao(db *gorm.DB) TenantDao {
	return &tenantDaoImpl{db}
}

func (c *tenantDaoImpl) Create(Ten *models.Tenant) (interface{}, error) {
	if err := c.db.Create(Ten).Error; err != nil {
		return nil, err
	}

	return nil, nil
}

func (c *tenantDaoImpl) Read(phone interface{}) (*models.Tenant, error) {
	var tenant models.Tenant
	if err := c.db.Where("Phone = ?", phone).First(&tenant).Error; err != nil {
		return nil, err
	}
	return &tenant, nil
}

func (c *tenantDaoImpl) ReadByID(tenID interface{}) (*models.Tenant, error) {
	var ten models.Tenant
	err := c.db.Where("ID = ?", tenID).First(&ten).Error
	return &ten, err
}

func (c *tenantDaoImpl) Update(Ten *models.Tenant) (interface{}, error) {
	return nil, nil
}

func (c *tenantDaoImpl) Delete(Ten *models.Tenant) error {
	return c.db.Exec(`DELETE FROM tenants WHERE ID = ? AND Phone = ?`, Ten.ID, Ten.Phone).Error
}


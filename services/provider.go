package services

import (
	"github.com/hongminhcbg/rent-car/daos"
	"github.com/jinzhu/gorm"
)

type ProviderService interface {
	GetUserService() UserService
}

type providerServiceImpl struct {
	db *gorm.DB
}

func NewProviderService(db *gorm.DB) ProviderService {
	return &providerServiceImpl{
		db:db,
	}
}
func (provider *providerServiceImpl) GetUserService() UserService {
	accountDao := daos.NewAccountDao(provider.db)
	customerDao := daos.NewCustomerDao(provider.db)
	tenantDao := daos.NewTenantDao(provider.db)

	return NewUserService(accountDao, customerDao, tenantDao)
}
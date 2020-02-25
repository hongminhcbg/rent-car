package services

import (
	"github.com/hongminhcbg/rent-car/common"
	"github.com/hongminhcbg/rent-car/daos"
	"github.com/hongminhcbg/rent-car/dtos"
	"github.com/hongminhcbg/rent-car/models"
)

type UserService interface {
	CreateCustomer(request dtos.CreateCustomerRequest) (*dtos.CreateCustomerResponse, error)

}

func NewUserService(accountDao daos.AccountDao,
	customerDao daos.CustomerDao,
	tenantDao daos.TenantDao,
) UserService {
	return &userServiceImpl{
		accountDao:  accountDao,
		customerDao: customerDao,
		tenantDao:   tenantDao,
	}
}

type userServiceImpl struct {
	accountDao  daos.AccountDao
	customerDao daos.CustomerDao
	tenantDao   daos.TenantDao
}

func (userService *userServiceImpl)CreateCustomer(request dtos.CreateCustomerRequest) (*dtos.CreateCustomerResponse, error){
	acc := models.Account{
		Phone:    request.Phone,
		Password: request.Password,
		Types:    common.CustomerAccoutType,
	}
	if err := userService.accountDao.Create(acc); err != nil {
		return nil, err
	}

	customer := models.Customer{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Phone:     request.Phone,
		Location:  request.Location,
	}
	if err := userService.customerDao.Create(&customer); err != nil {
		return nil, err
	}

	return &dtos.CreateCustomerResponse{ID:customer.ID}, nil
}
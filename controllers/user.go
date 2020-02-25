package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hongminhcbg/rent-car/dtos"
	"github.com/hongminhcbg/rent-car/services"
)

type UserController struct {
	UserService services.UserService
}

func (userController *UserController)CreateCustomer(context *gin.Context)  {
	var request dtos.CreateCustomerRequest
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(400, gin.H{
			"message":"error",
		})
	}

	if resp, err := userController.UserService.CreateCustomer(request); err != nil{
		context.JSON(400, gin.H{
			"message":"error",
		})
	} else {
		context.JSON(200, gin.H{
			"message":"create success",
			"data": resp,
		})
	}

}

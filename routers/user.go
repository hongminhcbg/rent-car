package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hongminhcbg/rent-car/conf"
	"github.com/hongminhcbg/rent-car/controllers"
	"github.com/hongminhcbg/rent-car/services"
	"github.com/jinzhu/gorm"
)

type Router struct {
	config *conf.MySQL
}

func NewRouter(config *conf.MySQL) Router {
	return Router{
		config:config,
	}
}

func (router *Router)InitGin() (*gin.Engine, error) {
	db, err := gorm.Open("mysql", router.config.URL)
	if err != nil {
		return nil, err
	}

	if err := db.DB().Ping(); err != nil {
		return nil, err
	} else {
		fmt.Println("ping db success")
	}

	providerService := services.NewProviderService(db)
	userController := controllers.UserController{UserService:providerService.GetUserService(),}

	engine := gin.Default()
	engine.POST("/api/v1/cutomer", userController.CreateCustomer)

	return engine, nil
}


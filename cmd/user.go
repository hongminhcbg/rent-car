package main

import (
	"github.com/hongminhcbg/rent-car/conf"
	"github.com/hongminhcbg/rent-car/routers"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	config := &conf.MySQL{URL: `root:1@tcp(localhost:3306)/rentcar?parseTime=true`}
	router := routers.NewRouter(config)
	engine, err := router.InitGin()
	if err != nil {
		panic(err)
	}
	engine.Run(":8080")

}

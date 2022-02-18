package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2"
	"micro-service/user/domain/repository"
	userService "micro-service/user/domain/service"
	"micro-service/user/handler"
	"micro-service/user/proto/user"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name("go.micro.service.user"),
		micro.Version("latest"),
	)

	//初始化服务
	srv.Init()

	//创建数据库连接
	db, err := gorm.Open("mysql", "root:123456@/micro?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	db.SingularTable(true)

	//只执行一次，数据表初始化
	//rp:=repository.NewUserRepository(db)
	//rp.InitTable()

	//创建服务实例
	userDataService := userService.NewUserDataService(repository.NewUserRepository(db))
	//注册Handler
	err = user.RegisterUserHandler(srv.Server(), &handler.User{UserDataService: userDataService})
	if err != nil {
		fmt.Println(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}

// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package router

import (
	"GFBackend/controller"
	"GFBackend/model/dao"
	"GFBackend/service"
)

// Injectors from wire.go:

func InitializeUserManageController() (*controller.UserManageController, error) {
	userDAO := dao.NewUserDAO()
	followDAO := dao.NewFollowDAO()
	userManageService := service.NewUserManageService(userDAO, followDAO)
	userManageController := controller.NewUserManageController(userManageService)
	return userManageController, nil
}

func InitializeCommunityManageController() (*controller.CommunityManageController, error) {
	communityDAO := dao.NewCommunityDAO()
	communityManageService := service.NewCommunityManageService(communityDAO)
	communityManageController := controller.NewCommunityManageController(communityManageService)
	return communityManageController, nil
}

func InitializeFileManageController() (*controller.FileManageController, error) {
	spaceDAO := dao.NewSpaceDAO()
	fileManageService := service.NewFileManageService(spaceDAO)
	fileManageController := controller.NewFileManageController(fileManageService)
	return fileManageController, nil
}
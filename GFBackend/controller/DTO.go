package controller

import (
	"GFBackend/model"
	"GFBackend/utils"
)

type ResponseMsg struct {
	Code    int    `json:"code" example:"200"`
	Message string `json:"message" example:"process successfully"`
}

type CommunityResponseMsg struct {
	Code    int             `json:"code" example:"200"`
	Message string          `json:"message" example:"process successfully"`
	Data    model.Community `json:"data"`
}

type UserInfo struct {
	Username    string `json:"Username" example:"jamesbond21"`
	Password    string `json:"Password" example:"f9ae5f68ae1e7f7f3fc06053e9b9b539"`
	NewPassword string `json:"NewPassword" example:"3ecb441b741bcd433288f5557e4b9118"`
	ForAdmin    bool   `json:"ForAdmin" example:true`
}

type NewUserInfo struct {
	Username   string `json:"Username" example:"jamesbond21"`
	Nickname   string `json:"Nickname" example:"Peter Park"`
	Birthday   string `json:"Birthday" example:"2022-02-30"`
	Gender     string `json:"Gender" example:"male/female/unknown"`
	Department string `json:"Department" example:"CS:GO"`
}

type CommunityInfo struct {
	Creator     string           `json:"Creator" example:"test1"`
	Name        string           `json:"Name" example:"community1"`
	Description string           `json:"Description" example:"this is a test community"`
	Create_Time *utils.LocalTime `json:"Create_Time" example:"2020-01-01"`
}

type UserFiles struct {
	ResponseMsg
	Filenames []string `json:"Filenames" example:"\"xxx.jpg\",\"xxx.png\",\"xxx.gif\""`
}
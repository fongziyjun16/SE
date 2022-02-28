package controller

import (
	"GFBackend/config"
	"GFBackend/middleware/auth"
	"GFBackend/model"
	"GFBackend/service"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"net/http"
	"strings"
	"sync"
)

var userManageControllerLock sync.Mutex
var userManageController *UserManageController

type UserManageController struct {
	userManageService service.IUserManageService
}

func NewUserManageController(userManageService service.IUserManageService) *UserManageController {
	if userManageController == nil {
		userManageControllerLock.Lock()
		if userManageController == nil {
			userManageController = &UserManageController{
				userManageService: userManageService,
			}
		}
		userManageControllerLock.Unlock()
	}
	return userManageController
}

var UserManageControllerSet = wire.NewSet(
	service.UserManageServiceSet,
	wire.Bind(new(service.IUserManageService), new(*service.UserManageService)),
	NewUserManageController,
)

// RegularRegister godoc
// @Summary Register a new Regular User
// @Description only need strings username & password
// @Tags User Manage
// @Accept json
// @Produce json
// @Param UserInfo body controller.UserInfo true "Regular User Register only needs Username, Password(encoded by md5) & ForAdmin with false."
// @Success 201 {object} controller.ResponseMsg "<b>Success</b>. User Register Successfully"
// @Failure 400 {object} controller.ResponseMsg "<b>Failure</b>. Bad Parameters or User Has Existed"
// @Failure 500 {object} controller.ResponseMsg "<b>Failure</b>. Server Internal Error."
// @Router /user/register [post]
func (userManageController *UserManageController) RegularRegister(context *gin.Context) {
	var registerInfo UserInfo
	if err := context.ShouldBindJSON(&registerInfo); err != nil {
		er := ResponseMsg{
			Code:    http.StatusBadRequest,
			Message: "Bad Parameters",
		}
		context.JSON(http.StatusBadRequest, er)
		return
	}

	err := userManageController.userManageService.Register(registerInfo.Username, registerInfo.Password, registerInfo.ForAdmin)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate") {
			er := ResponseMsg{
				Code:    http.StatusBadRequest,
				Message: "User Has Existed.",
			}
			context.JSON(http.StatusBadRequest, er)
		} else {
			er := ResponseMsg{
				Code:    http.StatusInternalServerError,
				Message: "Server Internal Error.",
			}
			context.JSON(http.StatusInternalServerError, er)
		}
		return
	}

	context.JSON(http.StatusCreated, ResponseMsg{
		Code:    http.StatusCreated,
		Message: "Create User Successfully",
	})
}

// AdminRegister godoc
// @Summary Register a new Admin User
// @Description only need strings username & password & ForAdmin, need token in cookie
// @Tags User Manage
// @Accept json
// @Produce json
// @Security ApiAuthToken
// @Param UserInfo body controller.UserInfo true "Admin User Register only needs Username, Password(encoded by md5) & ForAdmin with true."
// @Success 201 {object} controller.ResponseMsg "<b>Success</b>. User Register Successfully"
// @Failure 400 {object} controller.ResponseMsg "<b>Failure</b>. Bad Parameters or User Has Existed"
// @Failure 500 {object} controller.ResponseMsg "<b>Failure</b>. Server Internal Error."
// @Router /user/admin/register [post]
func (userManageController *UserManageController) AdminRegister(context *gin.Context) {
	userManageController.RegularRegister(context)
}

// UserLogin godoc
// @Summary Admin / Regular User login
// @Description only need strings username & password
// @Tags User Manage
// @Accept json
// @Produce json
// @Param UserInfo body controller.UserInfo true "only needs username and password"
// @Success 200 {object} controller.ResponseMsg "<b>Success</b>. User Login Successfully"
// @Failure 400 {object} controller.ResponseMsg "<b>Failure</b>. Bad Parameters or Username / Password incorrect"
// @Failure 500 {object} controller.ResponseMsg "<b>Failure</b>. Server Internal Error."
// @Router /user/login [post]
func (userManageController *UserManageController) UserLogin(context *gin.Context) {
	var userInfo UserInfo
	if err := context.ShouldBindJSON(&userInfo); err != nil {
		er := ResponseMsg{
			Code:    http.StatusBadRequest,
			Message: "Bad Parameters.",
		}
		context.JSON(http.StatusBadRequest, er)
		return
	}

	if token, err := userManageController.userManageService.Login(userInfo.Username, userInfo.Password); err != nil {
		if strings.Contains(err.Error(), "400") {
			er := ResponseMsg{
				Code:    http.StatusBadRequest,
				Message: "Username or Password is not correct",
			}
			context.JSON(http.StatusBadRequest, er)
		} else {
			er := ResponseMsg{
				Code:    http.StatusInternalServerError,
				Message: "Server Internal Error.",
			}
			context.JSON(http.StatusInternalServerError, er)
		}
		return
	} else {
		success := ResponseMsg{
			Code:    http.StatusOK,
			Message: token,
		}
		context.SetCookie("token", token, config.AppConfig.JWT.Expires*60, config.AppConfig.Server.BasePath, "localhost", false, true)
		context.JSON(http.StatusOK, success)
		return
	}
}

// UserLogout godoc
// @Summary Admin / Regular User logout
// @Description need strings username in post request, need token in cookie
// @Tags User Manage
// @Accept json
// @Produce json
// @Security ApiAuthToken
// @Param username body string true "username in post request body"
// @Router /user/logout [post]
func (userManageController *UserManageController) UserLogout(context *gin.Context) {
	type Info struct {
		Username string `json:"username"`
	}
	var info Info
	err := context.ShouldBind(&info)
	if err != nil {
		return
	}

	err = userManageController.userManageService.Logout(info.Username)
	if err != nil {
		return
	}
}

// UserUpdatePassword godoc
// @Summary Admin & Regular Update Password
// @Description need token in cookie, need Username, Password, NewPassword
// @Tags User Manage
// @Accept json
// @Produce json
// @Security ApiAuthToken
// @Param UserInfo body controller.UserInfo true "need Username, Password, NewPassword"
// @Success 201 {object} controller.ResponseMsg "<b>Success</b>. Update Password Successfully"
// @Failure 400 {object} controller.ResponseMsg "<b>Failure</b>. Bad Parameters or Password not match"
// @Failure 500 {object} controller.ResponseMsg "<b>Failure</b>. Server Internal Error."
// @Router /user/password [post]
func (userManageController *UserManageController) UserUpdatePassword(context *gin.Context) {
	var userInfo UserInfo
	if err1 := context.ShouldBindJSON(&userInfo); err1 != nil {
		er := ResponseMsg{
			Code:    http.StatusBadRequest,
			Message: "Bad Parameters.",
		}
		context.JSON(http.StatusBadRequest, er)
		return
	}

	err2 := userManageController.userManageService.UpdatePassword(userInfo.Username, userInfo.Password, userInfo.NewPassword)
	if err2 != nil {
		errMsg := ResponseMsg{
			Code:    http.StatusBadRequest,
			Message: "Bad Parameters.",
		}
		if strings.Contains(err2.Error(), "400") {
			errMsg.Message = "User old password not match"
		} else {
			errMsg.Code = http.StatusInternalServerError
			errMsg.Message = "Internal Server Error"
		}
		context.JSON(http.StatusBadRequest, errMsg)
		return
	}

	success := ResponseMsg{
		Code:    http.StatusOK,
		Message: "Update User Password Successfully",
	}
	context.JSON(http.StatusOK, success)
	return

}

// UserDelete godoc
// @Summary Admin delete Users, cannot self delete
// @Description need strings username in post request, need token in cookie
// @Tags User Manage
// @Accept json
// @Produce json
// @Security ApiAuthToken
// @Param username body string true "username in post request body"
// @Success 201 {object} controller.ResponseMsg "<b>Success</b>. Update Password Successfully"
// @Failure 400 {object} controller.ResponseMsg "<b>Failure</b>. Bad Parameters"
// @Failure 500 {object} controller.ResponseMsg "<b>Failure</b>. Server Internal Error."
// @Router /user/admin/delete [post]
func (userManageController *UserManageController) UserDelete(context *gin.Context) {
	type Info struct {
		Username string `json:"username"`
	}
	var info Info
	err1 := context.ShouldBind(&info)
	token, _ := context.Cookie("token")
	currentUsername, _ := auth.GetTokenUsername(token)
	if err1 != nil || info.Username == currentUsername {
		er := ResponseMsg{
			Code:    http.StatusBadRequest,
			Message: "Bad Parameters or Current User cannot delete self.",
		}
		context.JSON(http.StatusBadRequest, er)
		return
	}

	err2 := userManageController.userManageService.Delete(info.Username)
	if err2 != nil {
		er := ResponseMsg{
			Code:    http.StatusBadRequest,
			Message: "User not exist.",
		}
		if strings.Contains(err2.Error(), "user Policy") {
			er.Code = http.StatusInternalServerError
			er.Message = "Internal Server Error"
		}
		context.JSON(er.Code, er)
		return
	}

	context.JSON(http.StatusCreated, ResponseMsg{
		Code:    http.StatusCreated,
		Message: "Delete User Successfully",
	})
}

// UserUpdate godoc
// @Summary Update user information including Nickname, Birthday(yyyy-mm-dd), Gender(male / female / unknown), Department
// @Description need token in cookie, need Nickname, Birthday(yyyy-mm-dd), Gender(male / female / unknown), Department
// @Tags User Manage
// @Accept json
// @Produce json
// @Security ApiAuthToken
// @Param NewUserInfo body controller.NewUserInfo true "need Nickname, Birthday(yyyy-mm-dd), Gender(male / female / unknown), Department"
// @Success 201 {object} controller.ResponseMsg "<b>Success</b>. Update Password Successfully"
// @Failure 400 {object} controller.ResponseMsg "<b>Failure</b>. Bad Parameters"
// @Failure 500 {object} controller.ResponseMsg "<b>Failure</b>. Server Internal Error."
// @Router /user/update [post]
func (userManageController *UserManageController) UserUpdate(context *gin.Context) {
	var newUserInfo NewUserInfo
	if err1 := context.ShouldBindJSON(&newUserInfo); err1 != nil {
		er := ResponseMsg{
			Code:    http.StatusBadRequest,
			Message: "Bad Parameters.",
		}
		context.JSON(http.StatusBadRequest, er)
		return
	}

	userInfo := model.User{
		Username:   newUserInfo.Username,
		Nickname:   newUserInfo.Nickname,
		Birthday:   newUserInfo.Birthday,
		Gender:     newUserInfo.Gender,
		Department: newUserInfo.Department,
	}

	err2 := userManageController.userManageService.Update(userInfo)
	if err2 != nil {
		er := ResponseMsg{
			Code:    http.StatusBadRequest,
			Message: "Bad Parameters.",
		}
		context.JSON(http.StatusBadRequest, er)
		return

	}

	success := ResponseMsg{
		Code:    http.StatusOK,
		Message: "Update User Information Successfully",
	}
	context.JSON(http.StatusOK, success)
	return
}

// UserFollow godoc
// @Summary User Follow other users
// @Description need token in cookie, need username who is followed
// @Tags User Manage
// @Accept json
// @Produce json
// @Security ApiAuthToken
// @Param username body string true "username in post request body"
// @Success 201 {object} controller.ResponseMsg "<b>Success</b>. Follow Successfully"
// @Failure 400 {object} controller.ResponseMsg "<b>Failure</b>. Bad Parameters or User not exist."
// @Failure 500 {object} controller.ResponseMsg "<b>Failure</b>. Server Internal Error."
// @Router /user/follow [post]
func (userManageController *UserManageController) UserFollow(context *gin.Context) {
	type Info struct {
		Username string `json:"username"`
	}
	var info Info
	err1 := context.ShouldBind(&info)
	if err1 != nil {
		errMsg := ResponseMsg{
			Code:    http.StatusBadRequest,
			Message: "Bad Parameters.",
		}
		context.JSON(http.StatusBadRequest, errMsg)
		return
	}

	token, _ := context.Cookie("token")
	follower, _ := auth.GetTokenUsername(token)
	err2 := userManageController.userManageService.Follow(info.Username, follower)
	if err2 != nil {
		errMsg := ResponseMsg{
			Code:    http.StatusBadRequest,
			Message: "User not exist.",
		}
		if strings.Contains(err2.Error(), "500") {
			errMsg.Code = http.StatusInternalServerError
			errMsg.Message = "Internal Server Error"
		}
		context.JSON(http.StatusBadRequest, errMsg)
		return
	}

	success := ResponseMsg{
		Code:    http.StatusOK,
		Message: "Follow Successfully",
	}
	context.JSON(http.StatusOK, success)
	return
}

// UserUnfollow godoc
// @Summary User Unfollow other users
// @Description need token in cookie, need username who is followed
// @Tags User Manage
// @Accept json
// @Produce json
// @Security ApiAuthToken
// @Param username body string true "username in post request body"
// @Success 201 {object} controller.ResponseMsg "<b>Success</b>. Unfollow Successfully"
// @Failure 400 {object} controller.ResponseMsg "<b>Failure</b>. Bad Parameters or User not exist."
// @Failure 500 {object} controller.ResponseMsg "<b>Failure</b>. Server Internal Error."
// @Router /user/unfollow [post]
func (userManageController *UserManageController) UserUnfollow(context *gin.Context) {

	return
}
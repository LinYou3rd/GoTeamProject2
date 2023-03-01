package user

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Login(context *gin.Context) {

	var user userModel
	var toolUser userDataModel
	var admin adminModel
	var toolAdmin adminDataModel

	context.ShouldBindJSON(&toolUser)
	println(toolUser.Mode)
	println(toolUser.Account)
	println(toolUser.Password)

	if toolUser.Mode == "一般用户" || toolUser.Mode == "注册" {
		dbUser.Where("account=?", toolUser.Account).First(&user)
		println(user.ID)
		if user.ID == 0 {
			data := userLoginModel{
				Code:    404,
				Message: "账号不存在",
			}
			context.JSON(http.StatusNotFound, data)
			return
		}

		err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(toolUser.Password))
		if err != nil {
			data := userLoginModel{
				Code:    412,
				Message: "密码错误",
			}
			context.JSON(http.StatusPreconditionFailed, data)
			return
		}

		token, err := GenerateTokenForUser(user)
		if err != nil {
			data := userLoginModel{
				Code:    500,
				Message: "token签发失败",
			}
			context.JSON(http.StatusInternalServerError, data)
			return
		}

		data := userLoginModel{
			Model: gorm.Model{
				ID: user.ID,
			},
			Code:    200,
			Token:   token,
			Message: "登录成功",
			Name:    user.Name,
		}
		context.JSON(http.StatusOK, data)

	} else if toolUser.Mode == "管理员" {
		context.ShouldBindJSON(&toolAdmin)
		dbUser.Where("account=?", toolAdmin.Account).First(&admin)
		if admin.ID == 0 {
			data := adminLoginModel{
				Code:    404,
				Message: "账号不存在",
			}
			context.JSON(http.StatusNotFound, data)
			return
		}

		err := bcrypt.CompareHashAndPassword([]byte(admin.PasswordDigest), []byte(toolAdmin.Password))
		if err != nil {
			data := adminLoginModel{
				Code:    412,
				Message: "密码错误",
			}
			context.JSON(http.StatusPreconditionFailed, data)
			return
		}

		token, err := GenerateTokenForAdmin(admin)
		if err != nil {
			data := adminLoginModel{
				Code:    500,
				Message: "token签发失败",
			}
			context.JSON(http.StatusInternalServerError, data)
			return
		}

		data := adminLoginModel{
			Model: gorm.Model{
				ID: admin.ID,
			},
			Code:    200,
			Token:   token,
			Message: "登录成功",
		}
		context.JSON(http.StatusOK, data)

	}
}

func Enroll(context *gin.Context) {
	var toolUser userDataModel
	var user userModel
	context.ShouldBindJSON(&toolUser)
	println(toolUser.Account)
	dbUser.Where("account=?", toolUser.Account).First(&user)
	println(user.ID)
	if user.ID != 0 {
		data := userLoginModel{
			Code:    412,
			Message: "该账号已注册",
		}
		context.JSON(http.StatusPreconditionFailed, data)
		return
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(toolUser.Password), 1)
	if err != nil {
		data := userLoginModel{
			Code:    500,
			Message: "密码存储失败",
		}
		context.JSON(http.StatusInternalServerError, data)
		return
	}

	user.Account = toolUser.Account
	user.Contact = toolUser.Contact
	user.Name = toolUser.Name
	user.Identity = toolUser.Identity
	user.PasswordDigest = string(bytes)
	dbUser.Save(&user)

	data := userLoginModel{
		Code:    200,
		Message: "注册成功",
	}
	context.JSON(http.StatusOK, data)
	return

}

func ChangeInformation(context *gin.Context) {

	var user userModel
	var toolUser userDataModel
	context.ShouldBindJSON(&toolUser)
	dbUser.Where("id=?", toolUser.ID).First(&user)

	if newInformation := toolUser.Name; newInformation != "" {
		dbUser.Model(&user).Update("name", newInformation)
	}

	if newInformation := toolUser.Contact; newInformation != "" {
		dbUser.Model(&user).Update("contact", newInformation)
	}

	if newInformation := toolUser.Identity; newInformation != "" {
		dbUser.Model(&user).Update("identity", newInformation)
	}

	context.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "更新成功",
	})

}

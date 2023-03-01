package main

import (
	raise_project "ChineseHelpChinese/raise.project"
	"ChineseHelpChinese/user"
)
import "github.com/gin-gonic/gin"

func main() {
	user.Init()
	raise_project.Init()
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.POST("user/login", user.Login)
		v1.POST("user/enroll", user.Enroll)
		//退出登录...采用前端传入错误token的方式触发检测
		v1.PUT("user", user.ChangeInformation)
		v1.GET("admin", raise_project.AdminWork)
		project := v1.Group("/project")
		project.Use(user.JWT())
		{
			project.POST("", raise_project.AddProject)
			project.GET("", raise_project.GetTarget)
			project.GET("/all", raise_project.GetAll)
			project.PUT("", raise_project.Update)
			project.DELETE("", raise_project.DeleteProject)
			//有什么要实现的？ 获取单个项目信息 获取全部项目信息 更新项目（审核状态，资金）
			//检索单个项目（根据标题） 删除已筹集成功的项目（我觉得可以放在更新项目的部分）
			//添加单个项目
		}
	}

	err := router.Run(":8000")
	if err != nil {
		panic(err)
	}

}

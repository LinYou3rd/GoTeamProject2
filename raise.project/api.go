package raise_project

import (
	"ChineseHelpChinese/user"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AddProject(context *gin.Context) {

	var toolProject projectDataModel
	context.ShouldBindJSON(&toolProject)
	//怎么避免多个重复项目？以量取胜筹集资金
	data := forSave(toolProject)
	err := db.HMSet(toolProject.Title, data)

	if err != nil {
		data := projectService{
			Code:    500,
			Message: "存储失败",
		}
		context.JSON(http.StatusInternalServerError, data)
		return
	}

	information := projectService{
		Code:    200,
		Message: "创建成功",
	}

	context.JSON(http.StatusOK, information)
}

func GetTarget(context *gin.Context) {
	//用于展示界面获取项目具体信息、搜索项目返回数据
	var toolProject projectDataModel
	context.ShouldBindJSON(&toolProject)
	var title = toolProject.Title
	project, err := db.HGetAll(title).Result()

	if err != nil {
		data := projectService{
			Code:    500,
			Message: "获取项目信息失败",
		}
		context.JSON(http.StatusInternalServerError, data)
		return
	}

	information := forUse(project)
	information.Title = title
	information.Code = 200
	information.Message = "获取成功"
	context.JSON(http.StatusOK, information)
}

func GetAll(context *gin.Context) {
	//页面展示
	var i = 0
	var informations []projectService
	keys, err := db.Keys("*").Result()

	if err != nil {
		data := projectService{
			Code:    500,
			Message: "获取项目信息失败",
		}
		context.JSON(http.StatusInternalServerError, data)
		return
	}

	for _, key := range keys {
		var projects map[string]string
		projects, _ = db.HGetAll(key).Result()
		if projects["pass"] == "1" {
			informations[i] = forUse(projects)
			informations[i].Title = key
			i++
		}
	}

	context.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    informations,
	})
}

func Update(context *gin.Context) {
	//审核按钮，捐款按钮
	var adminWork user.AdminWorkModel
	context.ShouldBindJSON(&adminWork)
	mode := adminWork.Mode
	title := adminWork.Title

	if mode == "管理员" {
		pass := adminWork.Pass

		if pass == "false" {
			db.Del(title)
		} else {
			db.HSet(title, "havaAudit", true)
			db.HSet(title, "pass", true)
		}

	} else {
		money, _ := strconv.ParseInt(adminWork.Money, 10, 64)
		db.HIncrBy(title, "nowMoney", money)
	}

	information := projectService{
		Code:    200,
		Message: "更新成功",
	}
	context.JSON(http.StatusOK, information)

}

func DeleteProject(context *gin.Context) {
	//已筹集完毕项目删除
	var toolProject projectDataModel
	context.ShouldBindJSON(&toolProject)
	title := toolProject.Title
	db.Del(title)
	information := projectService{
		Code:    200,
		Message: "已删除",
	}
	context.JSON(http.StatusOK, information)
}

func AdminWork(context *gin.Context) {
	//获取未审核文件
	var i = 0
	var informations []projectService
	keys, err := db.Keys("*").Result()

	if err != nil {
		data := projectService{
			Code:    500,
			Message: "获取项目信息失败",
		}
		context.JSON(http.StatusInternalServerError, data)
		return
	}

	for _, key := range keys {
		var projects map[string]string
		projects, _ = db.HGetAll(key).Result()
		if projects["haveAudit"] == "0" {
			informations[i] = forUse(projects)
			i++
		}

	}

	context.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    informations,
	})
}

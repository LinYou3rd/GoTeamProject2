package user

import "github.com/jinzhu/gorm"

type (
	userModel struct {
		gorm.Model
		Name           string `json:"name"`
		Account        string `json:"account"`
		PasswordDigest string `json:"passwordDigest"`
		Identity       string `json:"identity"`
		Contact        string `json:"contact"`
	}

	userDataModel struct {
		gorm.Model
		Name     string `json:"name"`
		Account  string `json:"account"`
		Password string `json:"password"`
		Identity string `json:"identity"`
		Contact  string `json:"contact"`
		Mode     string `json:"mode"`
		Token    string `json:"token"`
	}

	userLoginModel struct {
		gorm.Model
		Code    int    `json:"code"`
		Token   string `json:"token"`
		Message string `json:"message"`
		Name    string `json:"name"`
	}
	//返回特定数据时补写一个结构体
)

func (userModel) TableName() string {
	return "users"
}

type (
	adminModel struct {
		gorm.Model
		Account        string `json:"account"`
		PasswordDigest string `json:"passwordDigest"`
	}

	adminDataModel struct {
		gorm.Model
		Account  string `json:"account"`
		Password string `json:"password"`
		Token    string `json:"token"`
		Mode     string `json:"mode"`
	}

	adminLoginModel struct {
		gorm.Model
		Code    int    `json:"code"`
		Token   string `json:"token"`
		Message string `json:"message"`
	}

	AdminWorkModel struct {
		Mode  string `json:"mode"`
		Title string `json:"title"`
		Pass  string `json:"pass"`
		Money string `json:"money"`
	}
)

func (adminModel) TableName() string {
	return "admins"
}

type tokenStature struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

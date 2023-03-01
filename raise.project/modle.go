package raise_project

import "strconv"

type (
	projectModel struct {
		Title       string `json:"title"`
		Introduce   string `json:"introduce"`
		Image       string `json:"image"`
		TargetMoney int    `json:"targetMoney"`
		NowMoney    int    `json:"nowMoney"`
		Pass        bool   `json:"pass" `
		HaveAudit   bool   `json:"haveAudit"`
	}

	projectDataModel struct {
		Title       string `json:"title"`
		Introduce   string `json:"introduce"`
		Image       string `json:"image"`
		TargetMoney int    `json:"targetMoney"`
		NowMoney    int    `json:"nowMoney"`
		Mode        string `json:"mode"`
		Pass        bool   `json:"pass" `
		HaveAudit   bool   `json:"haveAudit"`
	}

	projectService struct {
		Code        int    `json:"code"`
		Title       string `json:"title"`
		Introduce   string `json:"introduce"`
		Image       string `json:"image"`
		TargetMoney int    `json:"targetMoney"`
		NowMoney    int    `json:"nowMoney"`
		Message     string `json:"message"`
	}
)

func forSave(model projectDataModel) map[string]interface{} {
	data := make(map[string]interface{})
	data["introduce"] = model.Introduce
	data["image"] = model.Image
	data["targetMoney"] = model.TargetMoney
	data["nowMoney"] = model.NowMoney
	data["pass"] = model.Pass
	data["havaAudit"] = model.HaveAudit
	return data
}

func forUse(data map[string]string) projectService {
	targetMoney, _ := strconv.Atoi(data["targetMoney"])
	nowMoney, _ := strconv.Atoi(data["nowMoney"])
	information := projectService{
		Introduce:   data["introduce"],
		Image:       data["image"],
		TargetMoney: targetMoney,
		NowMoney:    nowMoney,
	}
	return information
}

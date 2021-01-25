package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

type ActtController struct {
	beego.Controller
}

// @Title ReportSource PainParts
// @Description 获取疼痛JSON
// @Param	id=default		format=json 	"获取疼痛JSON数据"
// @Success 200 json or xml
// @Failure 403 body is empty
// @router /painparts [get]
func (u *ActtController) PainParts() {
	req := httplib.Get(beego.AppConfig.String("acttapi"))
	req.Param("id", u.GetString("id"))
	req.Param("format", u.GetString("format"))
	switch u.GetString("format") {
	case "json":
		var result interface{}
		if err := req.ToJSON(&result); err != nil {
			u.Abort("500 "+ err.Error())
		}
		u.Data["json"] = result
		u.ServeJSON()
		break
	default:
		str, _ := req.String()
		u.Ctx.WriteString(str)
		break
	}
}
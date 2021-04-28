package controllers

import (
	"actt/models"
	"github.com/astaxie/beego"
	"net/http"
)

type UploadController struct {
	beego.Controller
}

// @Title GetAll
// @Description 文件上传
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [get]
func (u *UploadController) GetAll() {
	u.Data["json"] = "v1/upload/"
	u.ServeJSON()
}

// @Title Upload
// @Description 附件上传
// @Param	body		body 	models.Object	true		"The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UploadController) Post() {
	f, h, err := u.GetFile("files[]")
	if err != nil {
		http.Error(u.Ctx.ResponseWriter, err.Error(), http.StatusNoContent)
		return
	}
	defer f.Close()

	resp := make([]*models.OssResp, 1)
	resp[0], err = models.OssFileUpload(f, h)
	if err != nil {
		http.Error(u.Ctx.ResponseWriter, err.Error(), http.StatusBadRequest)
        return
	}

	u.SaveToFile("files[]", "upload/" + h.Filename)

	u.Data["json"] = resp
	u.ServeJSON()
}





package controllers

import (
	"actt/models"
	beego "github.com/beego/beego/v2/server/web"
	"io"
	"net/http"
	"os"
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
	u.Data["json"],_ = beego.AppConfig.String("AliOss::AccessKeyId")
	u.ServeJSON()
}

// @Title Upload
// @Description 附件上传
// @Param	body		body 	models.Object	true		"The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UploadController) Post() {
	u.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	files, err:=u.GetFiles("files[]")

	if err != nil {
		http.Error(u.Ctx.ResponseWriter, err.Error(), http.StatusNoContent)
		return
	}

	resps := make([]models.OssResp, len(files))

	for i, _ := range files {
		//for each fileheader, get a handle to the actual file
		file, err := files[i].Open()
		defer file.Close()
		if err != nil {
			http.Error(u.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
			return
		}

		resp, err := models.OssFileUpload(files[i].Filename, file)
		if err != nil {
			http.Error(u.Ctx.ResponseWriter, err.Error(), http.StatusBadRequest)
		}

		resps[i] = *resp

		//create destination file making sure the path is writeable.
		dst, err := os.Create("upload/" + resp.FileName)
		defer dst.Close()
		if err != nil {
			http.Error(u.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
			return
		}
		//copy the uploaded file to the destination file
		if _, err := io.Copy(dst, file); err != nil {
			http.Error(u.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	u.Data["json"] = resps
	u.ServeJSON()
}



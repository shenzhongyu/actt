package controllers

import (
	"actt/models"
	"github.com/astaxie/beego"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"time"
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
	time.Sleep(time.Second * 15)
	files, err:=u.GetFiles("files[]")
	if err != nil {
		http.Error(u.Ctx.ResponseWriter, err.Error(), http.StatusNoContent)
		return
	}

	res := make(chan models.OssResp, len(files))
	resp := make([]models.OssResp, len(files))

	for i, _ := range files {

		go func(n int, files []*multipart.FileHeader, res chan<- models.OssResp) {
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
			res <- *resp

			//create destination file making sure the path is writeable.
			dst, err := os.Create("upload/" + files[n].Filename)
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

		} (i, files, res)

		resp[i] = <-res
	}

	u.Data["json"] = resp
	u.ServeJSON()
}





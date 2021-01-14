package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type ErrorController struct {
	web.Controller
}


func (c *ErrorController) Error404() {
	c.Data["json"] = "page not found"
	c.ServeJSON()
}

func (c *ErrorController) Error500() {
	c.Data["json"] = "server error"
	c.ServeJSON()
}

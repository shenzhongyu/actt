package controllers

import (
	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}


func (c *ErrorController) Error404() {
	c.Data["json"] = "page not found"
	c.ServeJSON()
}

func (c *ErrorController) Error500() {
	c.Data["json"] = "server error"
	c.ServeJSON()
}

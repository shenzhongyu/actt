package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["actt/controllers:ActtController"] = append(beego.GlobalControllerRouter["actt/controllers:ActtController"],
        beego.ControllerComments{
            Method: "PainChecks",
            Router: "/painchecks",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["actt/controllers:ActtController"] = append(beego.GlobalControllerRouter["actt/controllers:ActtController"],
        beego.ControllerComments{
            Method: "PainParts",
            Router: "/painparts",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["actt/controllers:UploadController"] = append(beego.GlobalControllerRouter["actt/controllers:UploadController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["actt/controllers:UploadController"] = append(beego.GlobalControllerRouter["actt/controllers:UploadController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}

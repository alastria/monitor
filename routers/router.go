// @APIVersion 1.0.0
// @Title Alastria node monitor API
// @Description Monitoring system for Alastria nodes
// @Contact alfon.rocha@gmail.com
// @TermsOfServiceUrl http://alastria.io/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	// "github.com/alastria/monitor/controllers"
	"monitor/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/monitor",
			beego.NSInclude(
				&controllers.MonitorController{},
			),
		),

		beego.NSNamespace("/node",
			beego.NSInclude(
				&controllers.NodeController{},
			),
		),
	)
	beego.AddNamespace(ns)
}

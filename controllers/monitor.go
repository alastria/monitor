package controllers

import (
	"encoding/json"
	"monitor/models"

	"github.com/astaxie/beego"
)

// Operations about object
type MonitorController struct {
	beego.Controller
}

// @Title RestPostStatus
// @Description Check status of the monitor REST
// @Param	body		body 	models.StatusReq	true		"Request status of the monitor"
// @Success 200  {map[string]string} Status information
// @Failure 403 overload
// @router /status [post]
func (m *MonitorController) RestPostStatus() {
	var r models.StatusReq
	json.Unmarshal(m.Ctx.Input.RequestBody, &r)
	// m.Data["json"] = &StatusReq
	m.Data["json"] = map[string]string{"status": "ok", "testString": (&r).TestString}
	m.ServeJSON()
}

// @Title RestGetStatus
// @Description Check status of the monitor REST
// @Success 200 {status} string Status
// @Failure 403 : overload
// @router /status [get]
func (m *MonitorController) RestGetStatus() {
	output := make(map[string]string)
	output["status"] = "ok"
	m.Data["json"] = &output
	m.ServeJSON()
}

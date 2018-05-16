package controllers

import (
	"encoding/json"

	"github.com/alastria/monitor/lib"
	"github.com/alastria/monitor/models"

	"github.com/astaxie/beego"
)

// Monitor operations
type MonitorController struct {
	beego.Controller
}

// @Title RestPostStatus
// @Description Test simple POST method over the monitor
// @Param	body		body 	models.StatusReq	true		"String to test simple POST method"
// @Success 200  {map[string]string} Status information
// @Failure 403 overload
// @router /pingpong [post]
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

	_, port1 := lib.RunCommand("lsof -iTCP -sTCP:LISTEN -P -n | grep *:21000 | awk '{print $8 $9 $10}'")
	_, port2UDP := lib.RunCommand("lsof -iUDP -P -n | grep *:21000 | awk '{print $8 $9 $10}'")
	_, port2TCP := lib.RunCommand("lsof -iTCP -sTCP:LISTEN -P -n | grep *:22000 | awk '{print $8 $9 $10}'")
	_, port3 := lib.RunCommand("lsof -iTCP -sTCP:LISTEN -P -n | grep *:9000 | awk '{print $8 $9 $10}'")

	if port1 != "" {
		output["port21000"] = "up"
	} else {
		output["port21000"] = "down"
	}

	if port2UDP != "" && port2TCP != "" {
		output["port22000"] = "up"
	} else {
		output["port22000"] = "down"
	}

	if port3 != "" {
		output["port9000"] = "up"
	} else {
		output["port9000"] = "down"
	}

	output["status"] = "ok"
	m.Data["json"] = &output
	m.ServeJSON()
}

// @Title GetVersion
// @Description Check monitor version
// @Success 200 {status, version} string Status, string Version
// @Failure 403 : overload
// @router /version [get]
func (m *MonitorController) GetVersion() {
	output := make(map[string]string)
	_, current := lib.CurrentMonitorVersion()
	_, latest := lib.LatestMonitorVersion()
	if current != latest {
		output["status"] = "outdated"
	} else {
		output["status"] = "ok"
	}

	output["version"] = current
	m.Data["json"] = &output
	m.ServeJSON()
}

// @Title GetVersionUpdate
// @Description Check monitor version.
//	involved we decide to user GET.
// @Success 202 {status, version} string Status (latest | updated)
// @Failure 403 : overload
// @router /update [get]
func (m *MonitorController) GetVersionUpdate() {
	output := make(map[string]string)
	_, current := lib.CurrentMonitorVersion()
	_, latest := lib.LatestMonitorVersion()

	if current != latest {
		if lib.UpdateMonitor() {
			output["status"] = "updated"
		} else {
			output["status"] = "error"
		}
	} else {
		output["status"] = "latest"
	}

	m.Data["json"] = &output
	m.ServeJSON()
}

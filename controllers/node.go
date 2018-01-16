package controllers

import (
	"monitor/lib"

	"github.com/astaxie/beego"
)

// Operations about object
type NodeController struct {
	beego.Controller
}

// @Title UpdateFiles
// @Description Updates files and restart node
// @Success 200 {status} string
// @Failure 403 : error
// @router /update [get]
func (m *NodeController) UpdateFiles() {
	output := make(map[string]string)

	if lib.Update() {
		output["status"] = "ok"
	} else {
		output["status"] = "error occurred"
	}

	m.Data["json"] = &output
	m.ServeJSON()
}

// @Title StartNode
// @Description Starts the Alastria node
// @Success 200 {status} string
// @Failure 403 : error
// @router /start [get]
func (m *NodeController) StartNode() {
	output := make(map[string]string)
	if lib.Start() {
		output["status"] = "ok"
	} else {
		output["status"] = "error occurred"
	}

	m.Data["json"] = &output
	m.ServeJSON()
}

// @Title StopNode
// @Description Stops the Alastria node
// @Success 200 {status} string
// @Failure 403 : error
// @router /stop [get]
func (m *NodeController) StopNode() {
	output := make(map[string]string)
	if lib.Stop() {
		output["status"] = "ok"
	} else {
		output["status"] = "error occurred"
	}

	m.Data["json"] = &output
	m.ServeJSON()
}

// @Title restartNode
// @Description Restarts Alastria node
// @Success 200 {status} string
// @Failure 403 : error
// @router /restart [get]
func (m *NodeController) RestartNode() {
	output := make(map[string]string)
	if lib.Restart() {
		output["status"] = "ok"
	} else {
		output["status"] = "error occurred"
	}

	m.Data["json"] = &output
	m.ServeJSON()
}

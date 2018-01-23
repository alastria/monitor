package controllers

import (
	"encoding/json"

	"github.com/alastria/monitor/lib"
	"github.com/alastria/monitor/models"

	"github.com/astaxie/beego"
)

// Node operations
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

// @Title statusNode
// @Description Status Alastria node
// @Success 200 {status} string
// @Failure 403 : error
// @router /status [get]
func (m *NodeController) StatusNode() {
	output := make(map[string]string)
	if string(lib.Status()) != "" {
		output["status"] = "ok"
	} else {
		output["status"] = "stopped"
	}

	m.Data["json"] = &output
	m.ServeJSON()
}

// @Title ProposeCandidate
// @Description Propose new validator candidate
// @Param	body		body 	models.ProposeReq	true		"Propose a new candidate"
// @Success 200  {status} string
// @Failure 403 error in propose
// @router /propose [post]
func (m *NodeController) ProposeCandidate() {
	var r models.ProposeReq
	json.Unmarshal(m.Ctx.Input.RequestBody, &r)
	if lib.Propose((&r).Candidate, (&r).Value) {
		m.Data["json"] = map[string]string{"status": "ok"}
	} else {
		m.Data["json"] = map[string]string{"status": "propose failed"}
	}
	m.ServeJSON()
}

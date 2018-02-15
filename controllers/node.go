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

// @Title CleanStartNode
// @Description Starts the Alastria node clean
// @Success 200 {status} string
// @Failure 403 : error
// @router /cleanstart [get]
func (m *NodeController) CleanStartNode() {
	output := make(map[string]string)
	if lib.CleanStart() {
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

// @Title NodeRestartNetwork
// @Description Restart complete network for the node
// @Param	body		body 	models.RestartNetReq	true		"Restart the whole network"
// @Success 200  {status} string
// @Failure 403 error in restart
// @router /restartNetwork [post]
func (m *NodeController) NodeRestartNetwork() {
	var r models.RestartNetReq
	json.Unmarshal(m.Ctx.Input.RequestBody, &r)
	if lib.RestartNetwork((&r).NodeType, (&r).NodeName) {
		m.Data["json"] = map[string]string{"status": "ok"}
	} else {
		m.Data["json"] = map[string]string{"status": "restart failed"}
	}
	m.ServeJSON()
}

// @Title getLogs
// @Description Get coinbase of the node
// @Success 200 {status} coinbase
// @Failure 403 : error
// @router /coinbase [get]
func (m *NodeController) Coinbase() {
	output := make(map[string]string)
	ok, data := lib.GetCoinbase()
	if ok {
		output["data"] = data
	} else {
		output["status"] = "error occurred"
	}

	m.Data["json"] = &output
	m.ServeJSON()
}

// @Title getLogs
// @Description Get logs for the node
// @Success 200 {status} logData
// @Failure 403 : error
// @router /logsraw [get]
func (m *NodeController) GetLogs() {
	output := make(map[string]string)
	ok, data := lib.GetLog()
	if ok {
		output["data"] = data
	} else {
		output["status"] = "error occurred"
	}

	m.Data["json"] = &output
	m.ServeJSON()
}

// @Title GetVersion
// @Description Get current version of the node
// @Success 200 {status} logStatus {version} nodeVersion
// @Failure 403 : error
// @router /version [get]
func (m *NodeController) GetVersion() {
	output := make(map[string]string)
	ok, data := lib.NodeVersion()
	if ok {
		output["status"] = "ok"
		output["version"] = data
	} else {
		output["status"] = "error occurred"
	}

	m.Data["json"] = &output
	m.ServeJSON()
}

// @Title NodeLastRestart
// @Description Date/Hour of latest restart of the node
// @Success 200 {status} Status {date} lastUpdate
// @Failure 403 : error
// @router /lastrestart [get]
func (m *NodeController) NodeLastRestart() {
	output := make(map[string]string)
	ok, data := lib.LastNodeRestart()
	if ok {
		output["status"] = "ok"
		output["lastUpdate"] = data
	} else {
		output["status"] = "error occurred"
	}

	m.Data["json"] = &output
	m.ServeJSON()
}

// @Title getLogsJson
// @Description Get logs with a fancy formatted JSON
// @Success 200 {status} logData
// @Failure 403 : error
// @router /logsjson [get]
func (m *NodeController) GetLogsJson() {
	output := make(map[string]string)

	_, port1 := lib.RunCommand("lsof -i | grep *:21000 | awk '{print $8 $9 $10}'")
	_, port2 := lib.RunCommand("lsof -i | grep *:22000 | awk '{print $8 $9 $10}'")
	_, port3 := lib.RunCommand("lsof -i | grep *:9000 | awk '{print $8 $9 $10}'")

	_, nodeInfo := lib.RunCommand("geth -exec 'admin.nodeInfo' attach ~/alastria/data/geth.ipc")
	_, peers := lib.RunCommand("geth -exec 'admin.peers' attach ~/alastria/data/geth.ipc")
	_, blockNumber := lib.RunCommand("geth -exec 'eth.blockNumber' attach ~/alastria/data/geth.ipc")
	_, mining := lib.RunCommand("geth -exec 'eth.mining' attach ~/alastria/data/geth.ipc")
	_, syncing := lib.RunCommand("geth -exec 'eth.syncing' attach ~/alastria/data/geth.ipc")
	_, pendingTransactions := lib.RunCommand("geth -exec 'eth.pendingTransactions' attach ~/alastria/data/geth.ipc")
	_, candidates := lib.RunCommand("geth -exec 'istanbul.candidates' attach ~/alastria/data/geth.ipc")
	_, getValidators := lib.RunCommand("geth -exec 'istanbul.getValidators()' attach ~/alastria/data/geth.ipc")
	_, peerCount := lib.RunCommand("geth -exec 'net.peerCount' attach ~/alastria/data/geth.ipc")
	_, netVersion := lib.RunCommand("geth -exec 'net.version' attach ~/alastria/data/geth.ipc")
	_, txPool := lib.RunCommand("geth -exec 'txpool.content' attach ~/alastria/data/geth.ipc")

	output["port1"] = port1
	output["port2"] = port2
	output["port3"] = port3

	output["nodeInfo"] = nodeInfo
	output["peers"] = peers
	output["blockNumber"] = blockNumber
	output["mining"] = mining
	output["syncing"] = syncing
	output["pendingTransactions"] = pendingTransactions
	output["candidates"] = candidates
	output["getValidators"] = getValidators
	output["peerCount"] = peerCount
	output["netVersion"] = netVersion
	output["txPool"] = txPool

	m.Data["json"] = &output
	m.ServeJSON()
}

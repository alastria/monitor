package controllers

import (
	"encoding/json"
	"net/http"

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
// @Failure 405 Invalid Method
// @router /update [get]
func (m *NodeController) UpdateFiles() {
	output := make(map[string]string)

	if m.Ctx.Input.IsGet() {
		if lib.Update() {
			output["status"] = "ok"
		} else {
			output["status"] = "error occurred"
		}
		m.Data["json"] = &output
	} else {
		m.Ctx.Output.Header("Content-Type", "application/json")
		m.Data["json"] = map[string]string{"status": "Error", "testString": "Method Not Allowed"}
		m.Ctx.Output.SetStatus(http.StatusMethodNotAllowed)
	}
	m.ServeJSON()
}

// @Title StartNode
// @Description Starts the Alastria node
// @Success 200 {status} string
// @Failure 403 : error
// @Failure 405 Invalid Method
// @router /start/default [get]
func (m *NodeController) StartNode() {
	output := make(map[string]string)

	if m.Ctx.Input.IsGet() {
		if lib.Start() {
			output["status"] = "ok"
		} else {
			output["status"] = "error occurred"
		}

		m.Data["json"] = &output
	} else {
		m.Ctx.Output.Header("Content-Type", "application/json")
		m.Data["json"] = map[string]string{"status": "Error", "testString": "Method Not Allowed"}
		m.Ctx.Output.SetStatus(http.StatusMethodNotAllowed)
	}
	m.ServeJSON()
}

// @Title CleanStartNode
// @Description Starts the Alastria node clean
// @Success 200 {status} string
// @Failure 403 : error
// @Failure 405 Invalid Method
// @router /start/clean [get]
func (m *NodeController) CleanStartNode() {
	output := make(map[string]string)

	if m.Ctx.Input.IsGet() {
		if lib.CleanStart() {
			output["status"] = "ok"
		} else {
			output["status"] = "error occurred"
		}

		m.Data["json"] = &output
	} else {
		m.Ctx.Output.Header("Content-Type", "application/json")
		m.Data["json"] = map[string]string{"status": "Error", "testString": "Method Not Allowed"}
		m.Ctx.Output.SetStatus(http.StatusMethodNotAllowed)
	}
	m.ServeJSON()
}

// @Title TransactionsClean
// @Description Cleans the transaction queue from the node
// @Success 200 {status} string
// @Failure 403 : error
// @Failure 405 Invalid Method
// @router /transactions [delete]
func (m *NodeController) TransactionsClean() {
	output := make(map[string]string)

	if m.Ctx.Input.IsDelete() {
		if lib.CleanTransactions() {
			output["status"] = "ok"
		} else {
			output["status"] = "error occurred"
		}

		m.Data["json"] = &output
	} else {
		m.Ctx.Output.Header("Content-Type", "application/json")
		m.Data["json"] = map[string]string{"status": "Error", "testString": "Method Not Allowed"}
		m.Ctx.Output.SetStatus(http.StatusMethodNotAllowed)
	}
	m.ServeJSON()
}

// @Title StopNode
// @Description Stops the Alastria node
// @Success 200 {status} string
// @Failure 403 : error
// @Failure 405 Invalid Method
// @router /stop [post]
func (m *NodeController) StopNode() {
	output := make(map[string]string)

	if m.Ctx.Input.IsPost() {
		if lib.Stop() {
			output["status"] = "ok"
		} else {
			output["status"] = "error occurred"
		}

		m.Data["json"] = &output
	} else {
		m.Ctx.Output.Header("Content-Type", "application/json")
		m.Data["json"] = map[string]string{"status": "Error", "testString": "Method Not Allowed"}
		m.Ctx.Output.SetStatus(http.StatusMethodNotAllowed)
	}
	m.ServeJSON()
}

// @Title restartNode
// @Description Restarts Alastria node
// @Success 200 {status} string
// @Failure 403 : error
// @Failure 405 Invalid Method
// @router /restart [post]
func (m *NodeController) RestartNode() {
	output := make(map[string]string)

	if m.Ctx.Input.IsPost() {
		if lib.Restart() {
			output["status"] = "ok"
		} else {
			output["status"] = "error occurred"
		}

		m.Data["json"] = &output
	} else {
		m.Ctx.Output.Header("Content-Type", "application/json")
		m.Data["json"] = map[string]string{"status": "Error", "testString": "Method Not Allowed"}
		m.Ctx.Output.SetStatus(http.StatusMethodNotAllowed)
	}
	m.ServeJSON()
}

// @Title statusNode
// @Description Status Alastria node
// @Success 200 {status} string
// @Failure 403 : error
// @Failure 405 Invalid Method
// @router /status [get]
func (m *NodeController) StatusNode() {
	output := make(map[string]string)

	if m.Ctx.Input.IsGet() {
		if string(lib.Status()) != "" {
			output["status"] = "ok"
		} else {
			output["status"] = "stopped"
		}

		m.Data["json"] = &output
	} else {
		m.Ctx.Output.Header("Content-Type", "application/json")
		m.Data["json"] = map[string]string{"status": "Error", "testString": "Method Not Allowed"}
		m.Ctx.Output.SetStatus(http.StatusMethodNotAllowed)
	}
	m.ServeJSON()
}

// @Title ProposeCandidate
// @Description Propose new validator candidate
// @Param	body		body 	models.ProposeReq	true		"Propose a new candidate"
// @Success 200  {status} string
// @Failure 403 error in propose
// @Failure 405 Invalid Method
// @router /propose [post]
func (m *NodeController) ProposeCandidate() {
	var r models.ProposeReq

	if m.Ctx.Input.IsPost() {
		json.Unmarshal(m.Ctx.Input.RequestBody, &r)
		if lib.Propose((&r).Candidate, (&r).Value) {
			m.Data["json"] = map[string]string{"status": "ok"}
		} else {
			m.Data["json"] = map[string]string{"status": "propose failed"}
		}
	} else {
		m.Ctx.Output.Header("Content-Type", "application/json")
		m.Data["json"] = map[string]string{"status": "Error", "testString": "Method Not Allowed"}
		m.Ctx.Output.SetStatus(http.StatusMethodNotAllowed)
	}
	m.ServeJSON()
}

// @Title NodeRestartNetwork
// @Description Restart complete network for the node
// @Param	body		body 	models.RestartNetReq	true		"Restart the whole network"
// @Success 200  {status} string
// @Failure 403 error in restart
// @Failure 405 Invalid Method
// @router /network/restart [post]
func (m *NodeController) NodeRestartNetwork() {
	var r models.RestartNetReq
	if m.Ctx.Input.IsPost() {
		json.Unmarshal(m.Ctx.Input.RequestBody, &r)
		if lib.RestartNetwork((&r).NodeType, (&r).NodeName) {
			m.Data["json"] = map[string]string{"status": "ok"}
		} else {
			m.Data["json"] = map[string]string{"status": "restart failed"}
		}
	} else {
		m.Ctx.Output.Header("Content-Type", "application/json")
		m.Data["json"] = map[string]string{"status": "Error", "testString": "Method Not Allowed"}
		m.Ctx.Output.SetStatus(http.StatusMethodNotAllowed)
	}
	m.ServeJSON()
}

// @Title getCoinbase
// @Description Get coinbase of the node
// @Success 200 {status} coinbase
// @Failure 403 : error
// @Failure 405 Invalid Method
// @router /coinbase [get]
func (m *NodeController) Coinbase() {
	output := make(map[string]string)
	if m.Ctx.Input.IsGet() {
		ok, data := lib.GetCoinbase()
		if ok {
			output["data"] = data
		} else {
			output["status"] = "error occurred"
		}

		m.Data["json"] = &output
	} else {
		m.Ctx.Output.Header("Content-Type", "application/json")
		m.Data["json"] = map[string]string{"status": "Error", "testString": "Method Not Allowed"}
		m.Ctx.Output.SetStatus(http.StatusMethodNotAllowed)
	}
	m.ServeJSON()
}

// @Title MineStart
// @Description If the node is not mining start the miner
// @Success 200 {status} mining
// @Failure 403 : error
// @Failure 405 Invalid Method
// @router /ismining [post]
func (m *NodeController) MineStart() {
	output := make(map[string]string)

	if m.Ctx.Input.IsPost() {
		ok := lib.StartMining()
		if ok {
			output["status"] = "ok"
		} else {
			output["status"] = "error occurred"
		}

		m.Data["json"] = &output
	} else {
		m.Ctx.Output.Header("Content-Type", "application/json")
		m.Data["json"] = map[string]string{"status": "Error", "testString": "Method Not Allowed"}
		m.Ctx.Output.SetStatus(http.StatusMethodNotAllowed)
	}
	m.ServeJSON()
}

// @Title getLogs
// @Description Get logs for the node
// @Success 200 {status} logData
// @Failure 403 : error
// @Failure 405 Invalid Method
// @router /log/raw [get]
func (m *NodeController) GetLogs() {
	output := make(map[string]string)

	if m.Ctx.Input.IsGet() {
		ok, data := lib.GetLog()
		if ok {
			output["data"] = data
		} else {
			output["status"] = "error occurred"
		}

		m.Data["json"] = &output
	} else {
		m.Ctx.Output.Header("Content-Type", "application/json")
		m.Data["json"] = map[string]string{"status": "Error", "testString": "Method Not Allowed"}
		m.Ctx.Output.SetStatus(http.StatusMethodNotAllowed)
	}
	m.ServeJSON()
}

// @Title GetVersion
// @Description Get current version of the node
// @Success 200 {status} logStatus {version} nodeVersion
// @Failure 403 : error
// @Failure 405 Invalid Method
// @router /version [get]
func (m *NodeController) GetVersion() {
	output := make(map[string]string)

	if m.Ctx.Input.IsGet() {
		ok, data := lib.NodeVersion()
		if ok {
			output["status"] = "ok"
			output["version"] = data
		} else {
			output["status"] = "error occurred"
		}

		m.Data["json"] = &output
	} else {
		m.Ctx.Output.Header("Content-Type", "application/json")
		m.Data["json"] = map[string]string{"status": "Error", "testString": "Method Not Allowed"}
		m.Ctx.Output.SetStatus(http.StatusMethodNotAllowed)
	}
	m.ServeJSON()
}

// @Title NodeLastRestart
// @Description Date/Hour of latest restart of the node
// @Success 200 {status} Status {date} lastUpdate
// @Failure 403 : error
// @Failure 405 Invalid Method
// @router /start/latest [get]
func (m *NodeController) NodeLastRestart() {
	output := make(map[string]string)

	if m.Ctx.Input.IsGet() {
		ok, data := lib.LastNodeRestart()
		if ok {
			output["status"] = "ok"
			output["lastUpdate"] = data
		} else {
			output["status"] = "error occurred"
		}

		m.Data["json"] = &output
	} else {
		m.Ctx.Output.Header("Content-Type", "application/json")
		m.Data["json"] = map[string]string{"status": "Error", "testString": "Method Not Allowed"}
		m.Ctx.Output.SetStatus(http.StatusMethodNotAllowed)
	}
	m.ServeJSON()
}

// @Title getLogsJson
// @Description Get logs with a fancy formatted JSON
// @Success 200 {status} logData
// @Failure 403 : error
// @Failure 405 Invalid Method
// @router /log/json [get]
func (m *NodeController) GetLogsJson() {
	output := make(map[string]string)

	if m.Ctx.Input.IsGet() {
		_, port1 := lib.RunCommand("lsof -iTCP -sTCP:LISTEN -P -n | grep *:21000 | awk '{print $8 $9 $10}'")
		_, port2UDP := lib.RunCommand("lsof -iUDP -P -n | grep *:21000 | awk '{print $8 $9 $10}'")
		_, port2TCP := lib.RunCommand("lsof -iTCP -sTCP:LISTEN -P -n | grep *:22000 | awk '{print $8 $9 $10}'")
		_, port3 := lib.RunCommand("lsof -iTCP -sTCP:LISTEN -P -n | grep *:9000 | awk '{print $8 $9 $10}'")

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
		_, coinbase := lib.RunCommand("geth --exec 'eth.coinbase' attach ~/alastria/data/geth.ipc")

		output["port1"] = port1
		output["port2"] = port2UDP + port2TCP
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
		output["coinbase"] = coinbase

		m.Data["json"] = &output
	} else {
		m.Ctx.Output.Header("Content-Type", "application/json")
		m.Data["json"] = map[string]string{"status": "Error", "testString": "Method Not Allowed"}
		m.Ctx.Output.SetStatus(http.StatusMethodNotAllowed)
	}
	m.ServeJSON()
}

// @Title getIstanbulLog
// @Description Get Istanbul related logs with a fancy formatted JSON
// @Success 200 {status} logData
// @Failure 403 : error
// @Failure 405 Invalid Method
// @router /log/istanbul [get]
func (m *NodeController) GetIstanbulLog() {
	output := make(map[string]string)
	if m.Ctx.Input.IsGet() {
		_, enode := lib.RunCommand("geth --exec 'admin.nodeInfo' attach ~/alastria/data/geth.ipc | grep enode")
		_, coinbase := lib.RunCommand("geth --exec 'eth.coinbase' attach ~/alastria/data/geth.ipc")
		_, candidates := lib.RunCommand("geth --exec 'istanbul.candidates' attach ~/alastria/data/geth.ipc")
		_, validators := lib.RunCommand("geth --exec 'istanbul.getValidators()' attach ~/alastria/data/geth.ipc")

		output["enode"] = enode
		output["coinbase"] = coinbase
		output["candidates"] = candidates
		output["validators"] = validators
		m.Data["json"] = &output
	} else {
		m.Ctx.Output.Header("Content-Type", "application/json")
		m.Data["json"] = map[string]string{"status": "Error", "testString": "Method Not Allowed"}
		m.Ctx.Output.SetStatus(http.StatusMethodNotAllowed)
	}
	m.ServeJSON()
}

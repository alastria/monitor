package lib

import (
	"bytes"
	"fmt"
	//	"math/big"
	"encoding/hex"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/hashicorp/go-getter"
	//"github.com/alastria/monitor/services"
)

var log *logs.BeeLogger

var homeDir = os.Getenv("HOME")

var STATIC_NODES = homeDir + "/alastria/data/static-nodes.json"
var PERMISSIONED_NODES = homeDir + "/alastria/data/permissioned-nodes.json"

var IDENTITY, NODE_TYPE, STATIC, PERMISSIONED string

// Restart an Alastria node
func Restart() bool {
	Stop()
	Start()
	return true
}

func RunCommand(command string) (ok bool, salida string) {
	cmd := exec.Command("bash", "-c", command)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		ok = false
		salida = "Error: " + fmt.Sprint(err) + ": " + stderr.String()
		fmt.Println(salida)
		// log.Debug(salida)
		return
	}
	ok = true
	salida = out.String()
	// log.Debug(salida)
	fmt.Println("Result: " + salida)
	return
}

func RunCommandBackground(command string) (ok bool, salida string) {
	cmd := exec.Command("bash", "-c", command)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Start()
	if err != nil {
		ok = false
		salida = "Error: " + fmt.Sprint(err) + ": " + stderr.String()
		fmt.Println(salida)
		// log.Debug(salida)
		return
	}
	ok = true
	salida = out.String()
	// log.Debug(salida)
	fmt.Println("Result: " + salida)
	return
}

// Stop an Alastria node
func Stop() (ok bool) {
	_, _ = RunCommand(homeDir + "/alastria-node/scripts/stop.sh")
	return true
}

// Start an Alastria node
func Start() (ok bool) {
	ok, _ = RunCommand(homeDir + "/alastria-node/scripts/start.sh")
	return
}

// Clean Start an Alastria node
func CleanStart() (ok bool) {
	ok, _ = RunCommand(homeDir + "/alastria-node/scripts/start.sh clean")
	return
}

// Update config files and restart an Alastria node
func Update() bool {
	var err error
	RunCommand("cd " + homeDir + "/alastria-node && git pull")
	// log.Debug("%s, %s, %s, %s", IDENTITY, NODE_TYPE, STATIC, PERMISSIONED)
	stfile, static := getGithub("https://raw.githubusercontent.com/alastria/alastria-node/feature/ibft/data/static-nodes.json")
	pmfile, permissioned := getGithub("https://raw.githubusercontent.com/alastria/alastria-node/feature/ibft/data/permissioned-nodes_" + NODE_TYPE + ".json")
	if strings.Compare(static, STATIC) != 0 || strings.Compare(permissioned, PERMISSIONED) != 0 {

		// log.Trace("Son distintos")
		if strings.Compare(static, STATIC) != 0 {
			// log.Trace("Actualizando static-nodes")
			copy(stfile, STATIC_NODES)
		}
		if strings.Compare(permissioned, PERMISSIONED) != 0 {
			// log.Trace("Actualizando permissioned-nodes")
			copy(pmfile, PERMISSIONED_NODES)
		}
		// log.Debug(strings.Trim(static, "]"), strings.Trim(STATIC, "]"))
		if !strings.Contains(strings.Trim(static, "]"), strings.Trim(STATIC, "]")) ||
			!strings.Contains(strings.Trim(permissioned, "]"), strings.Trim(PERMISSIONED, "]")) {
			// log.Trace("Hay que reiniciar el nodo...")
			RunCommand(homeDir + "/alastria-node/scripts/stop.sh")
			time.Sleep(15000 * time.Millisecond)
			RunCommandBackground(homeDir + "/alastria-node/scripts/start.sh clean")
		}
	}
	if err != nil {
		return false
	}
	return true
}

// Log an Alastria node
func GetLog() (ok bool, data string) {
	ok, data = RunCommand(homeDir + "/alastria-node/scripts/log.sh")
	return
}

// Get current Version of the monitor
func CurrentMonitorVersion() (ok bool, version string) {
	// ok, version = RunCommand(homeDir + "/alastria-node/scripts/monitor.sh latest")
	ok, version = RunCommand("cd " + homeDir + "/alastria/monitor/src/github.com/alastria/monitor && git tag")
	return
}

// Get latest Version of the monitor
func LatestMonitorVersion() (ok bool, version string) {
	ok, version = RunCommand(homeDir + "/alastria-node/scripts/monitor.sh latest")
	return
}

// Get latest Version of the monitor
func UpdateMonitor() (ok bool) {
	ok1, _ := RunCommandBackground(homeDir + "/alastria-node/scripts/monitor.sh build")
	ok2, _ := RunCommandBackground(homeDir + "/alastria-node/scripts/monitor.sh start")
	if ok1 && ok2 {
		return true
	}
	return false
}

//Restart the complete network
func RestartNetwork(nodeType string, nodeName string) (ok bool) {
	_, _ = RunCommand("cd " + homeDir + "/alastria-node/ && git pull")
	_, _ = RunCommand(homeDir + "/alastria-node/scripts/stop.sh")
	ok1, _ := RunCommand("cd " + homeDir + "/alastria-node/scripts && ./init.sh backup " + nodeType + " " + nodeName)
	ok2, _ := RunCommand(homeDir + "/alastria-node/scripts/start.sh clean")
	if ok1 && ok2 {
		return true
	}
	return false
}

// Propose new candidate
func Propose(candidate string, value string) (ok bool) {
	cmdStr := "geth --exec 'istanbul.propose(\"" + candidate + "\", " + value + ")' attach http://localhost:22000"
	ok, _ = RunCommand(cmdStr)
	return
}

// Gets coinbase from node
func GetCoinbase() (ok bool, data string) {
	cmdStr := "geth --exec 'eth.coinbase' attach http://localhost:22000"
	ok, data = RunCommand(cmdStr)
	return
}

// Non-returning Update function for its use in CRON
func UpdateCron() {
	Update()
}

//Compute Status for a node
func Status() (salida string) {
	_, salida = RunCommand("ps aux | grep geth  | grep alastria/data | grep -v grep | awk '{print $2}'")
	return
}

//Last node/geth Restart
func LastNodeRestart() (ok bool, salida string) {
	ok, salida = RunCommand("ps aux | grep geth  | grep alastria/data | grep -v grep | awk '{print $9}'")
	return
}

//Node/Geth/Quorum Version
func NodeVersion() (ok bool, salida string) {
	ok, salida = RunCommand("geth version | grep geth")
	return
}

func getGithub(url string) (filename, contenido string) {
	filename = tempFileName("monitor", ".json")
	err := getter.GetFile(filename, url)
	if err != nil {
		// log.Warn("getGithub: %s", err)
	}
	if err == nil {
		contenido = getFile(filename)
	}
	return
}

func getFile(fichero string) (contenido string) {
	data, _ := ioutil.ReadFile(fichero)
	contenido = string(data)
	return
}

func copy(stfrom, stto string) {
	// log.Debug(stfrom, stto)
	from, err := os.Open(stfrom)
	if err != nil {
		// log.Warn("copy: ", err)
	}
	defer from.Close()
	to, err := os.Open(stto)
	if err != nil {
		// log.Warn("copy: ", err)
	}
	defer to.Close()
	io.Copy(from, to)
}

func tempFileName(prefix, suffix string) string {
	randBytes := make([]byte, 16)
	rand.Read(randBytes)
	return filepath.Join(os.TempDir(), prefix+hex.EncodeToString(randBytes)+suffix)
}

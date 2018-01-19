package lib

import (
	"bytes"
	"fmt"
	"log"
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

	"github.com/hashicorp/go-getter"
	//"github.com/alastria/monitor/services"
)

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

func debugCommand() {
	cmd := exec.Command(homeDir + "/alastria-node/scripts/stop.sh")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error: " + fmt.Sprint(err) + ": " + stderr.String())
		return
	}
	fmt.Println("Result: " + out.String())
}

// Stop an Alastria node
func Stop() bool {
	out, err := exec.Command(homeDir + "/alastria-node/scripts/stop.sh").Output()
	fmt.Println(out, err)
	time.Sleep(100 * time.Millisecond)
	if err != nil {
		return false
	}
	return true
}

// Start an Alastria node
func Start() bool {
	out, err := exec.Command(homeDir + "/alastria-node/scripts/start.sh").Output()
	fmt.Println(out, err)
	time.Sleep(100 * time.Millisecond)
	if err != nil {
		return false
	}
	return true
}

// Update config files and restart an Alastria node
func Update() bool {
	var err error
	fmt.Printf("%s, %s, %s, %s", IDENTITY, NODE_TYPE, STATIC, PERMISSIONED)
	stfile, static := getGithub("https://raw.githubusercontent.com/alastria/alastria-node/feature/ibft/data/static-nodes.json")
	pmfile, permissioned := getGithub("https://raw.githubusercontent.com/alastria/alastria-node/feature/ibft/data/permissioned-nodes_" + NODE_TYPE + ".json")
	if strings.Compare(static, STATIC) != 0 || strings.Compare(permissioned, PERMISSIONED) != 0 {

		fmt.Println("Son distintos")
		if strings.Compare(static, STATIC) != 0 {
			fmt.Println("Actualizando static-nodes")
			copy(stfile, STATIC_NODES)
		}
		if strings.Compare(permissioned, PERMISSIONED) != 0 {
			fmt.Println("Actualizando permissioned-nodes")
			copy(pmfile, PERMISSIONED_NODES)
		}
		fmt.Println(strings.Trim(static, "]"), strings.Trim(STATIC, "]"))
		if !strings.Contains(strings.Trim(static, "]"), strings.Trim(STATIC, "]")) ||
			!strings.Contains(strings.Trim(permissioned, "]"), strings.Trim(PERMISSIONED, "]")) {
			fmt.Println("Hay que reiniciar el nodo...")
			out, err := exec.Command(homeDir + "/alastria-node/scripts/stop.sh").Output()
			fmt.Println(out, err)
			time.Sleep(15000 * time.Millisecond)
			out, err = exec.Command(homeDir+"/alastria-node/scripts/start.sh", "all").Output()
			fmt.Println(out, err)
		}
	}
	if err != nil {
		return false
	}
	return true
}

// Propose new candidate
func Propose(candidate string, value string) bool {
	cmdStr := "geth --exec 'istanbul.propose(\"" + candidate + "\", " + value + ")' attach http://localhost:22000"
	out, err := exec.Command(cmdStr).Output()
	fmt.Println(out, err)
	time.Sleep(100 * time.Millisecond)
	if err != nil || string(out) != "null" {
		return false
	}
	return true
}

// Non-returning Update function for its use in CRON
func UpdateCron() {
	Update()
}

//Compute Status for a node
func Status() bool {
	out, err := exec.Command("ps aux | grep geth  | grep alastria/data | grep -v grep | awk '{print $2}')").Output()
	if err != nil {
		log.Fatal(err)
	}
	if string(out) != "" {
		return true
	}
	return false
}

func getGithub(url string) (filename, contenido string) {
	filename = tempFileName("monitor", ".json")
	err := getter.GetFile(filename, url)
	if err != nil {
		fmt.Println(err)
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
	fmt.Println(stfrom, stto)
	from, err := os.Open(stfrom)
	if err != nil {
		fmt.Println(err)
	}
	defer from.Close()
	to, err := os.Open(stto)
	if err != nil {
		fmt.Println(err)
	}
	defer to.Close()
	io.Copy(from, to)
}

func tempFileName(prefix, suffix string) string {
	randBytes := make([]byte, 16)
	rand.Read(randBytes)
	return filepath.Join(os.TempDir(), prefix+hex.EncodeToString(randBytes)+suffix)
}

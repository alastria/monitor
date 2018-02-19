package services

import (
	"crypto/tls"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

	"github.com/alastria/monitor/lib"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
)

var apppath string
var err error
var log *logs.BeeLogger

var validators = regexp.MustCompile(`\| +(.*) +\| +(.*) +\| +.* +\| +enode://(.*)@(.*):(.*)\?discport.* +\|`)
var regulars = regexp.MustCompile(`\| *(.*) *\| *(.*) *\| *.* *\| *(.*=) *\| *enode://([a-z0-9]*)@([0-9\.]*):([0-9]*)\?[a-z=0]+ *\|`)
var emails = regexp.MustCompile(`[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}`)

type Nodo struct {
	Entidad    string   `json: "entidad"`
	Contactos  []string `json: "contactos"`
	Enode      string   `json: "enode"`
	IP         string   `json: "ip"`
	Port       string   `json: "port"`
	PrivateFor string   `json: "privateFor"`
}

type NodeServices struct {
	nodos []Nodo
	cert  tls.Certificate
}

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ = filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	log = logs.GetBeeLogger()
}

// By default `ibft``
func NewNodeServices(feature string) (node *NodeServices) {
	node = new(NodeServices)
	nd := node.GetValidatorDirectory("https://raw.githubusercontent.com/alastria/alastria-node/feature/" + feature + "/DIRECTORY_VALIDATOR.md")
	for cont := 0; cont < len(nd); cont++ {
		node.nodos = append(node.nodos, nd[cont])
	}
	nd = node.GetGeneralDirectory("https://raw.githubusercontent.com/alastria/alastria-node/feature/" + feature + "/DIRECTORY_REGULAR.md")
	for cont := 0; cont < len(nd); cont++ {
		node.nodos = append(node.nodos, nd[cont])
	}

	// http://www.levigross.com/2015/11/21/mutual-tls-authentication-in-go/
	// Load our TLS key pair to use for authentication
	node.cert, err = tls.LoadX509KeyPair(
		apppath+beego.AppConfig.String("TLSClientCertFile"),
		apppath+beego.AppConfig.String("TLSClientKeyFile"),
	)
	if err != nil {
		log.Error("Unable to load cert", err)
	}
	return node
}

// uri example: /v1/monitor/status
func (n *NodeServices) call(ip string, uri string) (response string, err error) {
	// https://beego.me/docs/module/httplib.md

	//log.SetFlags(log.Lshortfile)
	//beego.Trace("Get config")
	req := httplib.Get("https://" + ip + ":8443" + uri)
	//beego.Trace("TLS Config")
	req.SetTLSClientConfig(&tls.Config{
		InsecureSkipVerify: true,
		Certificates:       []tls.Certificate{n.cert},
	})
	//beego.Trace("Running")
	response, err = req.Debug(true).String()

	//beego.Info("response: %s", response, "error: %s", err)

	return
}

func (n *NodeServices) GetValidatorDirectory(url string) (nodos []Nodo) {
	// "https://raw.githubusercontent.com/alastria/alastria-node/feature/ibft/DIRECTORY_VALIDATOR.md"
	stfile, _ := lib.GetGithub(url)
	result := validators.FindAllStringSubmatch(lib.GetFile(stfile), -1)
	nodos = make([]Nodo, len(result))
	for cont := 0; cont < len(result); cont++ {
		nodos[cont].Entidad = strings.TrimSpace(result[cont][1])
		mails := emails.FindAllStringSubmatch(result[cont][2], -1)
		nodos[cont].Contactos = make([]string, len(mails))
		for cmail := 0; cmail < len(mails); cmail++ {
			nodos[cont].Contactos[cmail] = strings.TrimSpace(mails[cmail][0])
		}
		nodos[cont].Enode = strings.TrimSpace(result[cont][3])
		nodos[cont].IP = strings.TrimSpace(result[cont][4])
		nodos[cont].Port = strings.TrimSpace(result[cont][5])
	}
	return
}

func (n *NodeServices) GetGeneralDirectory(url string) (nodos []Nodo) {
	// "https://raw.githubusercontent.com/alastria/alastria-node/feature/ibft/DIRECTORY_REGULAR.md"
	stfile, _ := lib.GetGithub(url)
	result := regulars.FindAllStringSubmatch(lib.GetFile(stfile), -1)
	nodos = make([]Nodo, len(result))
	for cont := 0; cont < len(result); cont++ {
		nodos[cont].Entidad = strings.TrimSpace(result[cont][1])
		mails := emails.FindAllStringSubmatch(result[cont][2], -1)
		nodos[cont].Contactos = make([]string, len(mails))
		for cmail := 0; cmail < len(mails); cmail++ {
			nodos[cont].Contactos[cmail] = strings.TrimSpace(mails[cmail][0])
		}
		nodos[cont].PrivateFor = strings.TrimSpace(result[cont][3])
		nodos[cont].Enode = strings.TrimSpace(result[cont][4])
		nodos[cont].IP = strings.TrimSpace(result[cont][5])
		nodos[cont].Port = strings.TrimSpace(result[cont][6])
	}
	return
}

func (n *NodeServices) NodeVerify() {
	for cont := 0; cont < len(n.nodos); cont++ {
		retorno, err := n.call(n.nodos[cont].IP, "/v1/monitor/status")
		if err == nil {
			retorno, err = n.call(n.nodos[cont].IP, "/v1/node/logsjson")
		}
		log.Info(n.nodos[cont].Entidad, n.nodos[cont].IP, retorno, err)
	}
}

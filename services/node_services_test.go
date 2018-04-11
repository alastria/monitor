package services

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"testing"
	"text/template"

	"github.com/astaxie/beego"
	// "github.com/ethereum/go-ethereum/log"
	. "github.com/smartystreets/goconvey/convey"
)

var nodeService *NodeServices

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ = filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

func TestLoadValidatorDirectory(t *testing.T) {
	nodeService := NewNodeServices("ibft")
	nodos := nodeService.GetValidatorDirectory("https://raw.githubusercontent.com/alastria/alastria-node/feature/ibft/DIRECTORY_VALIDATOR.md")
	Convey("Connection: Load Validator Directory\n", t, func() {
		Convey("Result don't must be empty", func() {
			So(nodos, ShouldNotBeNil)
		})
		Convey("Result must Be bigger that 0", func() {
			So(len(nodos), ShouldBeGreaterThan, 0)
		})
		Convey("Error first must be Alastria", func() {
			So(nodos[0].Entidad, ShouldEqual, "Alastria")
		})
		Convey("Error private for must be empty", func() {
			So(nodos[0].PrivateFor, ShouldBeEmpty)
		})
	})
}

func TestLoadGeneralDirectory(t *testing.T) {
	nodeService := NewNodeServices("ibft")
	nodos := nodeService.GetValidatorDirectory("https://raw.githubusercontent.com/alastria/alastria-node/feature/ibft/DIRECTORY_VALIDATOR.md")
	Convey("Connection: Load Validator Directory\n", t, func() {
		Convey("Result don't must be empty", func() {
			So(nodos, ShouldNotBeNil)
		})
		Convey("Result must Be bigger that 0", func() {
			So(len(nodos), ShouldBeGreaterThan, 0)
		})
		Convey("Error first must be Alastria", func() {
			So(nodos[0].Entidad, ShouldEqual, "Santander")
		})
		Convey("Error private for don't must be empty", func() {
			So(nodos[0].PrivateFor, ShouldNotBeNil)
		})
	})
}

// TestCheckPermissions verify if our nodes still have
// go test -timeout 99999999s -run TestCheckPermissions
func TestCheckPermissions(t *testing.T) {
	nodeService := NewNodeServices("ibft")
	log.Info("Verificando la lista de nodos:")
	for cont := 0; cont < len(nodeService.nodos); cont++ {
		nodo := nodeService.nodos[cont]
		log.Info("Entidad: %s, Contactos: %s, Enode: %s, IP: %s",
			nodo.Entidad, nodo.Contactos, nodo.Enode, nodo.IP)
	}
	incidencias := nodeService.CheckPermission()
	log.Info("Incidencias:")
	for cont := 0; cont < len(incidencias); cont++ {
		nodo := incidencias[cont]
		log.Info("Entidad: %s, Contactos: %s, Enode: %s, IP: %s, Incidencias: %s",
			nodo.Entidad, nodo.Contactos, nodo.Enode, nodo.IP, nodo.Incidencias)
		tmpl, err := template.New("test").Parse(`
Estimado socio,

Me pongo en contacto con usted porque está registrado en nuestro directorio como responsable del nodo {{.Entidad}} con enode '{{.Enode}}'.

Durante la revisiones periódicas que realizamos a todos los nodos de nuestra red Alastria, hemos identificado las siguientes carencias:{{.Incidencias}}

Rogamos que subsane estas incidencias lo antes posible cumpliendo con las políticas establecidas por el consorcio a tal efecto.

Para obtener ayuda puede acudir:
1. https://github.com/alastria/alastria-node/wiki 
2. https://alastria.slack.com
3. https://alastriaplatform.invisionzone.com/
4. https://tree.taiga.io/project/marcossanlab-alastria-platform/

Reciba un cordial saludo,

Alastria Platform Core Team
platform@alastria.io

`)
		if err != nil {
			panic(err)
		}
		err = tmpl.Execute(os.Stdout, nodo)
		if err != nil {
			panic(err)
		}
	}
	inc, _ := json.Marshal(incidencias)
	log.Debug(string(inc))
	Convey("Checking permissions on each node\n", t, func() {
		Convey("Verifiying that the report isn't nil\n", func() {
			So(incidencias, ShouldNotBeNil)
		})
	})
}

// go test -timeout 99999999s -run TestProposeValidators
/*
Indra: 0x59d9f63451811c2c3c287be40a2206d201dc3bff *
Everis: 0xd4e6453afcfbfc8c61d53b71fddcc484e14f45e0
CaixaBank: 0xa49ddde59f9c521be81297a4c49697a5497063b3
Alastria: 0xb87dc349944cc47474775dde627a8a171fc94532
Grant Thornton: 0xc18d744dd6d06f18b79544566d01dfa881d3626c
Santander: 0x7986c3fdb87a2149975d8fa2c1d65401da6b3399
Repsol: 0xf10067b13018211e17e880030dd2c38f1cdcb97 *
*/
func TestProposeValidators(t *testing.T) {
	nodeService := NewNodeServices("ibft")
	log.Debug("Incidencias: %s", nodeService.CheckPermission())
	//log.Debug("Validadores: %s", nodeService.ListValidators())
	result := nodeService.ProposeNodes("0x986f2503001e238ade8f401ca0b0930dd9ba4563") // Change me for propose/unpropose a node
	Convey("Connection: Propose should finish completely\n", t, func() {
		Convey("Result don't must be empty", func() {
			So(result, ShouldNotBeNil)
		})
		Convey("Result must be OK", func() {
			So(result, ShouldEqual, true)
		})
	})
}

// go test -timeout 99999999s -run TestListVolunteers
func TestListVolunteers(t *testing.T) {
	nodeService := NewNodeServices("ibft")
	log.Debug("Incidencias: %s", nodeService.CheckPermission())
	volunteers := nodeService.ListVolunteers()
	var activo string
	for key := range volunteers {
		aux := volunteers[key]

		i := sort.SearchStrings(aux.Validators, aux.Coinbase)

		if i < len(aux.Validators) && aux.Validators[i] == aux.Coinbase {
			activo = "SI"
		} else {
			activo = "NO"
		}

		log.Info("%s[%s]: %s - %s", aux.Entidad, activo, aux.Coinbase, aux.Validators)
	}
	return
}

// Actualizar la versión del monitor en todos los nodos
func TestVersionUpdate(t *testing.T) {
	nodeService := NewNodeServices("ibft")
	log.Debug("Incidencias: %s", nodeService.CheckPermission())
	nodeService.VersionUpdate()
	return
}

// Actualizar los ficheros de permisionado en todos los nodos
func TestUpdate(t *testing.T) {
	nodeService := NewNodeServices("ibft")
	log.Debug("Incidencias: %s", nodeService.CheckPermission())
	nodeService.Update()
	return
}

/*
func TestCalls(t *testing.T) {
	nodeService := NewNodeServices("ibft")

	//stop := "/v1/node/stop"
	//versionupdate := "/v1/monitor/versionupdate"
	update := "/v1/node/update"
	//start := "/v1/node/start"
	var nodos []string = ["/v1/node/update"]
	var uris  []string = []
	nodos = append(nodos, "")
	nodos = append(nodos, "")
	nodos = append(nodos, "")
	nodos = append(nodos, "")
	nodos = append(nodos, "")
	nodos = append(nodos, "")
	nodos = append(nodos, "")
	nodos = append(nodos, "")

	err := nodeService.Calls(nodos, uris)

}*/

package services

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
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

// >> /home/arocha/go/bin/go test -timeout 30s github.com/alastria/monitor/services -run ^TestProposeValidators$
func TestProposeValidators(t *testing.T) {
	nodeService := NewNodeServices("ibft")
	// log.Info("Verificando la lista de nodos:")
	// for cont := 0; cont < len(nodeService.nodos); cont++ {
	// 	nodo := nodeService.nodos[cont]
	// 	log.Info("Entidad: %s, Contactos: %s, Enode: %s, IP: %s",
	// 		nodo.Entidad, nodo.Contactos, nodo.Enode, nodo.IP)
	// }
	result := nodeService.ProposeNodes()
	Convey("Connection: Propose should finish completely\n", t, func() {
		Convey("Result don't must be empty", func() {
			So(result, ShouldNotBeNil)
		})
		Convey("Result must be OK", func() {
			So(result, ShouldEqual, "true")
		})
	})
}

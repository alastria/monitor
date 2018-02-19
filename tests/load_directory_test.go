package test

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/alastria/monitor/services"
	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
)

var nodeService *services.NodeServices

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ = filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

func TestLoadValidatorDirectory(t *testing.T) {
	nodeService := services.NewNodeServices("ibft")
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
	nodeService := services.NewNodeServices("ibft")
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

func TestNodeVerify(t *testing.T) {
	nodeService := services.NewNodeServices("ibft")
	nodeService.NodeVerify()
}

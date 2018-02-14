package test

import (
	"crypto/tls"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"github.com/ethereum/go-ethereum/log"
	. "github.com/smartystreets/goconvey/convey"
)

var apppath string

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ = filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

func TestConnection(t *testing.T) {
	// https://beego.me/docs/module/httplib.md

	beego.Trace(
		apppath+beego.AppConfig.String("TLSClientCertFile"),
		apppath+beego.AppConfig.String("TLSClientKeyFile"),
	)

	// http://www.levigross.com/2015/11/21/mutual-tls-authentication-in-go/
	// Load our TLS key pair to use for authentication
	cert, err := tls.LoadX509KeyPair(
		apppath+beego.AppConfig.String("TLSClientCertFile"),
		apppath+beego.AppConfig.String("TLSClientKeyFile"),
	)
	if err != nil {
		log.Error("Unable to load cert", err)
	}

	//log.SetFlags(log.Lshortfile)
	beego.Trace("Get config")
	req := httplib.Get("https://localhost:8443/v1/monitor/status")
	beego.Trace("TLS Config")
	req.SetTLSClientConfig(&tls.Config{
		InsecureSkipVerify: true,
		Certificates:       []tls.Certificate{cert},
	})
	beego.Trace("Running")
	rsp, err := req.Debug(true).String()

	beego.Info("response: %s", rsp, "error: %s", err)

	Convey("Connection: Test monitor status\n", t, func() {
		Convey("Status don't must be empty", func() {
			So(rsp, ShouldNotBeNil)
		})
		Convey("Status must Be OK", func() {
			So(strings.Split(rsp, "\"")[3], ShouldEqual, "ok")
		})
		Convey("Error must be emtpy", func() {
			So(err, ShouldBeNil)
		})
	})

}

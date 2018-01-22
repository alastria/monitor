package main

import (
	"io/ioutil"

	gh "github.com/alastria/monitor/middleware"
	hydra "github.com/ory-am/hydra/sdk"

	// "fmt"
	// "monitor/lib"

	"github.com/alastria/monitor/lib"
	_ "github.com/alastria/monitor/routers"
	// "time"

	"crypto/tls"
	"crypto/x509"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/robfig/cron"
)

var hc *hydra.Client
var log *logs.BeeLogger

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	} else {
		beego.BeeApp.Server.TLSConfig = configureTLS()
		configureOauth2()
	}

	// Start CRON
	c := cron.New()
	c.AddFunc("0 0 30 * * *", lib.UpdateCron)
	c.Start()

	// Start REST API
	beego.Run()

}

func configureTLS() *tls.Config {
	// log := *logs.GetBeeLogger()
	log.Trace("Preparando la configuración TLS ")

	// http://www.levigross.com/2015/11/21/mutual-tls-authentication-in-go/
	// Load our TLS key pair to use for authentication
	cert, err := tls.LoadX509KeyPair(beego.AppConfig.String("TLSCertFile"), beego.AppConfig.String("TLSKeyFile"))
	if err != nil {
		log.Error("Unable to load cert", err)
	}

	// Load our CA certificate
	clientCACert, err := ioutil.ReadFile(beego.AppConfig.String("TLSCACertFile"))
	if err != nil {
		log.Error("Unable to open cert", err)
	}
	clientCertPool := x509.NewCertPool()
	clientCertPool.AppendCertsFromPEM(clientCACert)

	tlsConfig := &tls.Config{
		ClientAuth:               tls.RequireAndVerifyClientCert,
		Certificates:             []tls.Certificate{cert},
		ClientCAs:                clientCertPool,
		CipherSuites:             []uint16{tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256, tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384},
		PreferServerCipherSuites: true,
		MinVersion:               tls.VersionTLS12,
	}

	tlsConfig.BuildNameToCertificate()

	log.Trace("Finalizando la configuración TLS")

	return tlsConfig
}

func configureOauth2() {
	var err error
	// Initialize Hydra and gin-hydra
	if hc, err = hydra.Connect(
		hydra.ClientID("ftl-client"),
		hydra.ClientSecret("7Vc7VZUtrbJ9sDCQ"),
		hydra.ClusterURL("http://localhost:4444"),
		hydra.Scopes("openid", "ftl"),
	); err != nil {
		panic(err)
	}

	gh.Init(hc)

	// Use the middleware `ftl.db.user.byid`
	beego.InsertFilter("/v1/user/certid/*", beego.BeforeRouter, gh.ScopesRequired("ftl.db.user.byid"))

	beego.InsertFilter("*", beego.BeforeRouter, gh.ScopesRequired("ftl"))
}

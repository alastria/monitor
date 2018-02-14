package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	gh "github.com/alastria/monitor/middleware"
	hydra "github.com/ory-am/hydra/sdk"

	"github.com/alastria/monitor/lib"
	_ "github.com/alastria/monitor/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/robfig/cron"
)

var hc *hydra.Client
var log *logs.BeeLogger
var path string

func main() {
	var err error
	log = logs.GetBeeLogger()
	log.Trace("main is IN")
	path, err = filepath.Abs(os.Args[0])
	path = path[0:strings.LastIndex(path, "/")]

	log.Debug(path, err)
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"

		/*		beego.InsertFilter("*", beego.BeforeRouter,cors.Allow(&cors.Options{
				AllowAllOrigins: true,
				AllowMethods: []string{"GET", "DELETE", "PUT", "PATCH", "OPTIONS"},
				AllowHeaders: []string{"Origin", "Access-Control-Allow-Origin"},
				ExposeHeaders: []string{"Content-Length", "Access-Control-Allow-Origin"},
				AllowCredentials: true,
			}))*/

	} else {
		//beego.BeeApp.Server.TLSConfig = configureTLS()
		configureOauth2()
	}
	beego.BeeApp.Server.TLSConfig = configureTLS()

	// Start CRON
	c := cron.New()
	c.AddFunc("0 0 30 * * *", lib.UpdateCron)
	c.Start()

	// Start REST API
	beego.Run()
	log.Trace("main is OUT!")
}

func configureTLS() *tls.Config {
	// log := *logs.GetBeeLogger()
	log.Trace("Preparando la configuración TLS ")

	// http://www.levigross.com/2015/11/21/mutual-tls-authentication-in-go/
	// Load our TLS key pair to use for authentication
	cert, err := tls.LoadX509KeyPair(
		path+beego.AppConfig.String("TLSCertFile"),
		path+beego.AppConfig.String("TLSKeyFile"),
	)
	if err != nil {
		log.Error("Unable to load cert", err)
	}

	// Load our CA certificate
	clientCACert, err := ioutil.ReadFile(path + beego.AppConfig.String("TLSCACertFile"))
	if err != nil {
		log.Error("Unable to open cert", err)
	}
	clientCertPool := x509.NewCertPool()
	clientCertPool.AppendCertsFromPEM(clientCACert)

	tlsConfig := &tls.Config{
		ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: []tls.Certificate{cert},
		ClientCAs:    clientCertPool,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
		},
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
		hydra.ClientID(path+beego.AppConfig.String("OAUTH2_CLIENT_ID")),
		hydra.ClientSecret(path+beego.AppConfig.String("OAUTH2_CLIENT_SECRET")),
		hydra.ClusterURL(path+beego.AppConfig.String("OAUTH2_CLUSTER_URL")),
		hydra.Scopes(path+beego.AppConfig.String("OAUTH2_SCOPES")),
	); err != nil {
		panic(err)
	}

	gh.Init(hc)

	// Use the middleware `ftl.db.user.byid`
	//beego.InsertFilter("/v1/user/certid/*", beego.BeforeRouter, gh.ScopesRequired("alastria.db.user.byid"))

	//beego.InsertFilter("*", beego.BeforeRouter, gh.ScopesRequired("alastria"))
}

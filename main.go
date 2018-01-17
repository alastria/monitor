package main

import (

	// "fmt"
	// "monitor/lib"

	"monitor/lib"
	_ "monitor/routers"
	// "time"

	"github.com/astaxie/beego"
	"github.com/robfig/cron"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	// Start CRON
	c := cron.New()
	c.AddFunc("0 0 30 * * *", lib.UpdateCron)
	c.Start()

	// Start REST API
	beego.Run()

}

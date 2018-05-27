package main

import (
	"go-webapp/config"
	"go-webapp/module/server"
	"go-webapp/routes"
	"runtime"
	log "github.com/sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

func init() {
	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

func main() {
	// Set maximum number of CPUs that can be executing simultaneously with the number of logical CPUs usable by the current process
	runtime.GOMAXPROCS(runtime.NumCPU())
	if config.GetEnv().DEBUG {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	router := routes.InitRouter()
	server.Run(router)
}
package routes

import (
	"go-webapp/module/debug"
	"go-webapp/module/server"

	"github.com/gin-gonic/gin"
)

func registerAPIRouter(router *gin.Engine) {
	api := router.Group("/api")
	api.GET("/api-test", server.TestApi)
	api.GET("/api-test-postgres", server.GetProducts)
	//router.HandleFunc("/resources", server.GetProducts).Methods("GET")

	debugger := router.Group("/api/debug")
	{
		debugger.GET("/pprof/", debug.IndexHandler())
		debugger.GET("/pprof/heap", debug.HeapHandler())
		debugger.GET("/pprof/goroutine", debug.GoroutineHandler())
		debugger.GET("/pprof/block", debug.BlockHandler())
		debugger.GET("/pprof/threadcreate", debug.ThreadCreateHandler())
		debugger.GET("/pprof/cmdline", debug.CmdlineHandler())
		debugger.GET("/pprof/profile", debug.ProfileHandler())
		debugger.GET("/pprof/symbol", debug.SymbolHandler())
		debugger.POST("/pprof/symbol", debug.SymbolHandler())
		debugger.GET("/pprof/trace", debug.TraceHandler())
	}

	router.GET("/version", server.Version)
}

//type Route struct {
//	Name        string
//	Method      string
//	Pattern     string
//	HandlerFunc http.HandlerFunc
//}
//
//type Routes []Route
//
//var routes = Routes{
//	Route{
//		"Products",
//		"GET",
//		"/api-test-postgres",
//		server.GetProducts(),
//	},
//}
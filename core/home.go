package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HomeArgs struct {
	ApiRouteDict gin.H
	Title string	
}

var HomeArgContext HomeArgs


func RegisterCoreRoutes(router *gin.RouterGroup, engine *gin.Engine) {
	engine.GET("/", getHome)
	router.GET("/", getHome)
}



func getHome(ctxt *gin.Context){
	// dict := gin.H{
	// 	"title": "Budget Application",
	// }

	dict := make(map[string]interface{})
	dict["args"] = HomeArgContext
	fmt.Printf("Setting variables to %s\n", dict)
	ctxt.HTML(http.StatusOK, "home.html", dict)
}
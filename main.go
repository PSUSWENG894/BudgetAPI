package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	msg := "Hello World!"
	fmt.Printf(msg)

	r := gin.Default()
	r.GET("/", func(ctxt *gin.Context){
		ctxt.JSON(200, gin.H{"message": msg},)
	})
	r.Run()
}
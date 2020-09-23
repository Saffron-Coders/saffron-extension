package app

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

func StartAPP() {
	URLMapping()
	if err := router.Run(":9090"); err != nil {
		fmt.Println("Couldn't start the server at 9090 port!")
		log.Println(err)
	}
}

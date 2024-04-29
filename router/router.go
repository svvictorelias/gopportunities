package router

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func Initialize() {

	router := gin.Default()

	initializeRoutes(router)

	fmt.Print("\n\n", os.Getenv("XX"), "aaaaaaaaaaa", "\n\n")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run("0.0.0.0:" + port)
}

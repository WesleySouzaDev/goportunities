package router

import "github.com/gin-gonic/gin"

func Initializer() {
	// Initialeze Router
	router := gin.Default()
	// Initialize the routes
	InitializerRouter(router)
	// Run the server
	router.Run() // listens on 0.0.0.0:8080 by default
}

//go:generate swag init
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// Swagger docs
	"github.com/mouse-blink/petgo/docs"
)

// User represents a user
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// @title           Minimal Go Swagger Example
// @version         1.0
// @description     Simple example of Swagger with Go and Gin.
// @BasePath        /api/v1

func main() {
	// Dynamically set Swagger host
	nodeIP := os.Getenv("NODE_IP")
	nodePort := os.Getenv("NODE_PORT")
	if nodeIP != "" && nodePort != "" {
		docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", nodeIP, nodePort)
	} else {
		log.Println("NODE_IP or NODE_PORT not set, using default host: localhost:8080")
		docs.SwaggerInfo.Host = "localhost:8080"
	}

	r := gin.Default()

	// API routes
	v1 := r.Group("/api/v1")
	{
		v1.GET("/user", getUser)
		v1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// K8s health probes
	r.GET("/healthz", liveness)
	r.GET("/readyz", readiness)
	r.GET("/", func(c *gin.Context) {
		fmt.Fprintln(c.Writer, "Hello, World!")
	})
	// Start server
	r.Run(":8080")
}

// getUser godoc
// @Summary      Get a user
// @Description  Get sample user
// @Tags         user
// @Produce      json
// @Success      200  {object}  User
// @Router       /user [get]
func getUser(c *gin.Context) {
	user := User{ID: 1, Name: "John Doe"}
	c.JSON(http.StatusOK, user)
}

// Liveness probe
func liveness(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "alive"})
}

// Readiness probe
func readiness(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ready"})
}

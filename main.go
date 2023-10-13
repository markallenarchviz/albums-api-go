package main

import (
	"github.com/gin-gonic/gin"
	"github.com/marchallenarchviz/albums-api-go/src/controllers/routes"
)

func main() {
	router := gin.Default()
	routes.AlbumsRoutes(&router.RouterGroup)
	router.Run("localhost:8080")
}

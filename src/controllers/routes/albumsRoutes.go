package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/marchallenarchviz/albums-api-go/src/controllers"
)

func AlbumsRoutes(r *gin.RouterGroup) {
	r.GET("/albums", controllers.GetAlbums)
	r.POST("/albums", controllers.PostAlbums)
	r.GET("/albums/:id", controllers.GetAlbumByID)
}

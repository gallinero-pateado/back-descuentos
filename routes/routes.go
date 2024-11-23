package routes

import (
	"descuentos/handlers"

	auth "descuentos/auth"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	protected := r.Group("/").Use(auth.AuthMiddleware)
	{
		protected.GET("/descuentos", handlers.MostrarDescuentos)
	}
}

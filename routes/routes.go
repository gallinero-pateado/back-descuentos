package routes

import (
	"descuentos/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/descuentos", handlers.MostrarDescuentos)
}

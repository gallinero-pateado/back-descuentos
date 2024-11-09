package main

import (
	"descuentos/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Endpoint para la sección de descuentos
	r.GET("/descuentos", handlers.MostrarDescuentos)

	r.Run(":8080") // Ejecuta el servidor en el puerto 8080
}

package main

import (
	"log"

	"descuentos/config"
	"descuentos/routes"
	"descuentos/services"

	"github.com/gin-gonic/gin"
)

func main() {

	services.EjecutarScraping()

	// Crear el enrutador de Gin
	r := gin.Default()

	// Configuración de CORS
	r.Use(config.CORSConfig())

	// Registrar rutas
	routes.RegisterRoutes(r)

	// Ejecutar el servidor en el puerto 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}

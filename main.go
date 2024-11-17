package main

import (
	"descuentos/handlers"

	"descuentos/scraping"

	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Ejecutar el scraping para Wendy's
	if err := scraping.ScrapingWendys("data/wendys.json"); err != nil {
		log.Fatalf("Error al realizar el scraping a la página de Wendy's: %v", err)
	}
	// Ejecutar el scraping para Burger King
	if err := scraping.ScrapingBurger("data/burgerking.json"); err != nil {
		log.Fatalf("Error al realizar el scraping a la página de Burger King: %v", err)
	}

	// Ejecutar el scraping para Little Caesars
	if err := scraping.ScrapingLittleCaesars("data/little_caesars.json"); err != nil {
		log.Fatalf("Error al realizar el scraping a la página de Little Caesars: %v", err)
	}

	// Ejecutar el scraping para Oxxo
	if err := scraping.ScrapingOxxo("data/oxxo.json"); err != nil {
		log.Fatalf("Error al realizar el scraping a la página de Oxxo: %v", err)
	}

	r := gin.Default()

	// Endpoint para la sección de descuentos
	r.GET("/descuentos", handlers.MostrarDescuentos)

	r.Run(":8080") // Ejecuta el servidor en el puerto 8080
}

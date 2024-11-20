package handlers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// Estructura unificada para los productos
type Product struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Category      string `json:"category"`
	Description   string `json:"description"`
	Price         string `json:"price"`
	PreviousPrice string `json:"previous_price,omitempty"`
	Image         string `json:"image"`
	Logo          string `json:"logo"`
	Type          string `json:"type"`
	Url           string `json:"url"`
}

// MostrarDescuentos procesa y unifica los datos de los diferentes JSON
func MostrarDescuentos(c *gin.Context) {
	// Rutas de los archivos JSON
	files := []string{
		"data/wendys.json",
		"data/burgerking.json",
		"data/little_caesars.json",
		"data/oxxo.json",
	}

	var allProducts []Product

	// Leer y procesar cada archivo JSON
	for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al leer el archivo", "file": file, "details": err.Error()})
			return
		}

		// Decodificar los datos en la estructura unificada
		var products []Product
		if err := json.Unmarshal(data, &products); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al parsear el archivo", "file": file, "details": err.Error()})
			return
		}

		allProducts = append(allProducts, products...)
	}

	// Responder con todos los productos unificados
	c.JSON(http.StatusOK, allProducts)
}

package handlers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"descuentos/models"
)

// MostrarDescuentos procesa y unifica los datos de los diferentes JSON
func MostrarDescuentos(c *gin.Context) {
	// Obtener el UID del contexto
	uid, exists := c.Get("uid")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No se encontró el UID"})
		return
	}

	// Rutas de los archivos JSON
	files := []string{
		"services/data/wendys.json",
		"services/data/burgerking.json",
		"services/data/little_caesars.json",
		"services/data/oxxo.json",
	}

	var allProducts []models.Product

	// Leer y procesar cada archivo JSON
	for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al leer el archivo", "file": file, "details": err.Error()})
			return
		}

		// Decodificar los datos en la estructura unificada
		var products []models.Product
		if err := json.Unmarshal(data, &products); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al parsear el archivo", "file": file, "details": err.Error()})
			return
		}

		allProducts = append(allProducts, products...)
	}

	// Responder con todos los productos unificados
	c.JSON(http.StatusOK, gin.H{
		"message":  "Acceso a descuentos",
		"user_id":  uid,
		"products": allProducts,
	})
}

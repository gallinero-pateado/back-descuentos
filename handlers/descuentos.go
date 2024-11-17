package handlers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type DescuentoWS struct {
	ID             int    `json:"id"`
	Titulo         string `json:"name"`
	Categoria      string `json:"category"`
	Descripcion    string `json:"description"`
	Precio         string `json:"price"`
	PrecioAnterior string `json:"previous_price"`
	Descuento      string `json:"discount"`
	Imagen         string `json:"image"`
	Logo           string `json:"logo"`
}

type DescuentoBK struct {
	ID          int    `json:"id"`
	Titulo      string `json:"name"`
	Categoria   string `json:"category"`
	Descripcion string `json:"description"`
	Precio      string `json:"price"`
	Imagen      string `json:"image"`
	Logo        string `json:"logo"`
}

type DescuentoLC struct {
	ID          int    `json:"id"`
	Titulo      string `json:"name"`
	Categoria   string `json:"category"`
	Descripcion string `json:"description"`
	Precio      string `json:"price"`
	Descuento   string `json:"discount"`
	Imagen      string `json:"image"`
	Logo        string `json:"logo"`
}

type DescuentoOx struct {
	ID          int    `json:"id"`
	Titulo      string `json:"name"`
	Categoria   string `json:"category"`
	Descripcion string `json:"description"`
	Precio      string `json:"price"`
	Imagen      string `json:"image"`
	Logo        string `json:"logo"`
}

func MostrarDescuentos(c *gin.Context) {
	// Leer el archivo JSON
	descuentoWSData, err := os.ReadFile("data/wendys.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo leer el archivo de descuentos de Wendys", "details": err.Error()})
		return
	}

	descuentoBKData, err := os.ReadFile("data/burgerking.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo leer el archivo de descuentos de Burger King", "details": err.Error()})
		return
	}

	descuentoLCData, err := os.ReadFile("data/little_caesars.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo leer el archivo de descuentos de Little Caesars", "details": err.Error()})
		return
	}

	descuentoOxData, err := os.ReadFile("data/oxxo.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo leer el archivo de descuentos de Oxxo", "details": err.Error()})
		return
	}

	// Convertir el archivo JSON a un slice de Productos
	var descuentosWS []DescuentoWS
	if err := json.Unmarshal(descuentoWSData, &descuentosWS); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al parsear el archivo de descuentos"})
		return
	}

	var descuentosBK []DescuentoBK
	if err := json.Unmarshal(descuentoBKData, &descuentosBK); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al parsear el archivo de descuentos"})
		return
	}

	var descuentosLC []DescuentoWS
	if err := json.Unmarshal(descuentoLCData, &descuentosLC); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al parsear el archivo de descuentos"})
		return
	}

	var descuentosOx []DescuentoBK
	if err := json.Unmarshal(descuentoOxData, &descuentosOx); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al parsear el archivo de descuentos"})
		return
	}

	// Enviar los descuentos como respuesta
	c.JSON(http.StatusOK, gin.H{
		"message":                  "Descuentos disponibles",
		"Descuentos Wendys":        descuentosWS,
		"Descuentos BurgerKing":    descuentosBK,
		"Descuentos LittleCaesars": descuentosLC,
		"Descuentos Oxxo":          descuentosOx,
	})
}

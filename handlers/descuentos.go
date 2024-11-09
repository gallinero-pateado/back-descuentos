package handlers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// Estructura para almacenar los descuentos del archivo JSON wendys
type DescuentoWS struct {
	ID             int    `json:"ID"`
	Titulo         string `json:"Titulo"`
	Categoria      string `json:"Categoria"`
	Descripcion    string `json:"Descripcion"`
	Precio         string `json:"Precio"`
	PrecioAnterior string `json:"Precio Anterior"`
	Descuento      string `json:"Descuento"`
	Imagen         string `json:"imagen"`
}

// Estructura para almacenar los descuentos del archivo JSON burger king
type DescuentoBK struct {
	ID          int    `json:"ID"`
	Titulo      string `json:"TItulo"`
	Liked       bool   `json:"liked"`
	Categoria   string `json:"Categoria"`
	Imagen      string `json:"Imagen"`
	Descripcion string `json:"Descipcion"`
	Precio      string `json:"Precio"`
	Condiciones string `json:"Condiciones"`
}

func MostrarDescuentos(c *gin.Context) {
	// Leer el archivo JSON
	descuentoWSData, err := os.ReadFile("data/descuentos_wendys.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo leer el archivo de descuentos", "details": err.Error()})
		return
	}

	// Leer el archivo JSON
	descuentoBKData, err := os.ReadFile("data/descuentos_bk.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo leer el archivo de descuentos", "details": err.Error()})
		return
	}

	// Convertir el archivo JSON a un slice de Productos
	var descuentosWS []DescuentoWS
	if err := json.Unmarshal(descuentoWSData, &descuentosWS); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al parsear el archivo de descuentos"})
		return
	}

	// Convertir el archivo JSON a un slice de Productos
	var descuentosBK []DescuentoBK
	if err := json.Unmarshal(descuentoBKData, &descuentosBK); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al parsear el archivo de descuentos"})
		return
	}

	// Enviar los descuentos como respuesta
	c.JSON(http.StatusOK, gin.H{
		"message":      "Descuentos disponibles",
		"descuentosWS": descuentosWS,
		"descuentosBK": descuentosBK,
	})
}

package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/PuerkitoBio/goquery"
)

// Estructura para almacenar los detalles del producto
type Producto struct {
	Titulo         string `json:"titulo"`
	Descuento      string `json:"descuento"`
	Detalles       string `json:"detalles"`
	Disponibilidad string `json:"disponibilidad"`
}

func obtenerProductos(c *gin.Context) {
	url := "https://descuentosrata.com/guata"
	// Realiza la solicitud HTTP
	resp, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al cargar la página"})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al cargar la página, código: " + resp.Status})
		return
	}

	// Analiza el HTML usando goquery
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al analizar la página"})
		return
	}

	var productos []Producto

	// Busca todos los productos en la página
	doc.Find("div.my-2.col-md-6.col-lg-4.col-12").Each(func(i int, s *goquery.Selection) {
		// Extrae el título del producto
		titulo := strings.TrimSpace(s.Find("div.card-body.d-flex").Text())
		if titulo == "" {
			titulo = "No disponible"
		}

		// Extrae el descuento del producto
		descuento := strings.TrimSpace(s.Find("span.font-weight-normal").Text())
		if descuento == "" {
			descuento = "No disponible"
		}

		// Extrae los detalles adicionales
		detalles := strings.TrimSpace(s.Find("span.badge.d-block-inline.ml-1.badge-dark.badge-pill").Text())
		if detalles == "" {
			detalles = "No disponible"
		}

		// Extrae la disponibilidad del producto
		disponibilidad := strings.TrimSpace(s.Find("span.guata-disp").Text())
		if disponibilidad == "" {
			disponibilidad = "No disponible"
		}

		// Agrega los datos del producto a la lista
		productos = append(productos, Producto{
			Titulo:         titulo,
			Descuento:      descuento,
			Detalles:       detalles,
			Disponibilidad: disponibilidad,
		})
	})

	// Devuelve el array de productos como JSON
	c.JSON(http.StatusOK, productos)
}

func main() {
	r := gin.Default()
	r.GET("/", obtenerProductos)
	r.Run(":8080")
}

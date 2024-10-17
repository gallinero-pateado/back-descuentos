package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

// Producto estructura para almacenar los detalles del producto
type Producto struct {
	Categoria      string `json:"categoria"`
	Titulo         string `json:"titulo"`
	Descripcion    string `json:"descripcion"`
	Precio         string `json:"precio"`
	PrecioAnterior string `json:"precio_anterior"`
	Descuento      string `json:"descuento"`
	Imagen         string `json:"imagen"`
}

// ObtenerProductos realiza el scraping de la página y devuelve los productos
func ObtenerProductos(c *gin.Context) {
	url := "https://www.wendys.cl/pedir"
	res, err := http.Get(url)
	if err != nil || res.StatusCode != 200 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error al cargar la página, código: %d", res.StatusCode)})
		return
	}
	defer res.Body.Close()

	// Carga el documento HTML
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al analizar el contenido HTML"})
		return
	}

	var productos []Producto

	// Busca todos los productos en la página
	doc.Find("div.product-card").Each(func(i int, s *goquery.Selection) {
		titulo := strings.TrimSpace(s.Find("span.line-clamp-2").Text()) // Extrae el título del producto
		if titulo == "" {
			titulo = "No disponible"
		}

		// Extrae la descripción del producto
		descripcion := strings.TrimSpace(s.Find("p.text-xs").Text())
		if descripcion == "" {
			descripcion = "No disponible"
		}

		// Extrae el precio y separa los precios
		precioTexto := strings.TrimSpace(s.Find("div.flex.gap-x-2.text-sm.flex-row").Text())
		precios := strings.Split(precioTexto, "$")

		var precio string
		var precioAnterior string

		if len(precios) > 2 { // Si hay dos precios
			precio = "$" + strings.TrimSpace(precios[1])         // Primer precio
			precioAnterior = "$" + strings.TrimSpace(precios[2]) // Segundo precio
		} else if len(precios) == 2 { // Si hay solo un precio
			precio = "$" + strings.TrimSpace(precios[1]) // Precio único
			precioAnterior = "No disponible"             // No hay precio anterior
		} else {
			precio = "No disponible" // Si no se encontró precio
			precioAnterior = "No disponible"
		}

		// Ajustar el selector para el descuento
		descuento := strings.TrimSpace(s.Find("span").Text())

		// Verificar si el contenido del descuento contiene el patrón HTML y asignar "No disponible" si es así
		if strings.Contains(descuento, `\u003Cimg alt=\`) {
			descuento = "No disponible"
		}

		// Extrae la imagen del producto
		imagen, exists := s.Find("img").Attr("src") // Obtener el atributo src de la imagen
		if !exists {
			imagen = "No disponible" // Si no se encuentra imagen
		}
		if imagen == "" {
			imagen = "No disponible" // Si no se encuentra imagen
		}

		productos = append(productos, Producto{
			Categoria:      "Comida Chatarra",
			Titulo:         titulo,
			Descripcion:    descripcion,
			Precio:         precio,
			PrecioAnterior: precioAnterior,
			Descuento:      descuento,
			Imagen:         imagen,
		})
	})

	// Devuelve los productos como JSON
	c.JSON(http.StatusOK, productos)
}

func main() {
	r := gin.Default()
	r.GET("/productos", ObtenerProductos)
	r.Run(":8080")
}

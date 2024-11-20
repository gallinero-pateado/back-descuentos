package scraping

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Producto estructura para almacenar los detalles del producto
type Producto struct {
	ID             int    `json:"id"`
	Titulo         string `json:"name"`
	Categoria      string `json:"category"`
	Descripcion    string `json:"description"`
	Precio         string `json:"price"`
	PrecioAnterior string `json:"previous_price"`
	Descuento      string `json:"discount"`
	Imagen         string `json:"image"`
	Logo           string `json:"logo"`
	Type           string `json:"type"`
	Url            string `json:"url"`
}

// Scrape realiza el scraping y devuelve los productos
func ScrapingWendys(filename string) error {
	url := "https://www.wendys.cl/pedir"
	res, err := http.Get(url)
	if err != nil || res.StatusCode != 200 {
		return fmt.Errorf("error al cargar la página, código: %d", res.StatusCode)
	}
	defer res.Body.Close()

	// Carga el documento HTML
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return fmt.Errorf("error al analizar el contenido HTML: %v", err)
	}

	// Obtener logo
	logo, exists := doc.Find("div._3FqAWjFlHSfPln4gH8Ox5B img").Attr("src")
	if !exists || logo == "" {
		logo = "No disponible"
	}

	var productos []Producto

	// Extraer datos de los productos
	doc.Find("div.product-card").Each(func(i int, s *goquery.Selection) {
		titulo := strings.TrimSpace(s.Find("span.line-clamp-2").Text())
		if titulo == "" {
			titulo = "No disponible"
		}

		descripcion := strings.TrimSpace(s.Find("p.text-xs").Text())
		if descripcion == "" {
			descripcion = "No disponible"
		}

		precioTexto := strings.TrimSpace(s.Find("div.flex.gap-x-2.text-sm.flex-row").Text())
		precios := strings.Split(precioTexto, "$")

		var precio, precioAnterior string
		if len(precios) > 2 {
			precio = "$" + strings.TrimSpace(precios[1])
			precioAnterior = "$" + strings.TrimSpace(precios[2])
		} else if len(precios) == 2 {
			precio = "$" + strings.TrimSpace(precios[1])
			precioAnterior = "No disponible"
		} else {
			precio = "No disponible"
			precioAnterior = "No disponible"
		}

		descuento := strings.TrimSpace(s.Find("span").Text())
		if strings.Contains(descuento, `\u003Cimg alt=\`) {
			descuento = "No disponible"
		}

		imagen, exists := s.Find("img").Attr("src")
		if !exists || imagen == "" {
			imagen = "No disponible"
		}

		productos = append(productos, Producto{
			ID:             i + 1,
			Titulo:         titulo,
			Categoria:      "wendys",
			Descripcion:    descripcion,
			Precio:         precio,
			PrecioAnterior: precioAnterior,
			Descuento:      descuento,
			Imagen:         imagen,
			Logo:           logo,
			Type:           "producto",
			Url:            url, // Asignar la URL aquí
		})
	})

	return SaveWendysToJSON(filename, productos)
}

// SaveToJSON guarda los productos en un archivo JSON
func SaveWendysToJSON(filename string, data interface{}) error {
	err := os.MkdirAll("data", os.ModePerm)
	if err != nil {
		return fmt.Errorf("error al crear la carpeta 'data': %v", err)
	}

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error al crear el archivo: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}

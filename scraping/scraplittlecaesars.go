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
type LittleCaesarsDescuento struct {
	ID          int    `json:"id"`
	Titulo      string `json:"name"`
	Categoria   string `json:"category"`
	Descripcion string `json:"description"`
	Precio      string `json:"price"`
	Descuento   string `json:"discount"`
	Imagen      string `json:"image"`
}

// Scrape realiza el scraping y devuelve los productos
func ScrapingLittleCaesars(filename string) error {
	url := "https://cl.littlecaesars.com/es-cl/menu/"
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

	var descuentos []LittleCaesarsDescuento

	// Extraer datos de los productos
	doc.Find("div.css-1x2zmgq").Each(func(i int, s *goquery.Selection) {
		titulo := strings.TrimSpace(s.Find("h2.css-1l246ro").Text())
		if titulo == "" {
			titulo = "No disponible"
		}

		descripcion := strings.TrimSpace(s.Find("p.css-vurnku").Text())
		if descripcion == "" {
			descripcion = "No disponible"
		}
		precio := strings.TrimSpace(s.Find("div.css-15n7wyn").Text())
		if precio == "" {
			precio = "No disponible"
		}

		descuento := "No disponible"

		imagen, exists := s.Find("img").Attr("src")
		if !exists || imagen == "" {
			imagen = "No disponible"
		}

		descuentos = append(descuentos, LittleCaesarsDescuento{
			ID:          i + 1,
			Titulo:      titulo,
			Categoria:   "Comida Chatarra",
			Descripcion: descripcion,
			Precio:      precio,
			Descuento:   descuento,
			Imagen:      imagen,
		})
	})

	return saveLittleToJSON(filename, descuentos)
}

// SaveToJSON guarda los productos en un archivo JSON
func saveLittleToJSON(filename string, data interface{}) error {
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

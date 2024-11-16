package scraping

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

type Discount struct {
	ID          int    `json:"id"`
	Titulo      string `json:"name"`
	Categoria   string `json:"category"`
	Descripcion string `json:"description"`
	Precio      string `json:"price"`
	Imagen      string `json:"image"`
}

type logoOxxo struct {
	Logo string `json:"logo"`
}

// ScrapingOxxo realiza el scraping a OXXO
func ScrapingOxxo(filename string) error {
	url := "https://oxxo.cl/promociones"
	res, err := http.Get(url)
	if err != nil || res.StatusCode != 200 {
		return fmt.Errorf("error al cargar la página, código: %d", res.StatusCode)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return fmt.Errorf("error al analizar el contenido HTML: %v", err)
	}

	// Obtener logo
	logo, exists := doc.Find("a.center-center img").Attr("src")
	if !exists || logo == "" {
		logo = "No disponible"
	}

	var discounts []Discount

	doc.Find("div.col-sm-4 img.img-fluid").Each(func(i int, s *goquery.Selection) {
		imagen, exists := s.Attr("src")
		if !exists || imagen == "" {
			imagen = "No disponible"
		}

		precio := "Cupón"

		discounts = append(discounts, Discount{
			ID:          i + 1,
			Titulo:      "No disponible",
			Categoria:   "Saludable",
			Descripcion: "No disponible",
			Precio:      precio,
			Imagen:      imagen,
		})
	})

	// Crear la estructura para el JSON final, agregando el logo al principio
	finalData := []interface{}{
		logoOxxo{Logo: logo},
	}

	// Añadir los cupones a los datos finales
	for _, cupon := range discounts {
		finalData = append(finalData, cupon)
	}

	// Guardar los datos en JSON
	return saveDiscountToJSON(filename, finalData)
}

// saveDiscountToJSON guarda los datos en un archivo JSON
func saveDiscountToJSON(filename string, data interface{}) error {
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

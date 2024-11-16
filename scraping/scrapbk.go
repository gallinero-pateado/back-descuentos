package scraping

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Cupon estructura de los datos del cupón
type Cupon struct {
	ID          int    `json:"id"`
	Titulo      string `json:"name"`
	Categoria   string `json:"category"`
	Descripcion string `json:"description"`
	Precio      string `json:"price"`
	Imagen      string `json:"image"`
	Logo        string `json:"logo,omitempty"`
}

// ScrapingLogo representa el logo en la estructura JSON
type ScrapingLogo struct {
	Logo string `json:"logo"`
}

// Scrap realiza el scraping a Burger King
func ScrapingBurger(filename string) error {
	url := "https://www.burgerking.cl/cupones/"
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
	logo, exists := doc.Find("img.header__brandLogo").Attr("src")
	if !exists || logo == "" {
		logo = "No disponible"
	}

	var cupones []Cupon

	// Iteramos sobre los descuentos
	doc.Find("button.card-tab").Each(func(i int, s *goquery.Selection) {
		titulo := strings.TrimSpace(s.Find("h6.coupon-name.mb-1").Text())
		if titulo == "" {
			titulo = "No disponible"
		}

		descripcion := strings.TrimSpace(s.Find("p.coupon-description.mb-0").Text())
		if descripcion == "" {
			descripcion = "No disponible"
		}

		imagen, _ := s.Find("img").Attr("src")
		if imagen == "" {
			imagen = "No disponible"
		}

		cupones = append(cupones, Cupon{
			ID:          i + 1,
			Titulo:      titulo,
			Categoria:   "Saludable",
			Descripcion: descripcion,
			Precio:      "Sin Precio",
			Imagen:      imagen,
			Logo:        "",
		})
	})

	// Crear la estructura para el JSON final, agregando el logo al principio
	finalData := []interface{}{
		ScrapingLogo{Logo: logo},
	}

	// Añadir los cupones a los datos finales
	for _, cupon := range cupones {
		finalData = append(finalData, cupon)
	}

	// Guardar los datos en JSON
	return saveToJSON(filename, finalData)
}

// saveToJSON guarda los datos en un archivo JSON
func saveToJSON(filename string, data interface{}) error {
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

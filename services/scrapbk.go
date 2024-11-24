package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// estructura de los datos unificados del producto
type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Price       string `json:"price"`
	Image       string `json:"image"`
	Logo        string `json:"logo"`
	Type        string `json:"type"`
	Url         string `json:"url"`
}

// ScrapingBurger realiza el scraping a Burger King
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

	logo, exists := doc.Find("img.header__brandLogo").Attr("src")
	if !exists || logo == "" {
		logo = "No disponible"
	}

	var products []Product

	// Extraer los productos
	doc.Find("button.card-tab").Each(func(i int, s *goquery.Selection) {
		name := strings.TrimSpace(s.Find("h6.coupon-name.mb-1").Text())
		if name == "" {
			name = "No disponible"
		}

		description := strings.TrimSpace(s.Find("p.coupon-description.mb-0").Text())
		if description == "" || isInvalidDescription(description) {
			description = "No disponible"
		}

		price := "Cupón"

		image, _ := s.Find("img").Attr("src")
		if image == "" {
			image = "No disponible"
		}

		// Agregar el producto a la lista
		products = append(products, Product{
			ID:          i + 1,
			Name:        name,
			Category:    "Burger King",
			Description: description,
			Price:       price,
			Image:       image,
			Logo:        logo,
			Type:        "Cupon",
			Url:         url,
		})
	})

	return saveToJSON(filename, products)
}

// isInvalidDescription verifica si una descripción es inválida
func isInvalidDescription(description string) bool {
	// Detectar patrones comunes que indiquen derechos de autor u otros textos no deseados
	lowerDesc := strings.ToLower(description)
	return strings.Contains(lowerDesc, "burger king") || strings.Contains(lowerDesc, "derechos reservados")
}

// saveToJSON guarda los datos en un archivo JSON
func saveToJSON(filename string, data interface{}) error {
	err := os.MkdirAll("services/data", os.ModePerm)
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

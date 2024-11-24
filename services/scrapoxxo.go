package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

type Discount struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Categoria   string `json:"category"`
	Description string `json:"description"`
	Price       string `json:"price"`
	Image       string `json:"image"`
	Logo        string `json:"logo"`
	Type        string `json:"type"`
	Url         string `json:"url"`
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
			Name:        "Promoción de Oxxo",
			Categoria:   "Oxxo",
			Description: "Más información en la página oficial.",
			Price:       precio,
			Image:       imagen,
			Logo:        logo,
			Type:        "Cupon",
			Url:         url,
		})
	})

	return saveDiscountToJSON(filename, discounts)
}

// saveDiscountToJSON guarda los datos en un archivo JSON
func saveDiscountToJSON(filename string, data interface{}) error {
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

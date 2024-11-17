package scraping

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Cupon struct {
	ID          int    `json:"id"`
	Titulo      string `json:"name"`
	Categoria   string `json:"category"`
	Descripcion string `json:"description"`
	Precio      string `json:"price"`
	Imagen      string `json:"image"`
	Logo        string `json:"logo"`
}

// realiza el scraping a Burger King
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

	var cupones []Cupon

	doc.Find("button.card-tab").Each(func(i int, s *goquery.Selection) {
		titulo := strings.TrimSpace(s.Find("h6.coupon-name.mb-1").Text())
		if titulo == "" {
			titulo = "No disponible"
		}

		descripcion := strings.TrimSpace(s.Find("p.coupon-description.mb-0").Text())
		if descripcion == "" {
			descripcion = "No disponible"
		}

		precio := "Cupón"

		imagen, _ := s.Find("img").Attr("src")
		if imagen == "" {
			imagen = "No disponible"
		}

		cupones = append(cupones, Cupon{
			ID:          i + 1,
			Titulo:      titulo,
			Categoria:   "Burger King",
			Descripcion: descripcion,
			Precio:      precio,
			Imagen:      imagen,
			Logo:        logo,
		})
	})

	return saveToJSON(filename, cupones)
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

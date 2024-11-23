package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Product estructura de los datos unificados del producto
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

	// Extraer el logo
	logo, exists := doc.Find("img.header__brandLogo").Attr("src")
	if !exists || logo == "" {
		logo = "No disponible"
	}

	var products []Product

	// Extraer los cupones
	doc.Find("button.card-tab").Each(func(i int, s *goquery.Selection) {
		// Extraer título
		name := strings.TrimSpace(s.Find("h6.coupon-name.mb-1").Text())
		if name == "" {
			name = "No disponible"
		}

		// Extraer descripción
		description := strings.TrimSpace(s.Find("p.coupon-description.mb-0").Text())
		if description == "" {
			description = "No disponible"
		} else {
			// Eliminar el texto repetitivo de derechos de autor
			description = removeCopyrightMessage(description)
		}

		// Precio actual (BK no muestra precio en los cupones)
		price := "Cupón"

		// Imagen del producto
		image, _ := s.Find("img").Attr("src")
		if image == "" {
			image = "No disponible"
		}

		// Agregar el producto a la lista
		products = append(products, Product{
			ID:          i + 1,
			Name:        name,
			Category:    "burgerking",
			Description: description,
			Price:       price,
			Image:       image,
			Logo:        logo,
			Type:        "Cupon",
			Url:         url,
		})
	})

	// Guardar los datos en el archivo JSON
	return saveToJSON(filename, products)
}

// removeCopyrightMessage elimina los mensajes repetitivos de derechos de autor
func removeCopyrightMessage(description string) string {
	// Lista de mensajes a eliminar
	messages := []string{
		"TM & © 2023 BURGER KING CORPORATION. SE UTILIZA BAJO LICENCIA. TODOS LOS DERECHOS RESERVADOS. IMAGENES REFERENCIALES.",
		"TM & © 2022 BURGER KING CORPORATION. SE UTILIZA BAJO LICENCIA. TODOS LOS DERECHOS RESERVADOS. IMAGENES REFERENCIALES.",
		"TM & © 2024 BURGER KING CORPORATION. SE UTILIZA BAJO LICENCIA. TODOS LOS DERECHOS RESERVADOS. IMAGENES REFERENCIALES. ESTE COMBO INCLUYE PAPAS Y BEBIDAS PEQUEÑAS.",
		"TM \u0026 © 2023 BURGER KING CORPORATION. SE UTILIZA BAJO LICENCIA. TODOS LOS DERECHOS RESERVADOS. IMAGENES REFERENCIALES.",
		"TM \u0026 © 2022 BURGER KING CORPORATION. SE UTILIZA BAJO LICENCIA. TODOS LOS DERECHOS RESERVADOS. IMAGENES REFERENCIALES.",
		"TM \u0026 © 2024 BURGER KING CORPORATION. SE UTILIZA BAJO LICENCIA. TODOS LOS DERECHOS RESERVADOS. IMAGENES REFERENCIALES. ESTE COMBO INCLUYE PAPAS Y BEBIDAS PEQUEÑAS.",
	}

	// Eliminar cualquier mensaje encontrado
	for _, msg := range messages {
		description = strings.Replace(description, msg, "", -1)
	}

	// Eliminar cualquier espacio extra que pueda quedar
	return strings.TrimSpace(description)
}

// saveToJSON guarda los datos en un archivo JSON
func saveToJSON(filename string, data interface{}) error {
	// Crear la carpeta "data" si no existe
	err := os.MkdirAll("data", os.ModePerm)
	if err != nil {
		return fmt.Errorf("error al crear la carpeta 'data': %v", err)
	}

	// Crear el archivo JSON
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error al crear el archivo: %v", err)
	}
	defer file.Close()

	// Codificar los datos como JSON
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}

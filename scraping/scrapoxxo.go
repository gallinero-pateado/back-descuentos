package scraping

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

// Cupon estructura de los datos del cupón
type Discount struct {
	ID          int    `json:"id"`
	Titulo      string `json:"name"`
	Categoria   string `json:"category"`
	Descripcion string `json:"description"`
	Precio      string `json:"price"`
	Imagen      string `json:"image"`
	Logo        string `json:"logo,omitempty"`
}

// Logo representa el logo en la estructura JSON
type logoOxxo struct {
	Logo string `json:"logo"`
}

// Scrappi realiza el scraping a OXXO
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
	logo, exists := doc.Find("a.css-115kwlw").Attr("src")
	if !exists || logo == "" {
		logo = "No disponible"
	}

	var discounts []Discount

	// Iteramos sobre los descuentos
	doc.Find("div.promotions__card").Each(func(i int, s *goquery.Selection) {
		// Verificación de los elementos dentro del div
		titulo := s.Find("h3").Text()                // Suponiendo que el título está en un <h3>
		descripcion := s.Find(".description").Text() // Ajustar el selector
		precio := s.Find(".price").Text()            // Ajustar el selector

		// Verificación de que se están extrayendo correctamente los valores
		if titulo == "" {
			titulo = "No disponible"
		}
		if descripcion == "" {
			descripcion = "No disponible"
		}
		if precio == "" {
			precio = "Sin precio"
		}

		imagen, _ := s.Find("img.img-fluid").Attr("src") // Ajustar el selector
		if imagen == "" {
			imagen = "No disponible"
		}

		// Mostrar información de los cupones para depuración
		fmt.Printf("Descuento %d - Titulo: %s, Descripción: %s, Precio: %s, Imagen: %s\n", i+1, titulo, descripcion, precio, imagen)

		discounts = append(discounts, Discount{
			ID:          i + 1,
			Titulo:      titulo,
			Categoria:   "Saludable", // Cambiar según la categoría correcta
			Descripcion: descripcion,
			Precio:      precio,
			Imagen:      imagen,
			Logo:        "", // No se necesita el logo en cada cupón
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

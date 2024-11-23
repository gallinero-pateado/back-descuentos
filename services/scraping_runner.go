package services

import (
	"log"
)

// ejecuta el scraping y guarda resultados
func EjecutarScraping() {

	if err := ScrapingWendys("services/data/wendys.json"); err != nil {
		log.Printf("Error al realizar el scraping a la página de Wendy's: %v", err)
	}

	if err := ScrapingBurger("services/data/burgerking.json"); err != nil {
		log.Printf("Error al realizar el scraping a la página de Wendy's: %v", err)
	}

	if err := ScrapingLittleCaesars("services/data/little_caesars.json"); err != nil {
		log.Printf("Error al realizar el scraping a la página de Wendy's: %v", err)
	}

	if err := ScrapingOxxo("services/data/oxxo.json"); err != nil {
		log.Printf("Error al realizar el scraping a la página de Wendy's: %v", err)
	}
}

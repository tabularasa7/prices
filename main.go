package main

import (
	"fmt"
	"log"

	webscraper "hospital-prices/web_scraper"
)

func main() {
	csvFiles, err := webscraper.ScrapeFiles("https://www.healthonecares.com/patient-resources/patient-financial-resources/pricing-transparency-cms-required-file-of-standard-charges")

	if err != nil {
		log.Fatalf("unable to scrape files due to error: %v", err)
	}

	fmt.Println("CSV Files:")

	for _, csv := range csvFiles {
		fmt.Println(csv)
	}
}

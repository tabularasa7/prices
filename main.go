package main

import (
	"fmt"
	"log"
	"strings"
	"sync"

	filemanager "hospital-prices/file_manager"
	"hospital-prices/hospital"
	webscraper "hospital-prices/web_scraper"
)

func main() {
	var wait sync.WaitGroup
	hospital := hospital.HospitalSystem{
		Name:       "HealthOne Cares",
		Url:        "https://www.healthonecares.com",
		PricingURL: "/patient-resources/patient-financial-resources/pricing-transparency-cms-required-file-of-standard-charges",
	}

	csvFiles, err := webscraper.ScrapeFiles(hospital.Url + hospital.PricingURL)

	if err != nil {
		log.Fatalf("unable to scrape files due to error: %v", err)
	}

	fmt.Println("CSV Files:")

	for _, csv := range csvFiles {
		wait.Add(1)
		fmt.Println(csv)
		parseStringsSlash := strings.Split(csv, "/")
		parseStringsQuestion := strings.Split(parseStringsSlash[7], "?")
		fmt.Println(parseStringsQuestion[0])

		go filemanager.DownloadFile(hospital.Url+csv, parseStringsQuestion[0], &wait)

		if err != nil {
			log.Fatalln(err)
		}
	}

	wait.Wait()
}

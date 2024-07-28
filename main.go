package main

import (
	"log"

	filemanager "hospital-prices/file_manager"
	fileparser "hospital-prices/file_parser"
	"hospital-prices/hospital"
	webscraper "hospital-prices/web_scraper"
)

// type Prices struct {
// 	hospital_name   string
// 	last_updated_on string
// 	version         string
// 	affirmation     Affirmation
// }

// type Affirmation struct {
// 	affirmation string
// 	confirm_affirmation bool
// }

func main() {
	hospital := hospital.HospitalSystem{
		Name:       "HealthOne Cares",
		Url:        "https://www.healthonecares.com",
		PricingURL: "/patient-resources/patient-financial-resources/pricing-transparency-cms-required-file-of-standard-charges",
	}

	files, err := webscraper.ScrapeFiles(hospital.Url + hospital.PricingURL)

	if err != nil {
		log.Fatalf("unable to scrape files due to error: %v", err)
	}

	fileNames, err := filemanager.ParseFileNamesAndDownload(files)

	if err != nil {
		log.Fatal(err)
	}

	_, err = fileparser.ParseJSONFiles(fileNames)

	if err != nil {
		log.Fatal(err)
	}

}

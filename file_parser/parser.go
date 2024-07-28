package fileparser

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func RunPythonFileParser(filename string) error {
	out, err := exec.Command("python3", "parser.py", filename).Output()

	if err != nil {
		return err
	}

	fmt.Println(out)

	return nil
}

func ParseJSONFiles(fileNames []string) ([]interface{}, error) {
	// TODO: create separate method for parsing JSON prices files
	file, err := os.Open(fileNames[0])

	if err != nil {
		return nil, err
	}

	jsonFile, err := io.ReadAll(file)

	if err != nil {
		return nil, err
	}

	var prices map[string]interface{}

	err = json.Unmarshal(jsonFile, &prices)

	if err != nil {
		return nil, err
	}

	var standardCharges []interface{}

	for k, v := range prices {
		if k == "standard_charge_information" {
			standardCharges = v.([]interface{})
		}
	}

	fmt.Println(standardCharges[0])
	return standardCharges, nil
}

func LoopInterfaces(value interface{}, fileNames *[]string, terminationVal string) {
	switch inVal := value.(type) {
	case map[string]interface{}:
		for _, val := range inVal {
			LoopInterfaces(val, fileNames, terminationVal)
		}
	case []interface{}:
		for _, val := range inVal {
			LoopInterfaces(val, fileNames, terminationVal)
		}
	case string:
		// currently scraping json files, looking for standard charges files
		if strings.Contains(inVal, terminationVal) {
			*fileNames = append(*fileNames, inVal)
		}
	default:
		// do nothing
	}
}

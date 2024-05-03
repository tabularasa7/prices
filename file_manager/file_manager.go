package filemanager

import (
	"bytes"
	"encoding/csv"
	"os"
)

// func DownloadFile(fileUrl string)

func NewFile(fileName string) (*csv.Reader, error) {
	file, err := os.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(bytes.NewReader(file))

	return reader, nil
}

func Read(reader *csv.Reader) ([]string, error) {

	record, err := reader.Read()

	if err != nil {
		return record, err
	}

	return record, nil

}

func ReadFile(reader *csv.Reader) ([][]string, error) {

	records, err := reader.ReadAll()

	if err != nil {
		return nil, err
	}

	return records, nil
}

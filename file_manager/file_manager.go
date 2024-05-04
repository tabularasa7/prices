package filemanager

import (
	"bytes"
	"encoding/csv"
	"io"
	"net/http"
	"os"
	"path"
	"sync"
)

func DownloadFile(fileUrl, filePath string, wait *sync.WaitGroup) error {
	defer wait.Done()
	resp, err := http.Get(fileUrl)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	out, err := os.Create(path.Base(filePath))

	if err != nil {
		return err
	}

	defer out.Close()

	_, err = io.Copy(out, resp.Body)

	return err

}

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

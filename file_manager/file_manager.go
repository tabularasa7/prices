package filemanager

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
	"sync"
)

func ParseFileNamesAndDownload(files []string) ([]string, error) {
	var wait sync.WaitGroup
	fmt.Println("Files:", files)

	fileNames := make([]string, len(files))

	// TODO: handle errors from Download file
	for index, file := range files {
		wait.Add(1)
		fmt.Println(file)
		parseStringsSlash := strings.Split(file, "/")
		parseStringsQuestion := strings.Split(parseStringsSlash[6], "?")
		parseStringsEndPunc := strings.TrimRight(parseStringsQuestion[0], ".")
		fileNames[index] = parseStringsEndPunc + ".json"
		fmt.Println(parseStringsEndPunc)

		go DownloadFile(file, parseStringsEndPunc+".json", &wait)

	}

	fmt.Printf("fileName: %v\n", fileNames)

	wait.Wait()

	return fileNames, nil
}

func DownloadFile(fileUrl, saveFileName string, wait *sync.WaitGroup) error {
	defer wait.Done()
	resp, err := http.Get(fileUrl)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	out, err := os.Create(path.Base(saveFileName))

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

package httpclient

import (
	"fmt"
	"io"
	"net/http"
)

func GetURL(url string) (io.ReadCloser, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	fmt.Println("Response Status:", resp.Status)

	return resp.Body, nil

}

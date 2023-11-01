package parse

import (
	"io"
	"net/http"
	"strings"
)

func GrabAPI(url string, params map[string]string) []byte {
	var queryString string
	if len(params) > 0 {
		var queryParams []string
		for key, value := range params {
			queryParams = append(queryParams, key+"="+value)
		}
		queryString = "?" + strings.Join(queryParams, "&")
	}

	fullURL := url + queryString

	response, err := http.Get(fullURL)
	if err != nil {
		return nil
	}
	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return nil
	}

	return responseData
}

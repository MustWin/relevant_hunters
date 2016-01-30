package api

import (
	"io/ioutil"
	"net/http"
	"os"
)

func readGet(link string) ([]byte, error) {
	request, err := http.NewRequest("GET", link, nil)
	if err != nil {
		return []byte{}, err
	}
	request.Header.Add("Authorization", "Bearer "+os.Getenv("PH_BEARER_TOKEN"))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Accept", "application/json")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return []byte{}, err
	}
	defer response.Body.Close()

	return ioutil.ReadAll(response.Body)
}

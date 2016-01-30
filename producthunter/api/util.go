package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/MustWin/relevant_hunters/producthunter"
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

	limit, _ := strconv.ParseInt(response.Header.Get("X-Rate-Limit-Limit"), 10, 32)
	remaining, _ := strconv.ParseInt(response.Header.Get("X-Rate-Limit-Remaining"), 10, 32)
	reset, _ := strconv.ParseInt(response.Header.Get("X-Rate-Limit-Reset"), 10, 32)
	log.Printf("Request Limits: limit: %d, remaining: %d, til reset: %d\n", limit, remaining, reset)

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []byte{}, err
	}

	if err = testForError(body); err != nil {
		return []byte{}, err
	}

	return body, nil
}

func testForError(body []byte) error {
	var eo producthunter.ErrorResponse
	err := json.Unmarshal(body, &eo)
	if err != nil {
		return err
	}

	if eo.Error != "" {
		return fmt.Errorf("%s\n%s", eo.Error, eo.Description)
	}

	return nil
}

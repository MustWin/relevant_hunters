package api

import (
	"encoding/json"
	"fmt"

	"github.com/MustWin/relevant_hunters/producthunter"
)

const (
	base_url = "https://api.producthunt.com/"
)

func GetCollection(id int) (producthunter.Collection, error) {
	var coll producthunter.CollectionResponse

	body, err := readGet(fmt.Sprintf("%s/v1/collections/%d", base_url, id))
	if err != nil {
		return coll.Collection, err
	}

	if err = testForError(body); err != nil {
		return coll.Collection, err
	}

	err = json.Unmarshal(body, &coll)
	if err != nil {
		return coll.Collection, err
	}

	return coll.Collection, nil
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

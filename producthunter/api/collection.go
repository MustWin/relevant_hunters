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

	err = json.Unmarshal(body, &coll)
	if err != nil {
		return coll.Collection, err
	}

	return coll.Collection, nil
}

func GetPostVotes(postId int) ([]producthunter.Vote, error) {
	var votes producthunter.VoteResponse

	body, err := readGet(fmt.Sprintf("%s/v1/posts/%d/votes", base_url, postId))
	if err != nil {
		return votes.Votes, err
	}

	err = json.Unmarshal(body, &votes)
	if err != nil {
		return votes.Votes, err
	}

	return votes.Votes, nil
}

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
	var voteResponse producthunter.VoteResponse
	var votes []producthunter.Vote
	var last int
	var request_url string
	base_request_url := fmt.Sprintf("%s/v1/posts/%d/votes", base_url, postId)
	hasMore := true
	for hasMore {
		if last > 0 {
			request_url = fmt.Sprintf("%s?older=%d", base_request_url, last)
		} else {
			request_url = base_request_url
		}
		body, err := readGet(request_url)
		if err != nil {
			return votes, err
		}

		err = json.Unmarshal(body, &voteResponse)
		if err != nil {
			return votes, err
		}
		votes = append(votes, voteResponse.Votes...)
		hasMore = len(voteResponse.Votes) == 50
		last = voteResponse.Votes[len(voteResponse.Votes)-1].Id
	}

	return votes, nil
}

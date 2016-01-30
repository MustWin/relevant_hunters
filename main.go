package main

import (
	"fmt"
	"os"

	"github.com/MustWin/relevant_hunters/producthunter"
	"github.com/MustWin/relevant_hunters/producthunter/api"
)

func main() {
	coll, err := api.GetCollection(35215)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return
	}

	fmt.Fprintf(os.Stdout, "Collection:\n%+v\n", coll)

	users := map[int]producthunter.User{}
	for _, post := range coll.Posts {
		votes, err := api.GetPostVotes(post.Id)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			continue
		}
		fmt.Fprintf(os.Stdout, "Post %q; %d(%d) votes\n", post.Name, post.VotesCount, len(votes))
		for _, vote := range votes {
			if _, present := users[vote.UserId]; !present {
				continue
				if user, err := api.GetUserDetails(vote.UserId); err != nil {
					fmt.Fprintf(os.Stderr, "Problem retrieving userId %d; %v\n", vote.UserId, err)
				} else {
					users[vote.UserId] = user
				}
			}
		}

		fmt.Fprintf(os.Stdout, "Users for post %q:\n%+v\n", post.Name, users)
	}
}

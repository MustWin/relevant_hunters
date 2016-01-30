package main

import (
	"fmt"
	"os"

	"github.com/MustWin/relevant_hunters/producthunter/api"
)

func main() {
	coll, err := api.GetCollection(35215)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return
	}

	fmt.Fprintf(os.Stdout, "Collection:\n%+v\n", coll)

	for _, post := range coll.Posts {
		votes, err := api.GetPostVotes(post.Id)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			continue
		}
		fmt.Fprintf(os.Stdout, "%d votes\n", len(votes))
	}
}

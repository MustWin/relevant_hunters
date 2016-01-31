package main

import (
	"encoding/csv"
	"sync"
	"time"
	"fmt"
	"os"
	"strings"
	"strconv"
	//"github.com/MustWin/relevant_hunters/producthunter"
	//"github.com/MustWin/relevant_hunters/producthunter/api"
	"github.com/MustWin/relevant_hunters/Godeps/_workspace/src/github.com/MustWin/gohunt/gohunt"
)

type RelevantUser struct {
	User            *gohunt.User
	RelevantUpvotes []string
}

func (RelevantUser) ToCSVHeaders() []string {
    return []string{"ID", "TWITTER", "NAME", "FOLLOWER COUNT", "VOTE COUNT", "POST COUNT", "RELEVANT UPVOTE COUNT", "RELEVANT UPVOTES"}
}
func (ref RelevantUser) ToCSV() []string {
    return []string{strconv.Itoa(ref.User.ID), ref.User.TwitterUsername, ref.User.Name, strconv.Itoa(len(ref.User.Followers)),
                    strconv.Itoa(int(ref.User.VotesCount)), strconv.Itoa(int(ref.User.PostsCount)), strconv.Itoa(len(ref.RelevantUpvotes)),
                    strings.Join(ref.RelevantUpvotes, " | ")}
}

// GLOBAL STATE, not syncrhonized
var users map[int]RelevantUser
var wg sync.WaitGroup
var phClient *gohunt.Client

func fetchCollection(channel chan func() error, id int) {
    wg.Add(1)
    go func() {
	    channel <- func() error {
            defer wg.Done()
	    	fmt.Fprintf(os.Stderr, "Fetching collection: %v\n", id)
	    	coll, err := phClient.GetCollection(id)
	    	if err != nil {
	    		return err
	    	}
	    	//fmt.Fprintf(os.Stderr, "Collection:\n%+v\n", coll)
	    	for _, post := range coll.Posts {
	    		fetchPostVotes(channel, post)
	    	}
            return nil
	    }

    }()

}

func addRelevantUpvote(user gohunt.User, post gohunt.Post) {
	newUser := users[user.ID]
	newUser.RelevantUpvotes = append(users[user.ID].RelevantUpvotes, post.DiscussionUrl)
	users[user.ID] = newUser
}

func fetchPostVotes(channel chan func() error, post gohunt.Post) {
    wg.Add(1)
    go func() {
	    channel <- func() error {
            defer wg.Done()
	    	fmt.Fprintf(os.Stderr, "Fetching votes for: %v\n", post.ID)
	    	votes, err := phClient.GetPostVotes(post.ID, -1, -1, 1000, "")
	    	if err != nil {
	    		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	    		return err
	    	}
	    	fmt.Fprintf(os.Stderr, "Post %q; %d(%d) votes\n", post.Name, post.VotesCount, len(votes))
	    	for _, vote := range votes {
	    		if _, present := users[vote.User.ID]; !present {
	    			fetchUser(channel, post, vote.User)
	    		} else {
	    			addRelevantUpvote(vote.User, post)
	    		}
	    	}
	    	fmt.Fprintf(os.Stderr, "Users for post %q:\n%+v\n", post.Name, users)
	    	return nil
	    }
    }()
}

func fetchUser(channel chan func() error, post gohunt.Post, user gohunt.User) {
	// Short circuit if it already exists
	if _, present := users[user.ID]; present {
		addRelevantUpvote(user, post)
		return
	}
    wg.Add(1)
    go func() {
	    channel <- func() error {
            defer wg.Done()
	    	fmt.Fprintf(os.Stderr, "Fetching user: %d\n", user.ID)
	    	if user, err := phClient.GetUser(strconv.Itoa(user.ID)); err != nil {
	    		fmt.Fprintf(os.Stderr, "Problem retrieving userId %d; %v\n", user.ID, err)
	    		return err
	    	} else {
	    		users[user.ID] = RelevantUser{User: &user, RelevantUpvotes: []string{post.DiscussionUrl}}
	    	}
	    	return nil
	    }
    }()
}

func outputCSV(users map[int]RelevantUser) {
    w := csv.NewWriter(os.Stdout)
    headersPrinted := false
    for id, relUser := range users {
        if !headersPrinted {
            if err := w.Write(relUser.ToCSVHeaders()); err != nil {
			    fmt.Fprintf(os.Stderr, "error writing headers to CSV")
            }
            headersPrinted = true
        }
        if err := w.Write(relUser.ToCSV()); err != nil {
		    fmt.Fprintf(os.Stderr, "error writing user %d to CSV", id)
        }
    }
}


func main() {

	phClient = gohunt.NewUserClient(os.Getenv("PH_BEARER_TOKEN"))
    users = make(map[int]RelevantUser)

	requestQueue := make(chan func() error)

    go func() {
        // Process the request queue one at a time, sleep if we get rate limited
        for {
            fn := <- requestQueue
	    	err := fn()
	    	if err != nil {
	    		fmt.Fprintf(os.Stdout, "Error: %v\n", err)
	    		// detect rate limit, sleep until rate limit is up, retry
                if respErr, ok := err.(gohunt.ResponseError); ok {
	                fmt.Fprintf(os.Stderr, "HERE3\n")
                    wg.Add(1)
	    		    requestQueue <- fn
                     time.Sleep(time.Duration(respErr.LimitedForSecs) * time.Second)
                }
	    	}
	    }
    }()

    // TODO: abstract this
    // Begin queing up the fetching
    fetchCollection(requestQueue, 35215)


    // Wait for all requests to complete
    wg.Wait()
    // Output Users as CSV
    outputCSV(users)
	fmt.Fprintf(os.Stderr, "DONE")
}

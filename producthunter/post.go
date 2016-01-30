package producthunter

type Post struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	Tagline       string `json:"tagline"`
	CategoryId    int    `json:"category_id"`
	CommentsCount int    `json:"comments_count"`
	VotesCount    int    `json:"votes_count"`
}

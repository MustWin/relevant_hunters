package producthunter

type VoteResponse struct {
	Votes []Vote `json:"votes"`
}

type Vote struct {
	Id     int  `json:"id"`
	PostId int  `json:"post_id"`
	UserId int  `json:"user_id"`
	User   User `json:"user"`
}

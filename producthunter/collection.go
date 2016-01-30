package producthunter

type CollectionResponse struct {
	Collection Collection `json:"collection"`
}

type Collection struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Title       string `json:"title"`
	Subscribers int    `json:"subscriber_count"`
	CategoryId  int    `json:"category_id"`
	Url         string `json:"collection_url"`
	Posts       []Post `json:"posts"`
}

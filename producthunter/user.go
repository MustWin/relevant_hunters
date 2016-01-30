package producthunter

type User struct {
	Id          int               `json:"id"`
	Name        string            `json:"name"`
	Headline    string            `json:"headline"`
	UserName    string            `json:"username"`
	TwitterName string            `json:"twitter_username"`
	Website     string            `json:"website_url"`
	Profile     string            `json:"profile_url"`
	Images      map[string]string `json:"image_url"`
}

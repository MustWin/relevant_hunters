package producthunter

type ErrorResponse struct {
	Error       string `json:"error"`
	Description string `json:"error_description"`
}

package repository

type LambdaRequest struct {
	Body string `json:"body"`
}

type LambdaResponse struct {
	StatusCode int    `json:"statusCode"`
	Body       string `json:"body"`
}

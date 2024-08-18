type ErrorResponse struct {
	Message string `json:"message"`
}

type Handler struct {
	storage Storage
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type Handler struct {
	storage Storage
}

// NewHandler returns pointer to the Handler
// and implements Dependency Injection pattern
func NewHandler(storage Storage) *Handler {
	return &Handler{
		storage: storage
	}
}
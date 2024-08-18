type Storage interface {
	Insert(e *Employee)
	Get(id int) (Employee, error)
	List() []Employee
	Update(id int, e *Employee) error
	Delete(id int) error
}

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


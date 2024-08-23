package interfaces

type Logger interface {
	Error(msg string, args ...any)
}

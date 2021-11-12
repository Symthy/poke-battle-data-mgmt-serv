package value

type OptionalId interface {
	Get() (*int, error)
	isEmpty() bool
}

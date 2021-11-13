package optional

type OptionalId interface {
	Get() (*int, error)
}

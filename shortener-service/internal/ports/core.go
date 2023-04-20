package ports

type CorePort interface {
	RandomURL() (string, error)
}

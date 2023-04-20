package ports

type CorePort interface {
	Hash(password []byte) ([]byte, error)
}

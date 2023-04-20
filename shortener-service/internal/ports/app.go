package ports

type AppPort interface {
	CreateURL(url string) (string, error)
	GetURL(shortCode string) (string, error)
	DeleteURL(id string) error
}

package ports

type RedisPort interface {
	SetToken(token string) error
	IsHasToken(token string) bool
}

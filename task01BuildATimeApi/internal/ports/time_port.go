package ports

type TimeProvider interface {
	GetCurrentTime(tz string) (string, error)
}
package reader

type I interface {
	Start(filter string) error
	Stop() error
}

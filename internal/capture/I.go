package capture

type I interface {
	Start(dev string) error
	Pause() error
	Resume() error
	Stop() error
}

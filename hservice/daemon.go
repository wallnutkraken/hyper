package hservice

const (
	daemonName = "hyperd"
	daemonDesc = "A glorified web redirector"
)

type Daemon interface {
	Install() error
	Name() string
	Description() string
	Uninstall() error
	Start() error
	Stop() error
}
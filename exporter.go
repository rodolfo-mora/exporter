package exporter

type Exporter interface {
	Register() error
	Export()
}

package exporter

type Exporter interface {
	Export()
}

func NewExporter(exporter string) Exporter {
	pe := NewPrometheusExporter(":8080")
	return pe
}

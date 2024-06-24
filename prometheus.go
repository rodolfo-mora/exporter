package exporter

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	Name = "gauge"
	Help = "Generic gauge label"
)

type Prometheus struct {
	Gauge   *prometheus.GaugeVec
	Port    string
	Tracker []string
}

func NewPrometheusExporter(port string) Prometheus {
	prom := Prometheus{
		Gauge: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: Name,
				Help: Help,
			},
			[]string{"floki"},
		),
		Tracker: []string{},
		Port:    port,
	}
	prometheus.MustRegister(prom.Gauge)
	return prom
}

func (p *Prometheus) Register(gauge string) {
	p.Gauge.WithLabelValues(gauge).Set(float64(1))
}

func (p Prometheus) Export() {
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(p.Port, nil))
}

func (p *Prometheus) Track(name string) {
	p.Tracker = append(p.Tracker, name)
}

func (p Prometheus) HostDown(hostname string) {
	p.Gauge.WithLabelValues(hostname).Set(float64(0))
}

func (p Prometheus) Display() {
	log.Println(p.Tracker)
}

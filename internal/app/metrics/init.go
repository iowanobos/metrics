package metrics

import (
	"github.com/iowanobos/metrics/internal/app/config"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

const (
	namespace = "iowanobos"
	subsystem = "metrics"

	currencyLabel = "currency"
)

var (
	currenciesMetrics = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "currency_rate",
			Help:      "Current currency rate by configuration file",
		},
		[]string{currencyLabel},
	)
)

func init() {
	log.Debug("Initialize metrics registration")
	prometheus.MustRegister(currenciesMetrics)
	cfg := config.Get()
	registerByConfig(cfg)
}

func registerByConfig(cfg config.Configuration) {
	for _, currency := range cfg.Currencies {
		labels := prometheus.Labels{currencyLabel: currency.Name}
		currenciesMetrics.With(labels).Set(currency.Value)
	}
}

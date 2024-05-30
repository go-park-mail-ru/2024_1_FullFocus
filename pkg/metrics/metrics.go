package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type Collector interface {
	IncreaseHits(path string)
	IncreaseErr(statusCode string, path string)
	AddDurationToHistogram(path string, duration time.Duration)
	AddDurationToSummary(statusCode string, path string, duration time.Duration)
}

type Metrics struct {
	totalHits         *prometheus.CounterVec
	totalErrors       *prometheus.CounterVec
	durationHistogram *prometheus.HistogramVec
	durationSummary   *prometheus.SummaryVec
}

func NewMetrics(reg prometheus.Registerer) *Metrics {
	labelHits := []string{"path"}
	totalHits := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "hits_total",
	}, labelHits)
	reg.MustRegister(totalHits)

	labelErrors := []string{"status_code", "path"}
	totalErrors := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "errors_total",
	}, labelErrors)
	reg.MustRegister(totalErrors)

	labelHistogram := []string{"path"}
	durationHistogram := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "request_durations",
	}, labelHistogram)
	reg.MustRegister(durationHistogram)

	labelSummary := []string{"status_code", "path"}
	durationSummary := prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name: "durations_summary",
		Objectives: map[float64]float64{
			0.5:  0.1,
			0.8:  0.1,
			0.9:  0.1,
			0.95: 0.1,
			0.99: 0.1,
			1:    0.1,
		}}, labelSummary)
	reg.MustRegister(durationSummary)

	return &Metrics{
		totalHits:         totalHits,
		totalErrors:       totalErrors,
		durationHistogram: durationHistogram,
		durationSummary:   durationSummary,
	}
}

func (m *Metrics) IncreaseHits(path string) {
	m.totalHits.WithLabelValues(path).Inc()
}

func (m *Metrics) IncreaseErr(statusCode string, path string) {
	m.totalErrors.WithLabelValues(statusCode, path).Inc()
}

func (m *Metrics) AddDurationToHistogram(path string, duration time.Duration) {
	m.durationHistogram.WithLabelValues(path).Observe(duration.Seconds())
}

func (m *Metrics) AddDurationToSummary(statusCode string, path string, duration time.Duration) {
	m.durationSummary.WithLabelValues(statusCode, path).Observe(duration.Seconds())
}

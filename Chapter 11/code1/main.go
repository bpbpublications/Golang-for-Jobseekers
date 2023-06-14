package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	apiRequests = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_request",
		Help: "The total number of processed events",
	})
	durationRequest = promauto.NewHistogram(prometheus.HistogramOpts{
		Name: "myapp_request_duration_seconds",
		Help: "Duration of myapp api request call",
	})
	errorsRequest = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_request_errors",
		Help: "Errors from processing myapp api request",
	})
)

type API struct{}

func (z *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	defer func() {
		timeDiff := time.Since(startTime)
		durationRequest.Observe(timeDiff.Seconds())
	}()
	apiRequests.Inc()
	sleepTime := 1 * time.Second
	rawSleepTime := r.URL.Query().Get("sleep")
	if rawSleepTime != "" {
		sleepTimeInt, _ := strconv.Atoi(rawSleepTime)
		sleepTime = time.Duration(sleepTimeInt)
	}
	time.Sleep(sleepTime)

	errorParam := r.URL.Query().Get("error")
	if errorParam == "true" {
		errorsRequest.Inc()
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("errored-out on request"))
		return
	}

	w.Write([]byte("testing"))
}

func main() {
	fmt.Println("Start server")
	http.Handle("/api/items", &API{})
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}

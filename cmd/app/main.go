package main

import (
	"flag"
	"net/http"
	"strconv"

	_ "github.com/iowanobos/metrics/internal/app/config"
	_ "github.com/iowanobos/metrics/internal/app/metrics"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

var port = flag.Int("p", 8080, "Provide HTTP-server port")

func init() {
	flag.Parse()
}

func main() {
	http.Handle("/metrics", promhttp.Handler())
	log.WithField("port", *port).Debug("Starting server")
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(*port), nil))
}

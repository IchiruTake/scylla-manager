// Copyright (C) 2017 ScyllaDB

package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/scylladb/mermaid"
)

var (
	currentVersion = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "scylla_manager",
		Subsystem: "server",
		Name:      "current_version",
		Help:      "Current Scylla Manager version.",
	}, []string{"version"})
)

func init() {
	prometheus.MustRegister(currentVersion)
	currentVersion.WithLabelValues(mermaid.Version()).Set(0)
}

package collectors

import (
	"os/exec"
	"os"
	"strconv"
	"strings"

	"github.com/golang/glog"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	cpuHeatCelsiusKey = "cpu_heat_celsius"
	nodeNameKey = "NODE"
	podNameKey = "HOSTNAME"
)

var (
	cpuHeatCelsius = prometheus.NewDesc(
		cpuHeatCelsiusKey,
		"CPU temperatur in celsius",
		[]string{"node", "pod"}, nil,
	)
)

type machineCollector struct {
	
}

// NewMachineCollector creates a new wmbusmeters prometheus collector.
func NewMachineCollector() prometheus.Collector {
	return &machineCollector{}
}

// Describe implements the prometheus.Collector interface.
func (collector *machineCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- cpuHeatCelsius
}

// Collect implements the prometheus.Collector interface.
func (collector *machineCollector) Collect(ch chan<- prometheus.Metric) {

	addGauge := func(desc *prometheus.Desc, node string, pod string, v float64, lv ...string) {
		lv = append([]string{node, pod}, lv...)
		ch <- prometheus.MustNewConstMetric(desc, prometheus.GaugeValue, v, lv...)
	}

	cmd := exec.Command("cat", "/sys/class/thermal/thermal_zone0/temp")
	out, err := cmd.Output()
	if err != nil {
		glog.Errorf("could not get CPU temp : %s ", err)
	}
	line := strings.TrimSuffix(string(out), "\n");
	value, err := strconv.ParseFloat(line, 64)
	if err != nil {
		glog.Errorf("could not parse CPU temp : %s  (%s)", string(out), err)
	}
	addGauge(cpuHeatCelsius, os.Getenv(nodeNameKey), os.Getenv(podNameKey), value/1000)
}

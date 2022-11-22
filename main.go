package main

import (
	"firstwork/aa_collector"
	"firstwork/bb_collector"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func main() {
	// 创建一个自定义的注册表
	registry1 := prometheus.NewRegistry()

	registry2 := prometheus.NewRegistry()
	// 创建一个简单呃 gauge 指标。
	temp1 := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "home_temperature_celsius11111",
		Help: "The current temperature in degrees Celsius111111111.",
	})
	temp2 := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "home_temperature_cels22222ius",
		Help: "The current temperature in degrees Celsius222222.",
	})

	// 使用我们自定义的注册表注册 gauge
	registry1.MustRegister(temp1)
	registry1.MustRegister(temp2)
	registry2.MustRegister(temp2)

	// 设置 gague 的值为 39
	temp1.Set(11111111)
	temp2.Set(222222222)

	workerAA := aa_collector.NewClusterManager("aaa")
	workerBB := bb_collector.NewClusterManager("bbb")

	rega := prometheus.NewPedanticRegistry()
	rega.MustRegister(workerAA)
	regb := prometheus.NewPedanticRegistry()
	regb.MustRegister(workerBB)
	// 暴露自定义指标
	http.Handle("/metrics/1", promhttp.HandlerFor(registry1, promhttp.HandlerOpts{Registry: registry1}))
	http.Handle("/metrics/2", promhttp.HandlerFor(registry2, promhttp.HandlerOpts{Registry: registry2}))
	http.Handle("/metrics/aa", promhttp.HandlerFor(rega, promhttp.HandlerOpts{}))
	http.Handle("/metrics/bb", promhttp.HandlerFor(regb, promhttp.HandlerOpts{}))
	http.ListenAndServe(":8080", nil)
}

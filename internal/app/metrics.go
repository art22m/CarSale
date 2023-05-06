package app

import (
	"github.com/prometheus/client_golang/prometheus"
)

var NewCarSalesCounter = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "new_car_sale",
	Help: "New car sale created",
})

var NewSellersCounter = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "new_seller",
	Help: "New seller created",
})

func init() {
	prometheus.MustRegister(NewCarSalesCounter)
	prometheus.MustRegister(NewSellersCounter)
}

package main

import (
	"orderbook/orderbook"

	"github.com/gin-gonic/gin"
)

var ob *orderbook.OrderBook

func init() {
	ob = orderbook.NewOrderBook()
}

func main() {

	r := gin.Default()
	r.GET("/", Index)
	r.Run(":8088")
}

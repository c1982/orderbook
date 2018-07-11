package main

import (
	"html/template"
	"orderbook/orderbook"

	"github.com/gin-gonic/gin"
)

var ob *orderbook.OrderBook

func init() {
	ob = orderbook.NewOrderBook()
}

func main() {

	t := template.Must(template.ParseGlob("ui/*.html"))

	r := gin.Default()
	r.Static("/js", "./ui/js")
	r.SetHTMLTemplate(t)

	r.GET("/", Index)
	r.Run(":9099")
}

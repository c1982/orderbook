package main

import (
	"html/template"
	"orderbook/orderbook"
	"time"

	"github.com/gin-gonic/gin"
	melody "gopkg.in/olahol/melody.v1"
)

var ob *orderbook.OrderBook
var m *melody.Melody

func init() {
	ob = orderbook.NewOrderBook()
	m = melody.New()
	loadTestData()
}

func main() {

	t := template.Must(template.ParseGlob("ui/*.html"))

	r := gin.Default()
	r.SetHTMLTemplate(t)

	r.Static("/js", "./ui/js")

	r.GET("/", Index)
	r.POST("/add", AddOrder)

	r.GET("/book", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleConnect(func(s *melody.Session) {
		msg := ob.ToList()
		s.Write(msg)
	})

	r.Run(":9099")
}

func loadTestData() {

	ob.AddOrder(orderbook.Order{ID: 1, UserID: 100, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 1.1, Price: 30500, Time: time.Now()})
	ob.AddOrder(orderbook.Order{ID: 2, UserID: 101, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 0.20, Price: 30300, Time: time.Now()})
	ob.AddOrder(orderbook.Order{ID: 3, UserID: 102, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 0.8, Price: 30250, Time: time.Now()})

	//SELL
	ob.AddOrder(orderbook.Order{ID: 4, UserID: 104, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 0.90, Price: 30600, Time: time.Now()})
	ob.AddOrder(orderbook.Order{ID: 5, UserID: 105, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 0.75, Price: 30700, Time: time.Now()})
	ob.AddOrder(orderbook.Order{ID: 6, UserID: 107, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 0.20, Price: 31000, Time: time.Now()})
}

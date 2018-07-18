package main

import (
	"fmt"
	"net/http"
	"orderbook/orderbook"
	"time"

	"github.com/shopspring/decimal"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type OrderRequest struct {
	Side    string `json:"side"`
	Type    string `json:"type"`
	Amount  string `json:"amount"`
	Price   string `json:"price"`
	Easy    bool   `json:"easy"`
	Trigger string `json:"trigger"`
	Stop    bool   `json:"stop"`
}

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func AddOrder(c *gin.Context) {

	req := OrderRequest{}

	if err := c.ShouldBindWith(&req, binding.JSON); err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	amount, _ := decimal.NewFromString(req.Amount)
	price, _ := decimal.NewFromString(req.Price)
	trigger, _ := decimal.NewFromString(req.Trigger)

	o := orderbook.Order{}
	o.ID = time.Now().Unix()
	o.Type = req.Type
	o.Side = req.Side
	o.Time = time.Now()
	o.Base = "BTC"
	o.Second = "TRY"
	o.Easy = req.Easy
	o.Amount = amount
	o.Price = price
	o.Stop = trigger

	fmt.Printf("ID: %d, Side: %s, Type: %s, Amount: %s, Price: %s, Stop: %s\r\n", o.ID,
		o.Side,
		o.Type,
		o.Amount.String(),
		o.Price.String(),
		o.Stop.String())

	if req.Stop {
		ob.AddStop(o)
	} else {
		ob.AddOrder(o)
	}

	export := ob.ToList()

	fmt.Sprintln(string(export))

	ob.Debug()

	m.Broadcast(export)

	c.JSON(http.StatusOK, nil)
}

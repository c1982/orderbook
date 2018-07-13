package main

import (
	"fmt"
	"net/http"
	"orderbook/orderbook"
	"strconv"
	"time"

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

	amount, _ := strconv.ParseFloat(req.Amount, 64)
	price, _ := strconv.ParseFloat(req.Price, 64)
	trigger, _ := strconv.ParseFloat(req.Trigger, 64)

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

	fmt.Printf("ID: %d, Side: %s, Type: %s, Amount: %f, Price: %f, Stop: %f\r\n", o.ID, o.Side, o.Type, o.Amount, o.Price, o.Stop)

	if req.Stop {
		ob.AddStop(o)
	} else {
		ob.AddOrder(o)
	}

	export := ob.ToList()
	ob.Debug()

	m.Broadcast(export)

	c.JSON(http.StatusOK, nil)
}

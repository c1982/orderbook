package main

import (
	"net/http"
	"orderbook/orderbook"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type OrderRequest struct {
	Side    string  `json:"side"`
	Type    string  `json:"type"`
	Amount  float64 `json:"amount"`
	Price   float64 `json:"price"`
	Easy    bool    `json:"easy"`
	Trigger float64 `json:"trigger"`
	Stop    bool    `json:"stop"`
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

	o := orderbook.Order{}
	o.ID = time.Now().Unix()
	o.Type = req.Type
	o.Side = req.Side
	o.Amount = req.Amount
	o.Price = req.Price
	o.Time = time.Now()
	o.Base = "BTC"
	o.Second = "TRY"
	o.Easy = req.Easy
	o.Stop = req.Trigger

	if req.Stop {
		ob.AddStop(o)
	} else {
		ob.AddOrder(o)
	}

	export := ob.ToList()
	m.Broadcast(export)

	c.JSON(http.StatusOK, nil)
}

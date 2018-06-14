package matching

import (
	"sync"
	"time"
)

type Order struct {
	Side   string
	Status string
	Stop   float64
	Amont  float64
	Price  float64
}

type OrderList struct {
	sync.RWMutex
	Orders []Order
}

func (ol *OrderList) Append(item Order) {
	ol.Lock()
	defer ol.Unlock()

	ol.Orders = append(ol.Orders, item)
}

func (ol *OrderList) Item() <-chan Order {
	c := make(chan Order)

	f := func() {
		ol.Lock()
		defer ol.Unlock()
		for _, v := range ol.Orders {
			c <- v
		}
		close(c)
	}
	go f()

	return c
}

func (ol *OrderList) First() Order {
	return ol.Orders[0]
}

var Takers *OrderList
var Makers *OrderList

func Run() {

	for {

		time.Sleep(time.Second * 3)
	}
}

func Match(takers *OrderList, makers *OrderList) {

	for {
		tk := Takers.First()
		completed, fee := Collect(tk)
	}
}

func Collect(o Order) (completed bool, fee float64) {

	expect := o.Amont
	filled := []Order{}

	for _, v := range Makers.Orders {
		if v.Amont == expect {
			v.Status = "completed"
			completed = true
			break
		}

		if v.Amont < expect {
			filled = append(filled, v)
		}
	}

	return
}

package matching

import (
	"fmt"
)

const (
	BID = 1
	ASK = 2
)

type BidAsk struct {
	Entry struct {
		Bid int
		Ask int
	}
}

func (BidAsk) spread() int {
	return 0
}

type OrderBook struct {
	Bids map[int]int
	Asks map[int]int
}

func (ob *OrderBook) Run() {

	for i, v := range ob.Asks {
		fmt.Printf("%d\t%d\r", i, v)
	}

	for i, v := range ob.Bids {
		fmt.Printf("%d\t%d\r", i, v)
	}
}

func (ob *OrderBook) isEmpty() bool {
	return (len(ob.Bids) == 0) && (len(ob.Asks) == 0)
}

func (ob *OrderBook) add(price, amount int, bid bool) {

	if bid {
		ob.Bids[price] += amount
	} else {
		ob.Asks[price] += amount
	}
}

func (ob *OrderBook) AddBid(price, amount int) {
	ob.add(price, amount, true)
}

func (ob *OrderBook) AddAsk(price, amount int) {
	ob.add(price, amount, false)
}

func (ob *OrderBook) GetBidAsk() (ba BidAsk) {
	return
}

func (ob *OrderBook) rbegin() {

	best_bid := ob.Bids[len(ob.Bids)-1]
}

package matching

import (
	"fmt"
	"sort"
	"time"
)

const (
	COMPLETE = iota + 1
	FILLED
	TRIGGIRED

	MARKET     = "market"
	LIMIT      = "limit"
	STOPMARKET = "stop_market"
	STOPLIMIT  = "stop_limit"
)

type Order struct {
	ID     int
	UserID int
	Base   string
	Second string
	Time   time.Time
	Status int
	Type   string //market, limit, stop market, stop limit
	Side   string
	Stop   float64
	Price  float64
	Amount float64
}

type Fill struct {
	MatchOrderID int
	OrderID      int
	Price        float64
	Amount       float64
}

//BidList Bid price yüksekten düşüğe doğru sıralanacak.
type BidList []Order

func (b BidList) Len() int           { return len(b) }
func (b BidList) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b BidList) Less(i, j int) bool { return b[i].Price < b[j].Price }

//AskList Ask price düşükten yükseğe doğru sıralanacak.
type AskList []Order

func (a AskList) Len() int           { return len(a) }
func (a AskList) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a AskList) Less(i, j int) bool { return a[i].Price > a[j].Price }

type OrderBook struct {
	bids  BidList
	asks  AskList
	stops StopList
	fills []Fill
}

type StopList []Order

func (ob *OrderBook) AddOrder(order Order) {

	if order.Side == "bid" {
		ob.BidAdd(order)
	} else if order.Side == "ask" {
		ob.AskAdd(order)
	}
}

func (ob *OrderBook) AskAdd(order Order) {

	ob.asks = append(ob.asks, order)
	sort.Sort(ob.asks)
	ob.fire()
	ob.execute(order)
}

func (ob *OrderBook) BidAdd(order Order) {

	ob.bids = append(ob.bids, order)
	sort.Sort(ob.bids)
	ob.fire()
	ob.execute(order)
}

func (ob *OrderBook) AddStop(order Order) {
	ob.stops = append(ob.stops, order)
}

func (ob *OrderBook) execute(order Order) {

	orderIndex := ob.getIndex(order)

	if order.Type == MARKET {
		if order.Side == "ask" {

			fmt.Printf("Amount: %f\r\n", order.Amount)

			for i, iter := range ob.bids {

				if order.Amount >= iter.Amount {

					order.Amount -= iter.Amount
					ob.asks[orderIndex].Amount = order.Amount

					ob.bids[i].Amount = 0
					ob.bids[i].Status = COMPLETE

					ob.fills = append(ob.fills, Fill{MatchOrderID: order.ID, OrderID: iter.ID, Amount: iter.Amount, Price: iter.Price})

				} else if order.Amount < iter.Amount {

					ob.bids[i].Amount -= order.Amount
					ob.fills = append(ob.fills, Fill{MatchOrderID: order.ID, OrderID: iter.ID, Amount: order.Amount, Price: iter.Price})
					order.Amount = 0
					ob.asks[orderIndex].Amount = order.Amount
				}

				if order.Amount == 0 {
					order.Status = COMPLETE
					ob.asks[orderIndex].Status = order.Status
					break
				}
			}
		}

		if order.Side == "bid" {

			fmt.Printf("Amount: %f\r\n", order.Amount)

			for i, iter := range ob.asks {

				if order.Amount >= iter.Amount {

					order.Amount -= iter.Amount
					ob.bids[orderIndex].Amount = order.Amount

					ob.asks[i].Amount = 0
					ob.asks[i].Status = COMPLETE

					ob.fills = append(ob.fills, Fill{MatchOrderID: order.ID, OrderID: iter.ID, Amount: iter.Amount, Price: iter.Price})

				} else if order.Amount < iter.Amount {

					ob.asks[i].Amount -= order.Amount
					ob.fills = append(ob.fills, Fill{MatchOrderID: order.ID, OrderID: iter.ID, Amount: order.Amount, Price: iter.Price})
					order.Amount = 0
					ob.bids[orderIndex].Amount = order.Amount
				}

				if order.Amount == 0 {
					order.Status = COMPLETE
					ob.bids[orderIndex].Status = order.Status
					break
				}
			}
		}
	}

	if order.Type == LIMIT {

		if order.Side == "ask" {
			for i, iter := range ob.bids {

				if iter.Price > order.Price {
					continue
				}

				if order.Amount >= iter.Amount {

					order.Amount -= iter.Amount
					ob.asks[orderIndex].Amount = order.Amount
					ob.bids[i].Amount = 0
					ob.bids[i].Status = COMPLETE

					ob.fills = append(ob.fills, Fill{MatchOrderID: order.ID, OrderID: iter.ID, Amount: iter.Amount, Price: iter.Price})

				} else if order.Amount < iter.Amount {

					ob.bids[i].Amount -= order.Amount
					ob.fills = append(ob.fills, Fill{MatchOrderID: order.ID, OrderID: iter.ID, Amount: order.Amount, Price: iter.Price})
					order.Amount = 0
					ob.asks[orderIndex].Amount = order.Amount
				}

				if order.Amount == 0 {
					order.Status = COMPLETE
					ob.asks[orderIndex].Status = order.Status

					break
				}
			}
		}

		if order.Side == "bid" {
			for i, iter := range ob.asks {

				if iter.Price < order.Price {
					continue
				}

				if order.Amount >= iter.Amount {

					order.Amount -= iter.Amount
					ob.bids[orderIndex].Amount = order.Amount
					ob.asks[i].Amount = 0
					ob.asks[i].Status = COMPLETE

					ob.fills = append(ob.fills, Fill{MatchOrderID: order.ID, OrderID: iter.ID, Amount: iter.Amount, Price: iter.Price})

				} else if order.Amount < iter.Amount {

					ob.asks[i].Amount -= order.Amount
					ob.fills = append(ob.fills, Fill{MatchOrderID: order.ID, OrderID: iter.ID, Amount: order.Amount, Price: iter.Price})
					order.Amount = 0
					ob.bids[orderIndex].Amount = order.Amount
				}

				if order.Amount == 0 {
					order.Status = COMPLETE
					ob.bids[orderIndex].Status = order.Status

					break
				}
			}
		}
	}

	ob.cleanComplete()
}

func (ob *OrderBook) fire() {

	bestAsk := ob.asks[0].Price

	fmt.Printf("Best Ask (BUY): %f\r\n", bestAsk)

	//STOP MARKET SELL
	for i := 0; i < len(ob.stops); i++ {
		v := ob.stops[i]

		if v.Stop <= bestAsk {

			ob.bids = append(ob.bids, v)
			sort.Sort(ob.bids)
			ob.execute(v)

			v.Status = COMPLETE

			fmt.Printf("Triggered: %d\r\n", v.ID)
		}
	}
}

func (ob *OrderBook) cleanComplete() {

	//SELL
	for i := 0; i < len(ob.bids); i++ {
		v := ob.bids[i]

		if v.Status != COMPLETE {
			continue
		}

		ob.bids = append(ob.bids[:i], ob.bids[i+1:]...)
		i--
	}

	//ASK
	for i := 0; i < len(ob.asks); i++ {
		v := ob.asks[i]

		if v.Status != COMPLETE {
			continue
		}

		ob.asks = append(ob.asks[:i], ob.asks[i+1:]...)
		i--
	}

	//STOP
	for i := 0; i < len(ob.stops); i++ {
		v := ob.stops[i]

		if v.Status != COMPLETE {
			continue
		}

		ob.stops = append(ob.stops[:i], ob.stops[i+1:]...)
		i--
	}
}

func (ob *OrderBook) getIndex(o Order) int {

	if o.Side == "ask" {
		for i := 0; i < len(ob.asks); i++ {
			if o.ID == ob.asks[i].ID {
				return i
			}
		}
	}

	if o.Side == "bid" {
		for i := 0; i < len(ob.bids); i++ {
			if o.ID == ob.bids[i].ID {
				return i
			}
		}
	}

	return -1
}

func NewOrderBook() *OrderBook {

	return &OrderBook{
		bids:  []Order{},
		asks:  []Order{},
		stops: []Order{},
		fills: []Fill{},
	}
}

func (ob *OrderBook) Run() {

	for {

	}
}

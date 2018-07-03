package matching

import (
	"log"
	"sort"
	"time"
)

const (
	COMPLETE = iota
	FILLED

	MARKET = "market"
)

type Order struct {
	ID     int
	UserID int
	Base   string
	Second string
	Time   time.Time
	Status int
	Type   string //market, limit, stop, stop limit
	Side   string
	Price  float64
	Amount float64
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
	bids BidList
	asks AskList
}

func (ob *OrderBook) AskAdd(order Order) {

	ob.asks = append(ob.asks, order)
	sort.Sort(ob.asks)

	ob.execute(order)
}

func (ob *OrderBook) BidAdd(order Order) {

	ob.bids = append(ob.bids, order)
	sort.Sort(ob.bids)
}

func (ob *OrderBook) execute(order Order) {

	if order.Type == MARKET {
		if order.Side == "ask" {

			log.Printf("Amount: %f", order.Amount)
			var amnt float64
			for i, iter := range ob.bids {

				if order.Amount == iter.Amount {
					//Full
				} else if order.Amount > iter.Amount {
					//Filled
					amnt = iter.Amount
					order.Amount -= iter.Amount
					ob.bids[i].Amount = 0
				} else if order.Amount < iter.Amount {
					//Filled
					amnt = order.Amount
					ob.bids[i].Amount -= order.Amount
					order.Amount = 0
				}

				log.Printf("Amount: %f Bid: %f", order.Amount, amnt)

				if order.Amount == 0 {
					break
				}
			}
		}
	}

	ob.cleanComplete()
}

func (ob *OrderBook) cleanComplete() {

	for _, v := range ob.bids {
		if v.Status != 1 {
			continue
		}

		//bids'den  bu order'ı sil.

	}
}

func NewOrderBook() *OrderBook {

	return &OrderBook{
		bids: []Order{},
		asks: []Order{},
	}
}

func (ob *OrderBook) Run() {

	for {

	}
}

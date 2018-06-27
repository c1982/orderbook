package matching

import (
	"sort"
	"time"
)

const (
	COMPLETE = iota
	FILLED
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
}

func (ob *OrderBook) BidAdd(order Order) {

	ob.bids = append(ob.bids, order)
	sort.Sort(ob.bids)
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

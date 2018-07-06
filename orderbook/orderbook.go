package orderbook

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

	ASKbuy  = "ask"
	BIDsell = "bid"
)

var MakerCommFee = 0.001
var TakerCommFee = 0.002

type Order struct {
	ID      int
	UserID  int
	Base    string
	Second  string
	Time    time.Time
	Status  int
	Type    string //market, limit, stop market, stop limit
	Side    string
	Stop    float64
	Price   float64
	SAmount float64
	Easy    bool
	Amount  float64
}

type OrderBook struct {
	bids  BidList
	asks  AskList
	stops StopList
	fills []Fill
}

type Fill struct {
	BidOrder Order
	AskOrder Order
	Price    float64
	Amount   float64
	Maker    bool
	Fee      float64
	SideFee  float64
	Taker    bool
}

func (f *Fill) fee() {

}

func (f *Fill) String() {

	var commFee float64

	if f.Taker {
		commFee = TakerCommFee
		f.Fee = f.Amount * TakerCommFee
		f.SideFee = f.Amount * MakerCommFee
	} else {
		commFee = MakerCommFee
		f.Fee = f.Amount * MakerCommFee
		f.SideFee = f.Amount * MakerCommFee
	}

	fmt.Printf("Price: %f, Amount: %f (Satıcı: %d, Alıcı: %d), Fee: %f (%f), Side Fee: %f Bid: %d, Ask: %d Taker: %v\r\n",
		f.Price,
		f.Amount,
		f.BidOrder.UserID,
		f.AskOrder.UserID,
		f.Fee,
		commFee,
		f.SideFee,
		f.BidOrder.ID,
		f.AskOrder.ID,
		f.Taker)
}

type AskList []Order

type StopList []Order

//BidList Bid price yüksekten düşüğe doğru sıralanacak.
type BidList []Order

func (b BidList) Len() int           { return len(b) }
func (b BidList) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b BidList) Less(i, j int) bool { return b[i].Price < b[j].Price }

//AskList Ask price düşükten yükseğe doğru sıralanacak.

func (a AskList) Len() int           { return len(a) }
func (a AskList) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a AskList) Less(i, j int) bool { return a[i].Price > a[j].Price }

func (ob *OrderBook) AddOrder(order Order) {

	order.SAmount = order.Price * order.Amount

	if order.Side == BIDsell {
		ob.bidAdd(order)
	} else if order.Side == ASKbuy {
		ob.askAdd(order)
	}

	ob.fire() //Stoplar
}

func (ob *OrderBook) askAdd(order Order) {

	ob.asks = append(ob.asks, order)
	sort.Sort(ob.asks)
	ob.execute(order)
}

func (ob *OrderBook) bidAdd(order Order) {

	ob.bids = append(ob.bids, order)
	sort.Sort(ob.bids)
	ob.execute(order)
}

func (ob *OrderBook) AddStop(order Order) {
	ob.stops = append(ob.stops, order)
}

func (ob *OrderBook) AddFill(bid, ask Order, price, amonth float64, taker bool) {

	ob.fills = append(ob.fills, Fill{BidOrder: bid, AskOrder: ask, Price: price, Amount: amonth, Taker: taker})
}

func (ob *OrderBook) executeMarketAsk(order Order, orderIndex int) {

	if order.Easy {

		for i, iter := range ob.bids {

			if order.Amount >= iter.SAmount {

				order.Amount -= iter.SAmount
				ob.asks[orderIndex].Amount = order.Amount

				ob.bids[i].Amount = 0
				ob.bids[i].Status = COMPLETE

				ob.AddFill(iter, order, iter.Price, iter.Amount, true)

			} else if order.Amount < iter.SAmount {

				ob.bids[i].Amount -= (order.Amount / iter.Price)
				ob.AddFill(iter, order, iter.Price, (order.Amount / iter.Price), true)

				order.Amount = 0
				ob.asks[orderIndex].Amount = order.Amount
			}

			if order.Amount == 0 {
				order.Status = COMPLETE
				ob.asks[orderIndex].Status = order.Status
				break
			}
		}

	} else {

		for i, iter := range ob.bids {

			if order.Amount >= iter.Amount {

				order.Amount -= iter.Amount
				ob.asks[orderIndex].Amount = order.Amount
				ob.bids[i].Amount = 0
				ob.bids[i].Status = COMPLETE
				ob.AddFill(iter, order, iter.Price, iter.Amount, true)

			} else if order.Amount < iter.Amount {

				ob.bids[i].Amount -= order.Amount
				ob.AddFill(iter, order, iter.Price, order.Amount, true)
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

}

func (ob *OrderBook) executeMarketBid(order Order, orderIndex int) {

	fmt.Printf("Sell Amount: %f\r\n", order.Amount)

	for i, iter := range ob.asks {

		if order.Amount >= iter.Amount {

			order.Amount -= iter.Amount
			ob.bids[orderIndex].Amount = order.Amount
			ob.asks[i].Amount = 0
			ob.asks[i].Status = COMPLETE

			taker := (order.Price <= iter.Price)
			ob.AddFill(order, iter, iter.Price, iter.Amount, taker)

		} else if order.Amount < iter.Amount {

			ob.asks[i].Amount -= order.Amount

			taker := (order.Price <= iter.Price)
			ob.AddFill(order, iter, iter.Price, order.Amount, taker)

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

func (ob *OrderBook) executeLimitAsk(order Order, orderIndex int) {
	for i, iter := range ob.bids {

		if iter.Price > order.Price {
			continue
		}

		if order.Amount >= iter.Amount {

			order.Amount -= iter.Amount

			ob.asks[orderIndex].Amount = order.Amount

			ob.bids[i].Amount = 0
			ob.bids[i].Status = COMPLETE

			taker := (order.Price <= iter.Price)
			ob.AddFill(iter, order, iter.Price, iter.Amount, taker)

		} else if order.Amount < iter.Amount {

			ob.bids[i].Amount -= order.Amount

			taker := (order.Price <= iter.Price)
			ob.AddFill(iter, order, iter.Price, order.Amount, taker)

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

func (ob *OrderBook) executeLimitBid(order Order, orderIndex int) {
	for i, iter := range ob.asks {

		if iter.Price < order.Price {
			continue
		}

		if order.Amount >= iter.Amount {

			order.Amount -= iter.Amount
			ob.bids[orderIndex].Amount = order.Amount
			ob.asks[i].Amount = 0
			ob.asks[i].Status = COMPLETE

			ob.AddFill(order, iter, iter.Price, iter.Amount, false)

		} else if order.Amount < iter.Amount {

			ob.asks[i].Amount -= order.Amount
			ob.AddFill(order, iter, iter.Price, order.Amount, true)
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

func (ob *OrderBook) execute(order Order) {

	orderIndex := ob.getIndex(order)

	logic := map[string]map[string]func(){
		MARKET: map[string]func(){
			ASKbuy:  func() { ob.executeMarketAsk(order, orderIndex) },
			BIDsell: func() { ob.executeMarketBid(order, orderIndex) },
		},
		LIMIT: map[string]func(){
			ASKbuy:  func() { ob.executeLimitAsk(order, orderIndex) },
			BIDsell: func() { ob.executeLimitBid(order, orderIndex) },
		},
	}

	logic[order.Type][order.Side]()

	ob.cleanComplete()
}

func (ob *OrderBook) fire() {

	var bestAsk float64
	var bestBid float64

	if len(ob.asks) > 0 {
		bestAsk = ob.asks[0].Price
	}

	if len(ob.bids) > 0 {
		bestBid = ob.bids[0].Price
	}

	for i := 0; i < len(ob.stops); i++ {
		v := ob.stops[i]

		if v.Side == ASKbuy {

			if v.Stop >= bestAsk {

				ob.askAdd(v)
				v.Status = COMPLETE

				fmt.Printf("Triggered: %d\r\n", v.ID)
			}

		} else if v.Side == BIDsell {

			if v.Stop <= bestBid {

				ob.bidAdd(v)
				v.Status = COMPLETE

				fmt.Printf("Triggered: %d\r\n", v.ID)
			}
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

	if o.Side == ASKbuy {
		for i := 0; i < len(ob.asks); i++ {
			if o.ID == ob.asks[i].ID {
				return i
			}
		}
	}

	if o.Side == BIDsell {
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

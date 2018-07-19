package orderbook

import (
	"encoding/json"
	"fmt"
	"sort"
	"time"

	"github.com/shopspring/decimal"
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

var MakerCommFee = decimal.New(1, -10)
var TakerCommFee = decimal.New(2, -10)
var ZeroNew = decimal.New(0, -10)

type Order struct {
	ID      int64           `json:"id"`
	UserID  int             `json:"-"`
	Base    string          `json:"base"`
	Second  string          `json:"second"`
	Time    time.Time       `json:"time"`
	Status  int             `json:"status"`
	Type    string          `json:"type"` //market, limit, stop market, stop limit
	Side    string          `json:"side"`
	Stop    decimal.Decimal `json:"stop"`
	Price   decimal.Decimal `json:"price"`
	SAmount decimal.Decimal `json:"-"`
	Easy    bool            `json:"easy"`
	Amount  decimal.Decimal `json:"amount"`
}

type OrderBook struct {
	bids  BidList
	asks  AskList
	stops StopList
	fills FillList
}

type Fill struct {
	BidOrder Order           `json:"-"`
	AskOrder Order           `json:"-"`
	Time     time.Time       `json:"time"`
	Price    decimal.Decimal `json:"price"`
	Amount   decimal.Decimal `json:"amount"`
	Maker    bool            `json:"maker"`
	Fee      decimal.Decimal `json:"fee"`
	SideFee  decimal.Decimal `json:"sidefee"`
	Taker    bool            `json:"taker"`
}

func (f *Fill) String() {

	var commFee = decimal.New(0, -10)

	if f.Taker {
		commFee = TakerCommFee
		f.Fee = f.Amount.Mul(TakerCommFee)
		f.SideFee = f.Amount.Mul(MakerCommFee)
	} else {
		commFee = MakerCommFee
		f.Fee = f.Amount.Mul(MakerCommFee)
		f.SideFee = f.Amount.Mul(MakerCommFee)
	}

	fmt.Printf("Price: %s, Amount: %s (Satıcı: %d, Alıcı: %d), Fee: %s (%s), Side Fee: %s Bid: %d, Ask: %d Taker: %v\r\n",
		f.Price.StringFixed(8),
		f.Amount.StringFixed(8),
		f.BidOrder.UserID,
		f.AskOrder.UserID,
		f.Fee.String(),
		commFee.String(),
		f.SideFee.String(),
		f.BidOrder.ID,
		f.AskOrder.ID,
		f.Taker)
}

type FillList []Fill

type AskList []Order

type StopList []Order

//BidList Bid price yüksekten düşüğe doğru sıralanacak.
type BidList []Order

func (b BidList) Len() int      { return len(b) }
func (b BidList) Swap(i, j int) { b[i], b[j] = b[j], b[i] }

//func (b BidList) Less(i, j int) bool { return b[i].Price < b[j].Price }
func (b BidList) Less(i, j int) bool { return b[i].Price.Cmp(b[j].Price) == -1 }

//AskList Ask price düşükten yükseğe doğru sıralanacak.

func (a AskList) Len() int      { return len(a) }
func (a AskList) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

//func (a AskList) Less(i, j int) bool { return a[i].Price > a[j].Price }
func (a AskList) Less(i, j int) bool { return a[i].Price.Cmp(a[j].Price) == 1 }

//Filed book sıralaması tarihe göre yapılır.
func (f FillList) Len() int           { return len(f) }
func (f FillList) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }
func (f FillList) Less(i, j int) bool { return f[i].Time.After(f[j].Time) }

func (ob *OrderBook) AddOrder(order Order) {

	//order.SAmount = order.Price * order.Amount
	order.SAmount = decimal.New(0, -10)
	//order.SAmount = order.SAmount.Mul(order.Price, order.Amount)
	order.SAmount = order.Price.Mul(order.Amount)

	if order.Side == BIDsell {
		ob.bidAdd(order)
	} else if order.Side == ASKbuy {
		ob.askAdd(order)
	}

	ob.fire()
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

//AddFill ...
func (ob *OrderBook) AddFill(bid, ask Order, price, amonth decimal.Decimal, taker bool) {

	ob.fills = append(ob.fills, Fill{BidOrder: bid, AskOrder: ask, Price: price, Amount: amonth, Taker: taker, Time: time.Now()})
	sort.Sort(ob.fills)
}

func (ob *OrderBook) ToList() []byte {

	export := struct {
		Bids  BidList  `json:"bids"`
		Asks  AskList  `json:"asks"`
		Fills FillList `json:"fills"`
	}{
		Bids:  ob.bids,
		Asks:  ob.asks,
		Fills: ob.fills,
	}

	data, err := json.Marshal(export)

	if err != nil {
		return nil
	}

	return data
}

func (ob *OrderBook) executeMarketAsk(order Order, orderIndex int) {

	if order.Easy {

		for i, iter := range ob.bids {

			//if order.Amount >= iter.SAmount {
			if order.Amount.Cmp(iter.SAmount) == 1 || order.Amount.Cmp(iter.SAmount) == 0 {

				//order.Amount -= iter.SAmount
				//order.Amount = order.Amount.Sub(order.Amount, iter.SAmount)
				order.Amount = order.Amount.Sub(iter.SAmount)
				ob.asks[orderIndex].Amount = order.Amount

				ob.bids[i].Amount = decimal.New(0, -10)
				ob.bids[i].Status = COMPLETE

				ob.AddFill(iter, order, iter.Price, iter.Amount, true)

				//} else if order.Amount < iter.SAmount {
			} else if order.Amount.Cmp(iter.SAmount) == -1 {

				sprice := order.Amount.Div(iter.Price)
				ob.bids[i].Amount = ob.bids[i].Amount.Sub(sprice)
				ob.AddFill(iter, order, iter.Price, sprice, true)

				order.Amount = ZeroNew
				ob.asks[orderIndex].Amount = order.Amount
			}

			//if order.Amount.Cmp(big.NewInt(0)) == 0 {
			if order.Amount.Cmp(ZeroNew) == 0 {
				order.Status = COMPLETE
				ob.asks[orderIndex].Status = COMPLETE
				break
			}
		}

	} else {

		for i, iter := range ob.bids {

			//if order.Amount >= iter.Amount {
			if order.Amount.Cmp(iter.Amount) == 1 || order.Amount.Cmp(iter.Amount) == 0 {

				//order.Amount -= iter.Amount
				//order.Amount = order.Amount.Sub(order.Amount, iter.Amount)
				order.Amount = order.Amount.Sub(iter.Amount)
				ob.asks[orderIndex].Amount = order.Amount
				//ob.bids[i].Amount = 0
				ob.bids[i].Amount = ZeroNew
				ob.bids[i].Status = COMPLETE
				ob.AddFill(iter, order, iter.Price, iter.Amount, true)

				//} else if order.Amount < iter.Amount {
			} else if order.Amount.Cmp(iter.Amount) == -1 {

				//ob.bids[i].Amount -= order.Amount
				//ob.bids[i].Amount = ob.bids[i].Amount.Sub(ob.bids[i].Amount, order.Amount)
				ob.bids[i].Amount = ob.bids[i].Amount.Sub(order.Amount)
				ob.AddFill(iter, order, iter.Price, order.Amount, true)
				order.Amount = ZeroNew
				ob.asks[orderIndex].Amount = order.Amount
			}

			//if order.Amount == 0 {
			if order.Amount.Cmp(ZeroNew) == 0 {
				order.Status = COMPLETE
				ob.asks[orderIndex].Status = COMPLETE
				break
			}
		}
	}

}

func (ob *OrderBook) executeMarketBid(order Order, orderIndex int) {

	for i, iter := range ob.asks {

		//if order.Amount >= iter.Amount {
		if order.Amount.Cmp(iter.Amount) == 1 || order.Amount.Cmp(iter.Amount) == 0 {

			//order.Amount -= iter.Amount
			//order.Amount = order.Amount.Sub(order.Amount, iter.Amount)
			order.Amount = order.Amount.Sub(iter.Amount)

			ob.bids[orderIndex].Amount = order.Amount
			ob.asks[i].Amount = ZeroNew
			ob.asks[i].Status = COMPLETE

			//taker := (order.Price <= iter.Price)
			taker := (order.Price.Cmp(iter.Price) == -1) || (order.Price.Cmp(iter.Price) == 0)
			ob.AddFill(order, iter, iter.Price, iter.Amount, taker)

			//} else if order.Amount < iter.Amount {
		} else if order.Amount.Cmp(iter.Amount) == -1 {

			//ob.asks[i].Amount -= order.Amount
			//ob.asks[i].Amount = ob.asks[i].Amount.Sub(ob.asks[i].Amount, order.Amount)
			ob.asks[i].Amount = ob.asks[i].Amount.Sub(order.Amount)

			//taker := (order.Price <= iter.Price)
			taker := (order.Price.Cmp(iter.Price) == -1) || (order.Price.Cmp(iter.Price) == 0)

			ob.AddFill(order, iter, iter.Price, order.Amount, taker)

			order.Amount = ZeroNew
			ob.bids[orderIndex].Amount = order.Amount
		}

		//if order.Amount == 0 {
		if order.Amount.Cmp(ZeroNew) == 0 {
			order.Status = COMPLETE
			ob.bids[orderIndex].Status = order.Status
			break
		}
	}

}

func (ob *OrderBook) executeLimitAsk(order Order, orderIndex int) {
	for i, iter := range ob.bids {

		//if iter.Price > order.Price {
		if iter.Price.Cmp(order.Price) == 1 {
			continue
		}

		//if order.Amount >= iter.Amount {
		if order.Amount.Cmp(iter.Amount) == 1 || order.Amount.Cmp(iter.Amount) == 0 {

			//order.Amount -= iter.Amount
			order.Amount = order.Amount.Sub(iter.Amount)

			ob.asks[orderIndex].Amount = order.Amount

			ob.bids[i].Amount = ZeroNew
			ob.bids[i].Status = COMPLETE

			//taker := (order.Price <= iter.Price)
			taker := (order.Price.Cmp(iter.Price) == -1) || (order.Price.Cmp(iter.Price) == 0)
			ob.AddFill(iter, order, iter.Price, iter.Amount, taker)

			//} else if order.Amount < iter.Amount {
		} else if order.Amount.Cmp(iter.Amount) == -1 {

			//ob.bids[i].Amount -= order.Amount
			//ob.bids[i].Amount = ob.bids[i].Amount.Sub(ob.bids[i].Amount, order.Amount)
			ob.bids[i].Amount = ob.bids[i].Amount.Sub(order.Amount)

			//taker := (order.Price <= iter.Price)
			taker := (order.Price.Cmp(iter.Price) == -1) || (order.Price.Cmp(iter.Price) == 0)
			ob.AddFill(iter, order, iter.Price, order.Amount, taker)

			order.Amount = ZeroNew
			ob.asks[orderIndex].Amount = order.Amount
		}

		//if order.Amount == 0 {
		if order.Amount.Cmp(ZeroNew) == 0 {
			order.Status = COMPLETE
			ob.asks[orderIndex].Status = order.Status

			break
		}
	}
}

func (ob *OrderBook) executeLimitBid(order Order, orderIndex int) {
	for i, iter := range ob.asks {

		//if iter.Price < order.Price {
		if iter.Price.Cmp(order.Price) == -1 {
			continue
		}

		//if order.Amount >= iter.Amount {
		if order.Amount.Cmp(iter.Amount) == 1 || order.Amount.Cmp(iter.Amount) == 0 {

			//order.Amount -= iter.Amount
			order.Amount = order.Amount.Sub(iter.Amount)

			ob.bids[orderIndex].Amount = order.Amount
			ob.asks[i].Amount = ZeroNew
			ob.asks[i].Status = COMPLETE

			ob.AddFill(order, iter, iter.Price, iter.Amount, false)

			//} else if order.Amount < iter.Amount {
		} else if order.Amount.Cmp(iter.Amount) == -1 {

			//ob.asks[i].Amount -= order.Amount
			ob.asks[i].Amount = ob.asks[i].Amount.Sub(order.Amount)

			ob.AddFill(order, iter, iter.Price, order.Amount, true)
			order.Amount = ZeroNew
			ob.bids[orderIndex].Amount = order.Amount
		}

		//if order.Amount == 0 {
		if order.Amount.Cmp(ZeroNew) == 0 {
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

	var bestAsk decimal.Decimal
	var bestBid decimal.Decimal

	if len(ob.asks) > 0 {
		bestAsk = ob.asks[0].Price
	}

	if len(ob.bids) > 0 {
		bestBid = ob.bids[0].Price
	}

	for i := 0; i < len(ob.stops); i++ {
		v := ob.stops[i]

		//TODO: Eşitlik?
		if v.Side == ASKbuy {

			//if v.Stop >= bestAsk {
			if v.Stop.Cmp(bestAsk) == 1 || v.Stop.Cmp(bestAsk) == 0 {

				ob.stops[i].Status = COMPLETE
				ob.askAdd(v)
				fmt.Printf("Triggered: %d (Side:%s)\r\n", v.ID, v.Side)
			}

		} else if v.Side == BIDsell {

			//if v.Stop <= bestBid {
			if v.Stop.Cmp(bestBid) == -1 || v.Stop.Cmp(bestBid) == 0 {

				ob.stops[i].Status = COMPLETE
				ob.bidAdd(v)

				fmt.Printf("Triggered: %d (Side:%s)\r\n", v.ID, v.Side)
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

func (ob *OrderBook) Debug() {
	fmt.Println("FILLED-BOOK")

	for i := 0; i < len(ob.fills); i++ {
		f := ob.fills[i]
		f.String()
	}

	fmt.Println("BUYS:")

	if len(ob.asks) == 0 {
		fmt.Println("empty")
	}
	for i := 0; i < len(ob.asks); i++ {
		fmt.Printf("ID: %d, Amonth: %s (%s)\r\n", ob.asks[i].ID, ob.asks[i].Amount.StringFixed(8), ob.asks[i].Price.StringFixed(8))
	}

	fmt.Println("SELLS:")

	if len(ob.bids) == 0 {
		fmt.Println("empty")
	}
	for i := 0; i < len(ob.bids); i++ {
		fmt.Printf("ID: %d, Amonth: %s (%s)\r\n", ob.bids[i].ID, ob.bids[i].Amount.StringFixed(8), ob.bids[i].Price.StringFixed(8))
	}

	fmt.Println("STOPS:")

	if len(ob.stops) == 0 {
		fmt.Println("empty")
	}
	for i := 0; i < len(ob.stops); i++ {
		fmt.Printf("ID: %d, Stop: %s, Amonth: %s (%s)\r\n", ob.stops[i].ID, ob.stops[i].Stop.StringFixed(8), ob.stops[i].Amount.StringFixed(8), ob.stops[i].Price.StringFixed(8))
	}
}

func NewOrderBook() *OrderBook {

	return &OrderBook{
		bids:  []Order{},
		asks:  []Order{},
		stops: []Order{},
		fills: []Fill{},
	}
}

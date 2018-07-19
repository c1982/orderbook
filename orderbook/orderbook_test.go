package orderbook

import (
	"testing"
	"time"

	"github.com/shopspring/decimal"
)

//Osijen
//Eşcenc
//Mayış
//Bitgoing
//Sikorta
//Laylon
//Eşki

func LoadTestData(ob *OrderBook) {
	ob.AddOrder(Order{ID: 1, UserID: 100, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: NewAmount("1.1"), Price: NewAmount("30500"), Time: time.Now()})
	ob.AddOrder(Order{ID: 2, UserID: 101, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: NewAmount("0.20"), Price: NewAmount("30300"), Time: time.Now()})
	ob.AddOrder(Order{ID: 3, UserID: 102, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: NewAmount("0.8"), Price: NewAmount("30250"), Time: time.Now()})

	ob.AddOrder(Order{ID: 4, UserID: 104, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: NewAmount("0.90"), Price: NewAmount("30600"), Time: time.Now()})
	ob.AddOrder(Order{ID: 5, UserID: 105, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: NewAmount("0.75"), Price: NewAmount("30700"), Time: time.Now()})
	ob.AddOrder(Order{ID: 6, UserID: 107, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: NewAmount("0.20"), Price: NewAmount("31000"), Time: time.Now()})
}

func NewAmount(value string) decimal.Decimal {
	d, _ := decimal.NewFromString(value)
	return d
}

func TestSortingOrderBook(t *testing.T) {

	ob := NewOrderBook()

	//BUY
	ob.AddOrder(Order{ID: 1, UserID: 100, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: NewAmount("1.0"), Price: NewAmount("6000"), Time: time.Now()})
	ob.AddOrder(Order{ID: 2, UserID: 101, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: NewAmount("0.90"), Price: NewAmount("6001"), Time: time.Now()})
	ob.AddOrder(Order{ID: 3, UserID: 102, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: NewAmount("1.01"), Price: NewAmount("6000"), Time: time.Now()})
	ob.AddOrder(Order{ID: 4, UserID: 103, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: NewAmount("1.05"), Price: NewAmount("5999"), Time: time.Now()})

	//SELL
	ob.AddOrder(Order{ID: 5, UserID: 104, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: NewAmount("1"), Price: NewAmount("6002"), Time: time.Now()})
	ob.AddOrder(Order{ID: 6, UserID: 105, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: NewAmount("2"), Price: NewAmount("6003"), Time: time.Now()})
	ob.AddOrder(Order{ID: 8, UserID: 107, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: NewAmount("1.02"), Price: NewAmount("6002.1"), Time: time.Now()})
	ob.AddOrder(Order{ID: 7, UserID: 106, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: NewAmount("1.02"), Price: NewAmount("6004"), Time: time.Now()})

	ob.Debug()
}

func TestBuyMarket(t *testing.T) {

	ob := NewOrderBook()

	LoadTestData(ob)

	ob.AddOrder(Order{ID: 9, UserID: 102, Base: "BTC", Second: "TRY", Type: "market", Side: "ask", Amount: NewAmount("10000"), Easy: true, Price: NewAmount("0"), Time: time.Now()})
	ob.AddOrder(Order{ID: 10, UserID: 102, Base: "BTC", Second: "TRY", Type: "market", Side: "ask", Amount: NewAmount("10000"), Easy: true, Price: NewAmount("0"), Time: time.Now()})
	ob.AddOrder(Order{ID: 11, UserID: 102, Base: "BTC", Second: "TRY", Type: "market", Side: "ask", Amount: NewAmount("10000"), Easy: true, Price: NewAmount("0"), Time: time.Now()})
	ob.Debug()
}

/*
func TestBuyLimitOrder(t *testing.T) {

	ob := NewOrderBook()

	//BUY
	ob.AddOrder(Order{ID: 3, UserID: 102, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 0.8, Price: 30250, Time: time.Now()})
	ob.AddOrder(Order{ID: 1, UserID: 100, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 1.1, Price: 30500, Time: time.Now()})
	ob.AddOrder(Order{ID: 2, UserID: 101, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 0.20, Price: 30300, Time: time.Now()})

	//SELL
	ob.AddOrder(Order{ID: 4, UserID: 104, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 0.90, Price: 30600, Time: time.Now()})
	ob.AddOrder(Order{ID: 5, UserID: 105, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 0.75, Price: 30700, Time: time.Now()})
	ob.AddOrder(Order{ID: 6, UserID: 107, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 0.20, Price: 31000, Time: time.Now()})

	//BUY
	ob.AddOrder(Order{ID: 9, UserID: 102, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 2.0, Price: 30600, Time: time.Now()})

	//SELL
	ob.AddOrder(Order{ID: 10, UserID: 203, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 1.10, Price: 30600, Time: time.Now()})

	ob.Debug()
}

func TestBuyLimitOrderPartial(t *testing.T) {

	ob := NewOrderBook()

	//BUY
	ob.AddOrder(Order{ID: 1, UserID: 100, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 1.1, Price: 30500, Time: time.Now()})
	ob.AddOrder(Order{ID: 2, UserID: 101, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 0.20, Price: 30300, Time: time.Now()})
	ob.AddOrder(Order{ID: 3, UserID: 102, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 0.8, Price: 30250, Time: time.Now()})

	//SELL
	ob.AddOrder(Order{ID: 4, UserID: 104, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 1.0, Price: 30600, Time: time.Now()})
	ob.AddOrder(Order{ID: 5, UserID: 105, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 0.75, Price: 30700, Time: time.Now()})
	ob.AddOrder(Order{ID: 6, UserID: 107, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 0.20, Price: 31000, Time: time.Now()})

	ob.AddOrder(Order{ID: 9, UserID: 102, Base: "BTC", Second: "TRY", Type: "market", Side: "ask", Amount: 30600, Price: 0, Time: time.Now()})

	ob.Debug()
}

func TestSellOrderSenaryoA(t *testing.T) {

	ob := NewOrderBook()

	//BUY
	ob.AddOrder(Order{ID: 1, UserID: 100, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 1.1, Price: 30500, Time: time.Now()})
	ob.AddOrder(Order{ID: 2, UserID: 101, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 0.20, Price: 30300, Time: time.Now()})
	ob.AddOrder(Order{ID: 3, UserID: 102, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 0.8, Price: 30250, Time: time.Now()})

	//SELL
	ob.AddOrder(Order{ID: 4, UserID: 104, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 0.90, Price: 30600, Time: time.Now()})
	ob.AddOrder(Order{ID: 5, UserID: 105, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 0.75, Price: 30700, Time: time.Now()})
	ob.AddOrder(Order{ID: 6, UserID: 107, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 0.20, Price: 31000, Time: time.Now()})

	//Orderbook'da istenilen price hazır bir şekilde varsa direkt işlem gerçekleşeceği için Taker olur.
	ob.AddOrder(Order{ID: 7, UserID: 102, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 1.1, Price: 30500, Time: time.Now()})

	//Orderbook'da istediği price olmadığı için mecvur bekleyecek. O nedenke maker olacak.
	//ob.AddOrder(Order{ID: 7, UserID: 102, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 0.70, Price: 30501, Time: time.Now()})

	ob.Debug()

}

func TestSellOrderSenaryoB(t *testing.T) {

	ob := NewOrderBook()

	//SELL
	ob.AddOrder(Order{ID: 4, UserID: 104, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 0.90, Price: 30600, Time: time.Now()})
	ob.AddOrder(Order{ID: 5, UserID: 105, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 0.75, Price: 30700, Time: time.Now()})
	ob.AddOrder(Order{ID: 6, UserID: 107, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 0.20, Price: 31000, Time: time.Now()})

	//BUY
	ob.AddOrder(Order{ID: 1, UserID: 100, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 1.1, Price: 30500, Time: time.Now()})
	ob.AddOrder(Order{ID: 2, UserID: 101, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 0.20, Price: 30300, Time: time.Now()})
	ob.AddOrder(Order{ID: 3, UserID: 102, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 0.80, Price: 30250, Time: time.Now()})

	ob.AddOrder(Order{ID: 7, UserID: 102, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 1.50, Price: 30500, Time: time.Now()})

	ob.Debug()
}

func TestSellOrderSenaryoC(t *testing.T) {

	ob := NewOrderBook()

	//SELL
	ob.AddOrder(Order{ID: 4, UserID: 104, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 0.90, Price: 30600, Time: time.Now()})
	ob.AddOrder(Order{ID: 5, UserID: 105, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 0.75, Price: 30700, Time: time.Now()})
	ob.AddOrder(Order{ID: 6, UserID: 107, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 0.20, Price: 31000, Time: time.Now()})

	//BUY
	ob.AddOrder(Order{ID: 1, UserID: 100, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 1.1, Price: 30500, Time: time.Now()})
	ob.AddOrder(Order{ID: 2, UserID: 101, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 0.20, Price: 30300, Time: time.Now()})
	ob.AddOrder(Order{ID: 3, UserID: 102, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 0.80, Price: 30250, Time: time.Now()})

	ob.AddOrder(Order{ID: 7, UserID: 102, Base: "BTC", Second: "TRY", Type: "market", Side: "bid", Amount: 1.20, Price: 0, Time: time.Now()})

	ob.Debug()

}

func TestSellOrderSenaryoD(t *testing.T) {

	ob := NewOrderBook()

	//SELL
	ob.AddOrder(Order{ID: 4, UserID: 104, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 0.90, Price: 30600, Time: time.Now()})
	ob.AddOrder(Order{ID: 5, UserID: 105, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 0.75, Price: 30700, Time: time.Now()})
	ob.AddOrder(Order{ID: 6, UserID: 107, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 0.20, Price: 31000, Time: time.Now()})

	//BUY
	ob.AddOrder(Order{ID: 1, UserID: 100, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 1.1, Price: 30500, Time: time.Now()})
	ob.AddOrder(Order{ID: 2, UserID: 101, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 0.20, Price: 30300, Time: time.Now()})
	ob.AddOrder(Order{ID: 3, UserID: 102, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 0.80, Price: 30250, Time: time.Now()})

	ob.AddOrder(Order{ID: 7, UserID: 102, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 1, Price: 30550, Time: time.Now()})

	ob.Debug()

}

func TestStopMarketSellSenaryoA(t *testing.T) {

	ob := NewOrderBook()

	//BUY
	ob.AddOrder(Order{ID: 1, UserID: 100, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 1.1, Price: 30500, Time: time.Now()})
	ob.AddOrder(Order{ID: 2, UserID: 101, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 0.20, Price: 30300, Time: time.Now()})
	ob.AddOrder(Order{ID: 3, UserID: 102, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 0.80, Price: 30250, Time: time.Now()})
	ob.AddOrder(Order{ID: 9, UserID: 102, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 0.80, Price: 30250, Time: time.Now()})

	//SELL
	ob.AddOrder(Order{ID: 4, UserID: 104, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 0.90, Price: 30600, Time: time.Now()})
	ob.AddOrder(Order{ID: 5, UserID: 105, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 0.75, Price: 30700, Time: time.Now()})
	ob.AddOrder(Order{ID: 6, UserID: 107, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 0.20, Price: 31000, Time: time.Now()})

	//STOP: 30.400
	ob.AddStop(Order{ID: 7, UserID: 104, Base: "BTC", Second: "TRY", Type: "market", Side: "bid", Stop: 30400, Amount: 1.01, Price: 0, Time: time.Now()})
	ob.AddStop(Order{ID: 10, UserID: 104, Base: "BTC", Second: "TRY", Type: "market", Side: "bid", Stop: 30400, Amount: 0.5, Price: 0, Time: time.Now()})

	ob.AddOrder(Order{ID: 8, UserID: 100, Base: "BTC", Second: "TRY", Type: "market", Side: "bid", Amount: 1.1, Price: 0, Time: time.Now()})

	ob.Debug()
}

func TestStopMarketSellSenaryoB(t *testing.T) {

	ob := NewOrderBook()

	//BUY
	ob.AddOrder(Order{ID: 1, UserID: 100, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 1.1, Price: 30500, Time: time.Now()})
	ob.AddOrder(Order{ID: 2, UserID: 101, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 0.20, Price: 30300, Time: time.Now()})
	ob.AddOrder(Order{ID: 3, UserID: 102, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 0.80, Price: 30250, Time: time.Now()})
	ob.AddOrder(Order{ID: 9, UserID: 102, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 0.80, Price: 30250, Time: time.Now()})

	//SELL
	ob.AddOrder(Order{ID: 4, UserID: 104, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 0.90, Price: 30600, Time: time.Now()})
	ob.AddOrder(Order{ID: 5, UserID: 105, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 0.75, Price: 30700, Time: time.Now()})
	ob.AddOrder(Order{ID: 6, UserID: 107, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 0.20, Price: 31000, Time: time.Now()})

	ob.AddStop(Order{ID: 7, UserID: 104, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Stop: 30400, Amount: 1.0, Price: 30300, Time: time.Now()})

	ob.AddOrder(Order{ID: 8, UserID: 100, Base: "BTC", Second: "TRY", Type: "market", Side: "bid", Amount: 1.1, Price: 0, Time: time.Now()})

	ob.Debug()
}

func TestStopMarketBUY(t *testing.T) {

	ob := NewOrderBook()

	//BUY
	ob.AddOrder(Order{ID: 1, UserID: 100, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 1.1, Price: 30500, Time: time.Now()})
	ob.AddOrder(Order{ID: 2, UserID: 101, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 0.20, Price: 30300, Time: time.Now()})
	ob.AddOrder(Order{ID: 3, UserID: 102, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 0.80, Price: 30250, Time: time.Now()})

	//SELL
	ob.AddOrder(Order{ID: 4, UserID: 104, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 1.10, Price: 30600, Time: time.Now()})
	ob.AddOrder(Order{ID: 5, UserID: 105, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 0.75, Price: 30700, Time: time.Now()})
	ob.AddOrder(Order{ID: 6, UserID: 107, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 0.20, Price: 31000, Time: time.Now()})

	//STOP
	ob.AddStop(Order{ID: 7, UserID: 104, Base: "BTC", Second: "TRY", Type: "market", Side: "ask", Stop: 30900, Amount: 1, Price: 0, Easy: true, Time: time.Now()})

	ob.AddOrder(Order{ID: 8, UserID: 100, Base: "BTC", Second: "TRY", Type: "market", Side: "ask", Amount: 1.1, Price: 0, Time: time.Now()})

	ob.Debug()
}

func TestStoSellLimit(t *testing.T) {

	ob := NewOrderBook()

	//BUY
	ob.AddOrder(Order{ID: 1, UserID: 100, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 1.1, Price: 30500, Time: time.Now()})
	ob.AddOrder(Order{ID: 2, UserID: 101, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 0.20, Price: 30300, Time: time.Now()})
	ob.AddOrder(Order{ID: 3, UserID: 102, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 0.80, Price: 30250, Time: time.Now()})
	ob.AddOrder(Order{ID: 4, UserID: 102, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 0.80, Price: 30250, Time: time.Now()})

	//SELL
	ob.AddOrder(Order{ID: 5, UserID: 104, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 0.90, Price: 30600, Time: time.Now()})
	ob.AddOrder(Order{ID: 6, UserID: 105, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 0.75, Price: 30610, Time: time.Now()})
	ob.AddOrder(Order{ID: 7, UserID: 107, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 0.20, Price: 31000, Time: time.Now()})

	ob.AddStop(Order{ID: 8, UserID: 106, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 0.3, Price: 30650, Stop: 30610, Time: time.Now()})
	ob.Debug()
	ob.AddOrder(Order{ID: 9, UserID: 108, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 1.0, Price: 30620, Time: time.Now()})
	ob.Debug()
}

func BenchmarkMarket(b *testing.B) {

	ob := NewOrderBook()

	for n := 0; n < b.N; n++ {
		ob.AddOrder(Order{ID: 1, UserID: 100, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 1.1, Price: 30500, Time: time.Now()})
		ob.AddOrder(Order{ID: 2, UserID: 104, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 1.10, Price: 30600, Time: time.Now()})
	}

}
*/

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
//Enşe
//kantır sıkrayt
//Gogıl

func LoadTestData(ob *OrderBook) {

	//BUYS
	ob.AddOrder(Order{ID: 1, UserID: 100, Base: "BTC", Second: "TRY", Type: "limit", Side: BUY, Amount: A("1.1"), Price: A("30500"), Time: time.Now()})
	ob.AddOrder(Order{ID: 2, UserID: 101, Base: "BTC", Second: "TRY", Type: "limit", Side: BUY, Amount: A("0.20"), Price: A("30300"), Time: time.Now()})
	ob.AddOrder(Order{ID: 3, UserID: 102, Base: "BTC", Second: "TRY", Type: "limit", Side: BUY, Amount: A("0.8"), Price: A("30250"), Time: time.Now()})

	//SELLS
	ob.AddOrder(Order{ID: 4, UserID: 104, Base: "BTC", Second: "TRY", Type: "limit", Side: SELL, Amount: A("0.90"), Price: A("30600"), Time: time.Now()})
	ob.AddOrder(Order{ID: 5, UserID: 105, Base: "BTC", Second: "TRY", Type: "limit", Side: SELL, Amount: A("0.75"), Price: A("30700"), Time: time.Now()})
	ob.AddOrder(Order{ID: 6, UserID: 107, Base: "BTC", Second: "TRY", Type: "limit", Side: SELL, Amount: A("0.20"), Price: A("31000"), Time: time.Now()})
}

func A(value string) decimal.Decimal {
	d, _ := decimal.NewFromString(value)
	return d
}

func TestSortingOrderBook(t *testing.T) {

	ob := NewOrderBook()

	//BUY
	ob.AddOrder(Order{ID: 1, UserID: 100, Base: "BTC", Second: "TRY", Type: "limit", Side: BUY, Amount: A("1.0"), Price: A("6000"), Time: time.Now()})
	ob.AddOrder(Order{ID: 2, UserID: 101, Base: "BTC", Second: "TRY", Type: "limit", Side: BUY, Amount: A("0.90"), Price: A("6001"), Time: time.Now()})
	ob.AddOrder(Order{ID: 3, UserID: 102, Base: "BTC", Second: "TRY", Type: "limit", Side: BUY, Amount: A("1.01"), Price: A("6000"), Time: time.Now()})
	ob.AddOrder(Order{ID: 4, UserID: 103, Base: "BTC", Second: "TRY", Type: "limit", Side: BUY, Amount: A("1.05"), Price: A("5999"), Time: time.Now()})

	//SELL
	ob.AddOrder(Order{ID: 5, UserID: 104, Base: "BTC", Second: "TRY", Type: "limit", Side: SELL, Amount: A("1"), Price: A("6002"), Time: time.Now()})
	ob.AddOrder(Order{ID: 6, UserID: 105, Base: "BTC", Second: "TRY", Type: "limit", Side: SELL, Amount: A("2"), Price: A("6003"), Time: time.Now()})
	ob.AddOrder(Order{ID: 8, UserID: 107, Base: "BTC", Second: "TRY", Type: "limit", Side: SELL, Amount: A("1.02"), Price: A("6002.1"), Time: time.Now()})
	ob.AddOrder(Order{ID: 7, UserID: 106, Base: "BTC", Second: "TRY", Type: "limit", Side: SELL, Amount: A("1.02"), Price: A("6004"), Time: time.Now()})

	ob.Debug()
}

func Test_Market_Buy_Easy(t *testing.T) {

	ob := NewOrderBook()

	LoadTestData(ob)

	ob.AddOrder(Order{ID: 9, UserID: 102, Base: "BTC", Second: "TRY", Type: "market", Side: BUY, Amount: A("10000"), Easy: true, Price: A("0"), Time: time.Now()})
	ob.AddOrder(Order{ID: 10, UserID: 102, Base: "BTC", Second: "TRY", Type: "market", Side: BUY, Amount: A("10000"), Easy: true, Price: A("0"), Time: time.Now()})
	//ob.AddOrder(Order{ID: 11, UserID: 102, Base: "BTC", Second: "TRY", Type: "market", Side: BUY, Amount: A("10000"), Easy: true, Price: A("0"), Time: time.Now()})
	ob.Debug()

}

func Test_Market_Buy_Partial(t *testing.T) {

	ob := NewOrderBook()
	LoadTestData(ob)

	ob.AddOrder(Order{ID: 9, UserID: 102, Base: "BTC", Second: "TRY", Type: "market", Side: BUY, Amount: A("45000"), Easy: true, Price: A("0"), Time: time.Now()})

	ob.Debug()
}

func Test_Buy_Limit(t *testing.T) {

	ob := NewOrderBook()

	LoadTestData(ob)

	//Orderbook'da istenilen price hazır bir şekilde varsa direkt işlem gerçekleşeceği için Taker olur.
	ob.AddOrder(Order{ID: 7, UserID: 102, Base: "BTC", Second: "TRY", Type: "limit", Side: BUY, Amount: A("1.1"), Price: A("30600"), Time: time.Now()})

	//Orderbook'da istediği price olmadığı için mecvur bekleyecek. O nedenke maker olacak.
	//ob.AddOrder(Order{ID: 8, UserID: 102, Base: "BTC", Second: "TRY", Type: "limit", Side: SELL, Amount: A("0.70"), Price: A("30501"), Time: time.Now()})

	ob.Debug()
}

func Test_Limit_Sell_A(t *testing.T) {

	ob := NewOrderBook()
	LoadTestData(ob)
	ob.AddOrder(Order{ID: 7, UserID: 102, Base: "BTC", Second: "TRY", Type: "limit", Side: SELL, Amount: A("1.50"), Price: A("30500"), Time: time.Now()})
	ob.Debug()
}

func Test_Limit_Sell_B(t *testing.T) {

	ob := NewOrderBook()
	LoadTestData(ob)

	ob.AddOrder(Order{ID: 7, UserID: 102, Base: "BTC", Second: "TRY", Type: "limit", Side: SELL, Amount: A("1"), Price: A("30550"), Time: time.Now()})
	ob.Debug()
}

func Test_Market_Sell(t *testing.T) {

	ob := NewOrderBook()
	LoadTestData(ob)

	ob.AddOrder(Order{ID: 7, UserID: 102, Base: "BTC", Second: "TRY", Type: "market", Side: SELL, Amount: A("1.20"), Price: A("0"), Time: time.Now()})
	ob.Debug()
}

func Test_Stop_Sell_Market(t *testing.T) {

	ob := NewOrderBook()
	LoadTestData(ob)

	//STOP: 30.400
	ob.AddStop(Order{ID: 7, UserID: 104, Base: "BTC", Second: "TRY", Type: "market", Side: SELL, Stop: A("30400"), Amount: A("1.01"), Price: A("0"), Time: time.Now()})
	ob.AddStop(Order{ID: 8, UserID: 104, Base: "BTC", Second: "TRY", Type: "market", Side: SELL, Stop: A("30400"), Amount: A("0.5"), Price: A("0"), Time: time.Now()})

	ob.AddOrder(Order{ID: 9, UserID: 100, Base: "BTC", Second: "TRY", Type: "market", Side: SELL, Amount: A("1.1"), Price: A("0"), Time: time.Now()})

	ob.Debug()
}

func Test_Stop_Limit_Sell(t *testing.T) {

	ob := NewOrderBook()
	LoadTestData(ob)

	ob.AddStop(Order{ID: 7, UserID: 104, Base: "BTC", Second: "TRY", Type: "limit", Side: SELL, Stop: A("30400"), Amount: A("1.0"), Price: A("30300"), Time: time.Now()})
	ob.AddOrder(Order{ID: 8, UserID: 100, Base: "BTC", Second: "TRY", Type: "market", Side: SELL, Amount: A("1.1"), Price: A("0"), Time: time.Now()})

	ob.Debug()
}

func Test_Stop_Market_Buy(t *testing.T) {

	ob := NewOrderBook()
	LoadTestData(ob)

	ob.AddStop(Order{ID: 7, UserID: 104, Base: "BTC", Second: "TRY", Type: "market", Side: BUY, Stop: A("30900"), Amount: A("1"), Price: A("0"), Easy: true, Time: time.Now()})
	ob.AddOrder(Order{ID: 8, UserID: 100, Base: "BTC", Second: "TRY", Type: "market", Side: BUY, Amount: A("1.1"), Price: A("0"), Time: time.Now()})

	ob.Debug()
}

func Test_Stop_Limit_Buy(t *testing.T) {

	ob := NewOrderBook()

	LoadTestData(ob)

	ob.AddStop(Order{ID: 7, UserID: 106, Base: "BTC", Second: "TRY", Type: "limit", Side: BUY, Amount: A("0.3"), Price: A("30650"), Stop: A("30610"), Time: time.Now()})
	ob.Debug()
	ob.AddOrder(Order{ID: 8, UserID: 108, Base: "BTC", Second: "TRY", Type: "limit", Side: BUY, Amount: A("1.0"), Price: A("30620"), Time: time.Now()})
	ob.Debug()
}

func BenchmarkMarket(b *testing.B) {

	ob := NewOrderBook()

	for n := 0; n < b.N; n++ {
		ob.AddOrder(Order{ID: 1, UserID: 100, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: decimal.RequireFromString("1.10"), Price: decimal.RequireFromString("30500"), Time: time.Now()})
		ob.AddOrder(Order{ID: 2, UserID: 104, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: decimal.RequireFromString("1.10"), Price: decimal.RequireFromString("30600"), Time: time.Now()})
	}
}

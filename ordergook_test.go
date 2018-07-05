package matching

import (
	"fmt"
	"testing"
	"time"
)

func printOrderBook(ob *OrderBook) {
	fmt.Println("FILLEDBOOK")
	for i := 0; i < len(ob.fills); i++ {
		fmt.Printf("Price: %f, Amonth: %f Order ID: %d, Matched ID: %d\r\n", ob.fills[i].Price, ob.fills[i].Amount, ob.fills[i].OrderID, ob.fills[i].MatchOrderID)
	}

	fmt.Println("ORDERBOOK BUYS")
	if len(ob.asks) == 0 {
		fmt.Println("---empty---")
	}
	for i := 0; i < len(ob.asks); i++ {
		fmt.Printf("ID: %d, Amonth: %f\r\n", ob.asks[i].ID, ob.asks[i].Amount)
	}

	fmt.Println("ORDERBOOK SELLS")
	if len(ob.bids) == 0 {
		fmt.Println("---empty---")
	}
	for i := 0; i < len(ob.bids); i++ {
		fmt.Printf("ID: %d, Amonth: %f\r\n", ob.bids[i].ID, ob.bids[i].Amount)
	}
}
func TestSortingOrderBook(t *testing.T) {

	ob := NewOrderBook()

	//BUY
	ob.AddOrder(Order{ID: 1, UserID: 100, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 1, Price: 6000, Time: time.Now()})
	ob.AddOrder(Order{ID: 2, UserID: 101, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 0.90, Price: 6001, Time: time.Now()})
	ob.AddOrder(Order{ID: 3, UserID: 102, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 1.01, Price: 6000, Time: time.Now()})
	ob.AddOrder(Order{ID: 4, UserID: 103, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 1.05, Price: 5999, Time: time.Now()})
	ob.AddOrder(Order{ID: 8, UserID: 103, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 1.05, Price: 6001.1, Time: time.Now()})

	//SELL
	ob.AddOrder(Order{ID: 5, UserID: 104, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 1, Price: 6002, Time: time.Now()})
	ob.AddOrder(Order{ID: 6, UserID: 105, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 2, Price: 6003, Time: time.Now()})
	ob.AddOrder(Order{ID: 8, UserID: 107, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 1.02, Price: 6002.1, Time: time.Now()})
	ob.AddOrder(Order{ID: 7, UserID: 106, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 1.02, Price: 6004, Time: time.Now()})
	ob.AddOrder(Order{ID: 9, UserID: 106, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 1.02, Price: 6004.01, Time: time.Now()})  //+
	ob.AddOrder(Order{ID: 10, UserID: 106, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 1.02, Price: 6004.01, Time: time.Now()}) //+

	printOrderBook(ob)
}

func TestBuyMarket(t *testing.T) {

	ob := NewOrderBook()

	//BUY
	ob.AddOrder(Order{ID: 1, UserID: 100, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 1.1, Price: 30500, Time: time.Now()})
	ob.AddOrder(Order{ID: 2, UserID: 101, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 0.20, Price: 30300, Time: time.Now()})
	ob.AddOrder(Order{ID: 3, UserID: 102, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 0.8, Price: 30250, Time: time.Now()})

	//SELL
	ob.AddOrder(Order{ID: 4, UserID: 104, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 0.90, Price: 30600, Time: time.Now()})
	ob.AddOrder(Order{ID: 5, UserID: 105, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 0.75, Price: 30700, Time: time.Now()})
	ob.AddOrder(Order{ID: 6, UserID: 107, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 0.20, Price: 31000, Time: time.Now()})

	ob.AddOrder(Order{ID: 9, UserID: 102, Base: "BTC", Second: "TRY", Type: "market", Side: "ask", Amount: 30000, Price: 0, Time: time.Now()})

	printOrderBook(ob)
}

func TestBuyLimitOrder(t *testing.T) {

	ob := NewOrderBook()

	//BUY
	ob.AddOrder(Order{ID: 1, UserID: 100, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 1.1, Price: 30500, Time: time.Now()})
	ob.AddOrder(Order{ID: 2, UserID: 101, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 0.20, Price: 30300, Time: time.Now()})
	ob.AddOrder(Order{ID: 3, UserID: 102, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 0.8, Price: 30250, Time: time.Now()})

	//SELL
	ob.AddOrder(Order{ID: 4, UserID: 104, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 0.90, Price: 30600, Time: time.Now()})
	ob.AddOrder(Order{ID: 5, UserID: 105, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 0.75, Price: 30700, Time: time.Now()})
	ob.AddOrder(Order{ID: 6, UserID: 107, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 0.20, Price: 31000, Time: time.Now()})

	ob.AddOrder(Order{ID: 9, UserID: 102, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 0.70, Price: 30600, Time: time.Now()})

	printOrderBook(ob)
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

	printOrderBook(ob)
}

func TestSellOrderSenaryoA(t *testing.T) {

	ob := NewOrderBook()

	//SELL
	ob.BidAdd(Order{ID: 4, UserID: 104, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 0.90, Price: 30600, Time: time.Now()})
	ob.BidAdd(Order{ID: 5, UserID: 105, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 0.75, Price: 30700, Time: time.Now()})
	ob.BidAdd(Order{ID: 6, UserID: 107, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 0.20, Price: 31000, Time: time.Now()})

	//BUY
	ob.AskAdd(Order{ID: 1, UserID: 100, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 1.1, Price: 30500, Time: time.Now()})
	ob.AskAdd(Order{ID: 2, UserID: 101, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 0.20, Price: 30300, Time: time.Now()})
	ob.AskAdd(Order{ID: 3, UserID: 102, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 0.8, Price: 30250, Time: time.Now()})

	ob.BidAdd(Order{ID: 7, UserID: 102, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 0.70, Price: 30500, Time: time.Now()})

	fmt.Println("FILLEDBOOK")
	for i := 0; i < len(ob.fills); i++ {
		fmt.Printf("Price: %f, Amonth: %f Order ID: %d, Matched ID: %d\r\n", ob.fills[i].Price, ob.fills[i].Amount, ob.fills[i].OrderID, ob.fills[i].MatchOrderID)
	}

	fmt.Println("ORDERBOOK BUYS")
	for i := 0; i < len(ob.asks); i++ {
		fmt.Printf("Amonth: %f\r\n", ob.asks[i].Amount)
	}

	fmt.Println("ORDERBOOK SELLS")
	for i := 0; i < len(ob.bids); i++ {
		fmt.Printf("Amonth: %f\r\n", ob.bids[i].Amount)
	}

}
func TestSellOrderSenaryoB(t *testing.T) {

	ob := NewOrderBook()

	//SELL
	ob.BidAdd(Order{ID: 4, UserID: 104, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 0.90, Price: 30600, Time: time.Now()})
	ob.BidAdd(Order{ID: 5, UserID: 105, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 0.75, Price: 30700, Time: time.Now()})
	ob.BidAdd(Order{ID: 6, UserID: 107, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 0.20, Price: 31000, Time: time.Now()})

	//BUY
	ob.AskAdd(Order{ID: 1, UserID: 100, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 1.1, Price: 30500, Time: time.Now()})
	ob.AskAdd(Order{ID: 2, UserID: 101, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 0.20, Price: 30300, Time: time.Now()})
	ob.AskAdd(Order{ID: 3, UserID: 102, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 0.80, Price: 30250, Time: time.Now()})

	ob.BidAdd(Order{ID: 7, UserID: 102, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 1.50, Price: 30500, Time: time.Now()})

	fmt.Println("FILLEDBOOK")
	for i := 0; i < len(ob.fills); i++ {
		fmt.Printf("Price: %f, Amonth: %f Order ID: %d, Matched ID: %d\r\n", ob.fills[i].Price, ob.fills[i].Amount, ob.fills[i].OrderID, ob.fills[i].MatchOrderID)
	}

	fmt.Println("ORDERBOOK BUYS")
	for i := 0; i < len(ob.asks); i++ {
		fmt.Printf("Amonth: %f\r\n", ob.asks[i].Amount)
	}

	fmt.Println("ORDERBOOK SELLS")
	for i := 0; i < len(ob.bids); i++ {
		fmt.Printf("Amonth: %f\r\n", ob.bids[i].Amount)
	}
}

func TestSellOrderSenaryoC(t *testing.T) {

	ob := NewOrderBook()

	//SELL
	ob.BidAdd(Order{ID: 4, UserID: 104, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 0.90, Price: 30600, Time: time.Now()})
	ob.BidAdd(Order{ID: 5, UserID: 105, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 0.75, Price: 30700, Time: time.Now()})
	ob.BidAdd(Order{ID: 6, UserID: 107, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 0.20, Price: 31000, Time: time.Now()})

	//BUY
	ob.AskAdd(Order{ID: 1, UserID: 100, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 1.1, Price: 30500, Time: time.Now()})
	ob.AskAdd(Order{ID: 2, UserID: 101, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 0.20, Price: 30300, Time: time.Now()})
	ob.AskAdd(Order{ID: 3, UserID: 102, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 0.80, Price: 30250, Time: time.Now()})

	ob.BidAdd(Order{ID: 7, UserID: 102, Base: "BTC", Second: "TRY", Type: "market", Side: "bid", Amount: 1.20, Price: 0, Time: time.Now()})

	fmt.Println("FILLEDBOOK")
	for i := 0; i < len(ob.fills); i++ {
		fmt.Printf("Price: %f, Amonth: %f Order ID: %d, Matched ID: %d\r\n", ob.fills[i].Price, ob.fills[i].Amount, ob.fills[i].OrderID, ob.fills[i].MatchOrderID)
	}

	fmt.Println("ORDERBOOK BUYS")
	for i := 0; i < len(ob.asks); i++ {
		fmt.Printf("Amonth: %f\r\n", ob.asks[i].Amount)
	}

	fmt.Println("ORDERBOOK SELLS")
	for i := 0; i < len(ob.bids); i++ {
		fmt.Printf("Amonth: %f\r\n", ob.bids[i].Amount)
	}

}

func TestSellOrderSenaryoD(t *testing.T) {

	ob := NewOrderBook()

	//SELL
	ob.BidAdd(Order{ID: 4, UserID: 104, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 0.90, Price: 30600, Time: time.Now()})
	ob.BidAdd(Order{ID: 5, UserID: 105, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 0.75, Price: 30700, Time: time.Now()})
	ob.BidAdd(Order{ID: 6, UserID: 107, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 0.20, Price: 31000, Time: time.Now()})

	//BUY
	ob.AskAdd(Order{ID: 1, UserID: 100, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 1.1, Price: 30500, Time: time.Now()})
	ob.AskAdd(Order{ID: 2, UserID: 101, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 0.20, Price: 30300, Time: time.Now()})
	ob.AskAdd(Order{ID: 3, UserID: 102, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 0.80, Price: 30250, Time: time.Now()})

	ob.BidAdd(Order{ID: 7, UserID: 102, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 1, Price: 30550, Time: time.Now()})

	fmt.Println("FILLEDBOOK")
	for i := 0; i < len(ob.fills); i++ {
		fmt.Printf("Price: %f, Amonth: %f Order ID: %d, Matched ID: %d\r\n", ob.fills[i].Price, ob.fills[i].Amount, ob.fills[i].OrderID, ob.fills[i].MatchOrderID)
	}

	fmt.Println("ORDERBOOK BUYS")
	for i := 0; i < len(ob.asks); i++ {
		fmt.Printf("Amonth: %f\r\n", ob.asks[i].Amount)
	}

	fmt.Println("ORDERBOOK SELLS")
	for i := 0; i < len(ob.bids); i++ {
		fmt.Printf("Amonth: %f\r\n", ob.bids[i].Amount)
	}

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

	printOrderBook(ob)
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

	//STOP: 30.400
	ob.AddStop(Order{ID: 7, UserID: 104, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Stop: 30400, Amount: 1.0, Price: 30300, Time: time.Now()})

	ob.AddOrder(Order{ID: 8, UserID: 100, Base: "BTC", Second: "TRY", Type: "market", Side: "bid", Amount: 1.1, Price: 0, Time: time.Now()})

	printOrderBook(ob)
}

func TestStopMarketSellSenaryoC(t *testing.T) {

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
	ob.AddStop(Order{ID: 7, UserID: 104, Base: "BTC", Second: "TRY", Type: "market", Side: "ask", Stop: 30900, Amount: 10000, Price: 0, Time: time.Now()})

	ob.AddOrder(Order{ID: 8, UserID: 100, Base: "BTC", Second: "TRY", Type: "market", Side: "bid", Amount: 1.1, Price: 0, Time: time.Now()})

	printOrderBook(ob)
}

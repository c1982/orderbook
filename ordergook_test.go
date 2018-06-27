package matching

import (
	"fmt"
	"testing"
	"time"
)

func TestOrderBook(t *testing.T) {

	ob := NewOrderBook()

	//Kullanıcının alış veriş sonucunda eline geçen paradan kesilecek komiyon oranı.
	//Bu kural ortder type'ın tipine göre değişiyor.
	// Type: Maker 0.1 - Limit
	// Tyope: Taker 0.2 - Market

	//BUY
	ob.AskAdd(Order{ID: 1, UserID: 100, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 1, Price: 6000, Time: time.Now()})
	ob.AskAdd(Order{ID: 2, UserID: 101, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 0.90, Price: 6001, Time: time.Now()})
	ob.AskAdd(Order{ID: 3, UserID: 102, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 1.01, Price: 6000, Time: time.Now()})
	ob.AskAdd(Order{ID: 4, UserID: 103, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 1.05, Price: 5999, Time: time.Now()})
	ob.AskAdd(Order{ID: 8, UserID: 103, Base: "BTC", Second: "TRY", Type: "limit", Side: "ask", Amount: 1.05, Price: 6001.1, Time: time.Now()})

	//SELL
	ob.BidAdd(Order{ID: 5, UserID: 104, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 1, Price: 6002, Time: time.Now()})
	ob.BidAdd(Order{ID: 6, UserID: 105, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 2, Price: 6003, Time: time.Now()})
	ob.BidAdd(Order{ID: 8, UserID: 107, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 1.02, Price: 6002.1, Time: time.Now()})
	ob.BidAdd(Order{ID: 7, UserID: 106, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 1.02, Price: 6004, Time: time.Now()})
	ob.BidAdd(Order{ID: 9, UserID: 106, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 1.02, Price: 6004.01, Time: time.Now()})  //+
	ob.BidAdd(Order{ID: 10, UserID: 106, Base: "BTC", Second: "TRY", Type: "limit", Side: "bid", Amount: 1.02, Price: 6004.01, Time: time.Now()}) //+

	//9 ve 10 orderbook tarafında toplanaxcak gösrerilecek. Depth olayu.
	//Fiyatı (price) aynı olan order'ların  Amont'larının toplamı orderlist'de gösterilecek.
	//Priceları toplanmış ve derinliği yapılmış liste derinlik tablosu olacak. Bu liste orderbook listesinden farklı sadece gösterim amaçlı olacak.

	fmt.Println("BUYS")
	for i := 0; i < len(ob.asks); i++ {
		fmt.Println(ob.asks[i])
	}

	fmt.Println("SELLS")
	for i := 0; i < len(ob.bids); i++ {
		fmt.Println(ob.bids[i])
	}

}

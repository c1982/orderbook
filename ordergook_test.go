package matching

import "testing"

func TestOrderBookEmpty(t *testing.T) {
	book := OrderBook{}

	if !book.isEmpty() {
		t.Error("order book is not empty by defaultß")
	}
}

func TestRun(t *testing.T) {

	ob := OrderBook{
		Bids: make(map[int]int),
		Asks: make(map[int]int),
	}

	ob.AddAsk(100, 10)
	ob.AddAsk(101, 10)
	ob.AddBid(99, 20)
	ob.Run()
}

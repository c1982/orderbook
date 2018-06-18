package matching

import (
	"container/heap"
	"fmt"

	"github.com/shopspring/decimal"
)

type OrderSide int
type OrderType int
type OrderStatus int

const (
	SideBUY OrderSide = iota + 1
	SideSELL

	OrderTypeMARKET = iota + 1
	OrderTypeLIMIT
	OrderTypeSTOP
	OrderTypeCANCEL

	OrderStatusNew OrderStatus = iota + 1
	OrderStatusPartiallyFilled
	OrderStatusFilled
	OrderStatusCancelled
	OrderStatusRejected
	OrderStatusPending
)

type Order struct {
	ID          int
	BaseAsset   string
	SecondAsset string
	Status      OrderStatus
	Side        OrderSide
	Type        OrderType
	Price       decimal.Decimal
	Amount      decimal.Decimal
}

func (o *Order) ToPair() string {
	return fmt.Sprintf("%s/%s", o.BaseAsset, o.SecondAsset)
}

type Level struct {
	Price decimal.Decimal `json:"price"`
	Order []*Order        `json:"orders"`
}

type LevelHeap []Level

func (h LevelHeap) Len() int           { return len(h) }
func (h LevelHeap) Less(i, j int) bool { return h[i].Price.LessThan(h[j].Price) }
func (h LevelHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *LevelHeap) Push(x interface{}) {
	for index, level := range *h {
		if x.(Level).Price.Equal(level.Price) {
			(*h)[index].Order = append((*h)[index].Order, x.(Level).Order...)
			return
		}
	}
	*h = append(*h, x.(Level))
}

func (h *LevelHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *LevelHeap) Peek() *Level {
	if h.Len() == 0 {
		return nil
	}
	return &(*h)[0]
}

func (h *LevelHeap) Remove(o Order) error {
	for index, level := range *h {
		if o.Price == level.Price {
			for i, order := range level.Order {
				if o.ID == order.ID {
					(*h)[index].Order = append((*h)[index].Order[:i], (*h)[index].Order[i+1:]...)
					if len((*h)[index].Order) == 0 {
						heap.Remove(h, index)
					}
					return nil
				}
			}
		}
	}

	return fmt.Errorf("failed")
}

func NewTrade(o Order) {

	switch o.Side {
	case SideBUY:
		//Bir şey
	case SideSELL:
	//Bir şeuy
	default:
		//Rejected

	}
}

package valueobject

import "time"

type Delivery struct {
	CreatedAt time.Time // data e hora da entrega
	Location  string    // balcão, mesa, delivery
	OrderOf   string    // entrega antes de tal coisa
}

func NewDelivery(location, orderOf string) Delivery {
	return Delivery{
		CreatedAt: time.Now(),
		Location:  location,
		OrderOf:   orderOf,
	}
}

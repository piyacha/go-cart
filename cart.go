package cart

type Cart struct {
	items	[]Item
}

type Item struct {
	name string 
	price float64
}

func (c Cart) AddToCart(i Item) []Item {
	c.items = append(c.items, i)

	return c.items
}

func (c Cart) TotalPrice() (total float64) {

	for _, item := range c.items {
		total += item.price
	}

	return
}

package resources

type Product struct {
	Id              string   `json:"Id"`
	Name            string   `json:"Name"`
	QuantityInStock Quantity `json:"QuantityInStock"`
	Price           Cent     `json:"Price"`
}

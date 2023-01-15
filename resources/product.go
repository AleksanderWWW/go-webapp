package resources


type Product struct {
	Id string `json:"Id"`
	Name string `json:"Name"`
	QuantityInStock int `json:"QuantityInStock"`
	Price int `json:"Price"`
}

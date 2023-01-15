package resources


type Product struct {
	Id string `json:"Id"`
	Name string `json:"Name"`
	QuantityInStock int16 `json:"QuantityInStock"`
	Price int16 `json:"Price"`
}

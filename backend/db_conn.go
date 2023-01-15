package backend

import "github.com/AleksanderWWW/go-webapp/resources"

var allProducts []resources.Product = []resources.Product{
	{Id: "C1", Name: "Chair", QuantityInStock: 200, Price: 7850},
	{Id: "C2", Name: "Chair", QuantityInStock: 100, Price: 9850},
	{Id: "B1", Name: "Bed", QuantityInStock: 900, Price: 23378},
}


func RetrieveAll() []resources.Product {
	return allProducts
}

func RetrieveByID(id string) resources.Product {
	for _, p := range allProducts {
		if p.Id == id {
			return p
		}
	}
	return resources.Product{}
}

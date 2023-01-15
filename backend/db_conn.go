package backend

import "github.com/AleksanderWWW/go-webapp/resources"


func RetrieveAll() []resources.Product {
	return []resources.Product{
		{Id: "C1", Name: "Chair", QuantityInStock: 200, Price: 7850},
		{Id: "C2", Name: "Chair", QuantityInStock: 100, Price: 9850},
		{Id: "B1", Name: "Bed", QuantityInStock: 900, Price: 23378},
	}
}

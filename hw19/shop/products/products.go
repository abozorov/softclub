package products

type Product struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

func TotalPrice(products []Product) float64 {
	tP := 0.0
	for _, v := range products {
		tP += v.Price * float64(v.Quantity)
	}
	return tP
}

func MostExpensive(products []Product) Product {
	mxPP := new(Product)
	for _, v := range products {
		if mxPP.Price < v.Price {
			mxPP = &v
		}
	}
	return *mxPP
}

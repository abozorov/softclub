package main

import (
	"errors"
	"fmt"
)

type Delivery interface {
	Calculate(weight float64) (float64, error)
}

type CarDelivery struct{}
func(с CarDelivery)	Calculate(weight float64) (float64, error) {
	return weight * 10.0, nil
}

type AirDelivery struct{}
func(a AirDelivery)	Calculate(weight float64) (float64, error) {
	if weight < 4.0 {
		return -1, errors.New("Summ not enougth")
	}
	return weight * 25.0, nil
}

type ExpressDelivery struct{}
func(e ExpressDelivery)	Calculate(weight float64) (float64, error) {
	if weight > 10.0 {
		return -1, errors.New("Weigth > 10")
	}
	return weight * 40.0, nil
}

func CalculateAll(deliveries []Delivery, weight float64) (calc []float64, errs []error) {
	for _, v := range deliveries {
		c, err := v.Calculate(weight)
		calc = append(calc, c)
		errs  = append(errs, err)
	}
	return
}

func main() {
	deliveries := []Delivery{
		CarDelivery{},
		AirDelivery{},
		ExpressDelivery{},
		CarDelivery{},
		AirDelivery{},
		ExpressDelivery{},
	}
	fmt.Println(CalculateAll(deliveries, 102))
	fmt.Println(CalculateAll(deliveries, 10))
	fmt.Println(CalculateAll(deliveries, 4))
	fmt.Println(CalculateAll(deliveries, 3))

}

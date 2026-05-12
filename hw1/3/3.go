package main

import "fmt"

func main() {
	var a, b, c, d, e, positive, negative, zero int
	fmt.Scan(&a, &b, &c, &d, &e)

	if a > 0 {
		positive++
	} else if a < 0 {
		negative++
	} else {
		zero++
	}

	if b > 0 {
		positive++
	} else if b < 0 {
		negative++
	} else {
		zero++
	}

	if c > 0 {
		positive++
	} else if c < 0 {
		negative++
	} else {
		zero++
	}

	if d > 0 {
		positive++
	} else if d < 0 {
		negative++
	} else {
		zero++
	}

	if e > 0 {
		positive++
	} else if e < 0 {
		negative++
	} else {
		zero++
	}

	fmt.Printf("Positive=%d Negative=%d Zero=%d\n", positive, negative, zero)
}

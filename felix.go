package main

import (
	"fmt"
	"time"
)

type Address struct {
	street  string
	city    string
	zip     string
	state   string
	country string
}

type Suppliers struct {
	Sup []Supplier
}

type Supplier struct {
	id             string
	name           string
	PaymentAddress Address
	Merchendize    []Product
}

type Customers struct {
	cust []Customer
}

type Customer struct {
	id        string
	FirstName string
	LastName  string
	Shipping  Address
	Billing   Address
	Cart      []Product
}

type Product struct {
	pks      string
	name     string
	desc     string
	category string
	retail   float32
	sale     float32
	discount float32
	qty      int
}

type inventry struct {
	id    string
	stock []Product
}

type purchaseOrder struct {
	id       string
	datetime string
	status   string
	items    []Product
}

func main() {
	var (
		tim int = 1000
	)

	fmt.Println("Retail Loop")
	for i := 0; i < 10; i++ {
		sleeper(tim)
		fmt.Println(i)
	}
}

func sleeper(tim int) {
	time.Sleep(time.Duration((tim)) * time.Millisecond)
}

package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	testProductList := []product{}
	testProductPointerList := []*product{}
	for i := 0; i < 100000; i++ {
		testProductList = append(testProductList, product{i, customer{"j", address{"test city", locationDetails{"mel"}}}})
		testProductPointerList = append(testProductPointerList, &product{i, customer{"j", address{"test city", locationDetails{"mel"}}}})
	}

	var (
		sumPointerFuncDuration int64
		sumFuncDuration        int64
	)
	retryTimes := 30
	for j := 0; j < retryTimes; j++ {
		startB := time.Now()
		sumPointerFuncDuration += testPointerFunc(testProductPointerList, startB)
		start := time.Now()
		sumFuncDuration += testFunc(testProductList, start)
	}
	println(fmt.Sprintf("total pointer func duration is: %v", sumPointerFuncDuration))
	println(fmt.Sprintf("total func duration is: %v", sumFuncDuration))
	println(fmt.Sprintf("average pointer func duration is: %v", sumPointerFuncDuration/int64(retryTimes)))
	println(fmt.Sprintf("average func duration is: %v", sumFuncDuration/int64(retryTimes)))
}

func testFunc(testList []product, startTime time.Time) (duration int64) {
	//init an order
	order := order{}
	//append the valid products into the order; for demo purpose we only check the customer name
	for _, p := range testList {
		if len(p.customer.name) > 0 {
			order.Products = append(order.Products, p)
		}
	}
	//marshal the order struct
	json.Marshal(order)
	elapsed := time.Since(startTime)
	return elapsed.Nanoseconds()
}

func testPointerFunc(testList []*product, startTime time.Time) (duration int64) {
	order := orderWithPointer{}
	for _, p := range testList {
		if len(p.customer.name) > 0 {
			order.Products = append(order.Products, p)
		}
	}
	json.Marshal(order)
	elapsed := time.Since(startTime)
	return elapsed.Nanoseconds()
}

type product struct {
	index    int
	customer customer
}
type customer struct {
	name    string
	address address
}
type address struct {
	locationName    string
	locationDetails locationDetails
}
type locationDetails struct {
	cityCode string
}

type orderWithPointer struct {
	Products []*product
}

type order struct {
	Products []product
}

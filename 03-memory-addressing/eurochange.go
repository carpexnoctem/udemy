package main

import "fmt"

const euroconv float64 = .85 //this is the current euro>usd converstion rate

func main() {
	var usd float64
	fmt.Print("How Many USD Dollars Do You Have?")
	fmt.Scan(&usd)
	euros := usd * euroconv
	fmt.Println(usd, "Dollars is currently", euros, "Euros")

}

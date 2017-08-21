package main

import "fmt"

const euroconv float64 = .85 //this is the current euro>usd converstion rate

func main() {
	var usd float64 //expect the variable usd to be a float64 value
	fmt.Print("How Many USD Dollars Do You Have?")
	fmt.Scan(&usd) //fmt.Scan will parse the returned input and the & stores the returned value in memory
	euros := usd * euroconv //just assigning a var to a formula
	fmt.Println(usd, "Dollars is currently", euros, "Euros")

}

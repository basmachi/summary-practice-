package main

import (
	"fmt"
)

func main() {
	limitCashback := 30
	percent := 10
	/*purchase1 := 1040
	purchase2 := 1020
	purchase3 := 1080
	purchase4 := 1050
	purchase5 := 1070
	//sumPurchases :=purchase1 + purchase2 + purchase3 + purchase4 + purchase5   //obshaya summa pokupok
	*/
	purchases := [5]int{10,20,7,14,9}
//	fmt.Println(purchases)
	sumPurchases := 0
	for _, purchase := range purchases {
		sumPurchases += purchase
		//fmt.Print(value, " ")
	}
	sumCashback := sumPurchases * percent / 100   //procent ot obshey summi
	if sumCashback > limitCashback {
		fmt.Println(limitCashback)
	} else {
		fmt.Println(sumCashback)
	}
}


	


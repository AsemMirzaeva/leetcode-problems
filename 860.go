package main

import "fmt"

func lemonadeChange(bills []int) bool {
	count5 := 0
	count10 := 0

	for _, v := range bills {

		if v == 5 {
			count5++
		} else if v == 10 {
			if count5 == 5 {
				return false
			}
			count10++
			count5--

		} else if v == 20 {
			if count10 > 0 && count5 > 0 {
				count10--
				count5--
			} else if count5 >= 3 {
				count5 -= 3
			} else {
				return false
			}
		}
		
	}
	return true
}

func main() {
	// b1 := []int{5, 5, 5, 10, 20}
	// b2 := []int{5,5,5,10,10,20,10,20,5,5}
	b3 := []int{5, 5, 5, 10, 5, 20, 5, 10, 5, 20}
	// fmt.Println(lemonadeChange(b1))
	// fmt.Println(lemonadeChange(b2))
	fmt.Println(lemonadeChange(b3))
}

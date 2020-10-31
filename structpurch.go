package main

import "fmt"

type purch struct {
	purchases []int64
	personCashback int64
  	SumOfCashback int64
}
func (c *purch) count(){
	SumOfPurchases := int64(0)
	for _, value := range c.purchases{
		SumOfPurchases += value
	}
	c.SumOfCashback = (SumOfPurchases) * c.personCashback / 100
	fmt.Println(c)
}

func main() {
	var c = purch{
		purchases: []int64{100, 150, 200},
		personCashback: 15,
		SumOfCashback: 0,
	}
	fmt.Println(c)
	c.count()
	//	count(&c)
	fmt.Println(c)
}



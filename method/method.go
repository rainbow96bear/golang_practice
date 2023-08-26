package main

import "fmt"

type money uint

func (dollar money) checkEnough(check money) bool {
	result := dollar >= check
	fmt.Println(result)
    return result
}
func (money) what(check money) bool {
	
	fmt.Println(check)
    return false
}
func (dollar *money) deposite(addAmount money) money {
	*dollar = *dollar + addAmount
    return *dollar
}

func (dollar *money) withdraw(subAmount money) money {
	*dollar = *dollar - subAmount
    return *dollar
}

func main(){
	depositeFn := (*money).deposite
	withdrawFn := (*money).withdraw
	checkEnoughFn := (money).checkEnough
	var myDollar money = 1000
	fmt.Println(myDollar)

	checkEnoughFn(myDollar,1500)
	
	depositeFn(&myDollar, 600)
	fmt.Println(myDollar)
	
	checkEnoughFn(myDollar, 1500)

	withdrawFn(&myDollar, 1500)
	fmt.Println(myDollar)
}






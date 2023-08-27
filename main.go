package main

import (
	"fmt"
)


func coinChange(cash *int, price *int) *map[int]int{
     coins := map[int]int{}
     coinType := []int{1,2,5,10,20,50,100,500,1000}
     remains := *cash - *price
    if(*cash < *price){
       fmt.Printf("Cash is not enough to pay.")
       return &coins
    }

    for remains > 0 {
        for index := len(coinType) - 1; index >= 0; index-- {               
             coinCount := remains / coinType[index]
             amount := coinCount * coinType[index]
             if(amount <= remains){
                coins[coinType[index]] = coinCount
                remains = remains - amount                
             }
            
        }
    }


    return &coins
}

func main() {
    var cash int;
    var price int;

    fmt.Printf("Enter Cash: ")
    fmt.Scanf("%d",&cash)
    fmt.Printf("Enter Total Price: ")
    fmt.Scanf("%d",&price)

    coins := coinChange(&cash, &price)
    fmt.Printf("Change %v ", coins)
 
}

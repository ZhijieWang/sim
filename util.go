package main

import (

)

//MarketMakerTrader is responsible for placing the original order in IPO process
type MarketMakerTrader struct {
}

//Receive implements the actor interface
//func (mm *MarketMakerTrader) Receive(context actor.Context) {
//	switch msg := context.Message().(type) {
//	case *actor.Started:
//		quotes, err := getRootContext().RequestFuture(getExchange(), GetQuoteMessage{}, time.Second).Result()
//		if err != nil {
//			panic("Exchange failed to return quotes for proper IPO")
//		}
//		resultList := quotes.(StockQuoteMessage).Value
//		for k := range resultList {
//			order := CreateMarketOrder(*context.Self(), k, 1, 1, Ask)
//			result, _ := getRootContext().RequestFuture(getExchange(), SubmitOrderRequest{k, order}, time.Second).Result()
//			if _, ok := result.(OrderConfirmation); !ok {
//				panic("Order Submission Failed")
//			} else {
//				fmt.Println("Order Confirmed")
//			}
//		}
//	case OrderConfirmation:
//		fmt.Println(msg)
//	case OrderFulfillment:
//		fmt.Printf("Order %v+: Partially fulfilled\n", msg.OrderID)
//	}

//}

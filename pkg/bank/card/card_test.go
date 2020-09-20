package card

import (
	"bank/pkg/bank/types"
	"fmt"
)


func ExamplePaymentSources()  {
	cards := []types.Card{
		{
			PAN: "5058 xxxx xxxx 8888",
			Balance: 10_000_00,
			Active: true,
		},
		{
			PAN: "5058 xxxx xxxx 9999",
			Balance: 10_000_00,
			Active: false,
		},
		{
			PAN: "5058 xxxx xxxx 8888",
			Balance: -10_000_00,
			Active: true,
		},

	}	
		paymentSources := PaymentSources(cards)
	for _, alif := range paymentSources {
		fmt.Println(alif.Number)		
	}
	//Output:
	//5058 xxxx xxxx 8888
	
}

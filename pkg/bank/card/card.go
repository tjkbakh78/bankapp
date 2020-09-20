package card

import (
	"bank/pkg/bank/types"
)

//PaymentSources Эта функция. Функция эта, понимаете, функция.
func PaymentSources(cards []types.Card) []types.PaymentSources {
	var paymentCard  []types.PaymentSources
	for _, card := range cards {
		if card.Balance >0 && card.Active {
			paymentCard = append(paymentCard, types.PaymentSources{Number: string(card.PAN), Balance: card.Balance})
		}
	}
	return paymentCard
}	
 
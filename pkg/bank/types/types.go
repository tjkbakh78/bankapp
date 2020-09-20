package types


//Money тип это деньги
type Money int64
 
//Currency тип валюты там наши или русские или всеми любимые америкосы
type Currency string

//PAN  пан?! Кажется это номер карты, если, конечно память не изменяет.
type PAN string 

//MinBalance это минимальный баланс на карте 
type MinBalance Money

//Card тип - структура 
type Card struct{
	ID 				int 
	PAN 			PAN
	Balance 		Money 
	MinBalance		Money
	Currency 		Currency
	Color 			string
	Name 			string 
	Active 			bool
}
//Payment Платежи картой. 
type Payment struct{
	ID				int
	Amount			Money
}

//PaymentSources этo Источник оплаты
type PaymentSources struct{
	Type 			string 	// 'card'
	Number 			string  // номер вида '5058 XXXX XXXX 8888'
	Balance			Money   // балансик в дирамонях
} 
package process_transaction

type TransactionInputDto struct {
	AccountId string
	Amount    float64
}

type TransactionOutputDto struct {
	Id           string
	Status       string
	ErrorMessage string
}

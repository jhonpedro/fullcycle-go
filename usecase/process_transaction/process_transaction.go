package process_transaction

import (
	"github.com/jhonpedro/fullcycle-go/entities"
	"github.com/jhonpedro/fullcycle-go/utility/constants"
)

type ProcessTransaction struct {
	repository       entities.TransactionRepository
	uniqueIdentifier entities.UniqueIdentifierService
}

func NewProcessTransaction(repository entities.TransactionRepository, uniqueIdentifier entities.UniqueIdentifierService) *ProcessTransaction {
	return &ProcessTransaction{repository, uniqueIdentifier}
}

func (p *ProcessTransaction) Execute(input TransactionInputDto) (TransactionOutputDto, error) {

	id := p.uniqueIdentifier.Generate()

	transaction := entities.NewTransaction(id, input.AccountId, input.Amount)

	isTransactionValid := transaction.IsValid()

	if isTransactionValid != nil {
		return p.rejectTransaction(input, id, isTransactionValid)
	}

	return p.aproveTransaction(input, id)
}

func (p *ProcessTransaction) rejectTransaction(input TransactionInputDto, id string, err error) (TransactionOutputDto, error) {
	repositoryInsertError := p.repository.Insert(id, input.AccountId, input.Amount, constants.Rejected, err.Error())

	if repositoryInsertError != nil {
		return TransactionOutputDto{}, repositoryInsertError
	}

	return TransactionOutputDto{
		Id:           id,
		Status:       constants.Rejected,
		ErrorMessage: err.Error(),
	}, nil
}

func (p *ProcessTransaction) aproveTransaction(input TransactionInputDto, id string) (TransactionOutputDto, error) {
	repositoryInsertError := p.repository.Insert(id, input.AccountId, input.Amount, constants.Approved, "")

	if repositoryInsertError != nil {
		return TransactionOutputDto{}, repositoryInsertError
	}

	return TransactionOutputDto{
		Id:           id,
		Status:       constants.Approved,
		ErrorMessage: "",
	}, nil
}

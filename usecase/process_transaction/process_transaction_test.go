package process_transaction

import (
	"testing"

	"github.com/golang/mock/gomock"
	mock_entities "github.com/jhonpedro/fullcycle-go/entities/mock_entities"
	"github.com/jhonpedro/fullcycle-go/utility/constants"
	"github.com/stretchr/testify/assert"
)

func TestProcessTransactionValidOne(t *testing.T) {
	input := TransactionInputDto{
		AccountId: "1",
		Amount:    200,
	}

	expectedOutput := TransactionOutputDto{
		Id:           "uuid",
		Status:       "approved",
		ErrorMessage: "",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repositoryMock := mock_entities.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().Insert("uuid", input.AccountId, input.Amount, constants.Approved, "").Return(nil)

	uniqueIdentifierMock := mock_entities.NewMockUniqueIdentifierService(ctrl)
	uniqueIdentifierMock.EXPECT().Generate().Return("uuid")

	usecase := NewProcessTransaction(repositoryMock, uniqueIdentifierMock)

	output, err := usecase.Execute(input)

	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}

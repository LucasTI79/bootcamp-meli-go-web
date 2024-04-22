package products

import (
	"encoding/json"
	"testing"

	"github.com/batatinha123/products-api/pkg/store"
	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	input := []Product{
		{
			ID:       1,
			Name:     "CellPhone",
			Category: "tech",
			Count:    3,
			Price:    250,
		},
		{
			ID:       1,
			Name:     "Notbook",
			Category: "tech",
			Count:    10,
			Price:    1750.5,
		},
	}

	dataJson, _ := json.Marshal(input)

	dbMock := store.Mock{
		Data: dataJson,
	}

	storeStub := store.FileStoreMock{
		FileName: "",
		Mock:     &dbMock,
	}

	myRepo := NewRepository(&storeStub)

	resp, _ := myRepo.GetAll()

	assert.Equal(t, input, resp)

}

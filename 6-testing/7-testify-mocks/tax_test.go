package tax

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTaxAndSave(t *testing.T) {
	repository := &TaxRepositoryMock{}
	repository.On("SaveTax", 10.0).Return(nil).Once()
	// repository.On("SaveTax", 10.0).Return(nil).Once()
	repository.On("SaveTax", 0.0).Return(errors.New("error saving tax"))
	// repository.On("SaveTax", mock.Anything).Return(errors.New("error saving tax"))

	err := CalculateTaxAndSave(1000.0, repository)
	assert.Nil(t, err)

	err = CalculateTaxAndSave(0.0, repository)
	assert.Error(t, err, "error saving tax")

	repository.AssertExpectations(t)
	repository.AssertNumberOfCalls(t, "SaveTax", 2)
}

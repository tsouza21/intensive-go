package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenAnEmptyID_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := Order {}
	assert.Error(t, order.isValid(), "invalid id")
}

func TestGivenAnEmptyPrice_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := Order{ ID: "123" }
	assert.Error(t, order.isValid(), "invalid price")
}

func TestGivenAnEmptyTax_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := Order{ ID: "123", Price: 22.33 }
	assert.Error(t, order.isValid(), "invalid tax")
}

func TestGivenAInvalidPriceAndTax_WhenICallNewOrderFunc_ThenIShouldReceibeAnError(t *testing.T) {
	order, err := NewOrder("123", 0, 0)
	
	assert.Error(t, err)
	assert.Nil(t, order)
}

func TestGivenAValidParams_WhenICallNewOrder_ThenIShouldReceiveCreatedOrderWithAllParams(t *testing.T) {
	order := Order{
		ID: "123", Price: 22.33, Tax: 2.2,
	}

	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 22.33, order.Price)
	assert.Equal(t, 2.2, order.Tax)
	assert.Nil(t, order.isValid())
}

func TestGivenAValidParams_WhenICallNewOrderFunc_ThenIShouldReceiveCreatedOrderWithAllParams(t *testing.T) {
	order, err := NewOrder("123", 10.0, 3.0)

	assert.Nil(t, err)
	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 3.0, order.Tax)
}

func TestGivenAValidParams_WhenICallCalculatePrice_ThenIShouldSetFinalPrice(t *testing.T) {
	order, err := NewOrder("123", 10.0, 3.0)

	assert.Nil(t, err)
	assert.Nil(t, order.CalculateFinalPrice())
	assert.Equal(t, 13.0, order.FinalPrice)
}
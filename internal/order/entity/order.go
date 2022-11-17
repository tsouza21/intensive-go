package entity

import "errors"

type Order struct {
	ID string
	Price float64
	Tax float64
	FinalPrice float64
}

func NewOrder(id string, price float64, tax float64) (*Order, error) {
	order := &Order{
		ID: id,
		Price: price,
		Tax: tax,
	}

	err := order.isValid()

	if err != nil {
		return nil, err		
	}
	
	return order, nil
}

func (order *Order) isValid() error {
	if order.ID == "" {
		return errors.New("invalid id")
	}

	if order.Price <= 0 {
		return errors.New("invalid price")
	}

	if order.Tax <= 0 {
		return errors.New("invalid tax")
	}

	return nil
}

func (order *Order) CalculateFinalPrice() error {
	err := order.isValid()

	if err != nil {
		return err
	}

	order.FinalPrice = order.Price + order.Tax

	return nil
}
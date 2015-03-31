package model

import (
	"errors"
)

type Production struct {
	productionType string
	initialPrice   CurrencyAmount
	dependsOn      *Production
}

type GoodsAmount int

func (p *Production) Type() string {
	return p.productionType
}

func (p *Production) InitialPrice() CurrencyAmount {
	return p.initialPrice
}

func NewProduction(commonName string, initialPrice CurrencyAmount) (*Production, error) {
	if commonName == "" {
		return nil, errors.New("Common name mustn't be empty string")
	}

	if initialPrice <= 0 {
		return nil, errors.New("Initial price mustn't be less then zero")
	}

	return &Production{
		productionType: commonName,
		initialPrice:   initialPrice,
	}, nil
}

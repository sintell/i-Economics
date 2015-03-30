package model

import (
	"errors"
	"math/rand"
)

// has: production amount field
//      production cost field
//      tradeing cost field
//      production fonds aging field
//      production rate field
type Company struct {
	id               int32
	name             string
	organizationForm string
	initialCapital   CurrencyAmount

	productionAmount int
	productionRate   int
	productionCost   CurrencyAmount
}

func (c *Company) RegistrationId() int32 {
	return c.id
}

func (c *Company) Name() string {
	return c.name
}

func (c *Company) OrganizationForm() string {
	return c.organizationForm
}

func (c *Company) InitialCapital() CurrencyAmount {
	return c.initialCapital
}

func (c *Company) CurrentCost() CurrencyAmount {
	return c.productionCost
}

func (c *Company) SellTo(recipient *Trader, subject *Goods) (*Deal, error) {
	return nil, errors.New("NEI")
}

func (c *Company) BuyFrom(donor *Trader, subject *Goods) (*Deal, error) {
	return nil, errors.New("NEI")
}

func NewCompany() *Company {
	return &Company{
		rand.Int31(),
		"Test",
		"OOO",
		300000.00,
		0,
		0,
		0,
	}
}

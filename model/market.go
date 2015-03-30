package model

import (
	"errors"
)

// has: production amount field
//      production cost field
//      production need amount field
type Market struct {
	lots []*Lot
	info marketInfo
}

type marketInfo struct {
	productionAmount     GoodsAmount
	productionNeedAmount GoodsAmount
	productionCost       CurrencyAmount
}

type Lot struct {
	object           *Goods
	subject          *Trader
	price            CurrencyAmount
	amount           GoodsAmount
	registrationDate ModelTime
}

func (c *Market) SellTo(recipient *Trader, subject *Goods) (*Deal, error) {
	return nil, errors.New("NEI")
}

func (c *Market) BuyFrom(donor *Trader, subject *Goods) (*Deal, error) {
	return nil, errors.New("NEI")
}

func (m *Market) Find(lotSearchParams ...interface{}) []*Lot {
	result := []*Lot{}

	for _, searchParam := range lotSearchParams {
		for _, lot := range m.lots {
			switch searchParam.(type) {
			// This cases temporally disapled due to nonexistent search criterias
			// case *Goods: {

			// }
			// case *Trader: {

			// }
			case CurrencyAmount:
				{
					if lot.price == searchParam.(CurrencyAmount) {
						result = append(result, lot)
					}
				}

			case GoodsAmount:
				{
					if lot.amount == searchParam.(GoodsAmount) {
						result = append(result, lot)
					}
				}
			}

		}

	}
	return result
}

func (m *Market) Info() marketInfo {
	return m.info
}

func NewMarket() *Market {
	return &Market{
		[]*Lot{},
		marketInfo{},
	}
}

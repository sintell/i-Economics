package model

type Money struct {
	amount float64
	owner  *Trader
}

type CurrencyAmount float64

func (m *Money) Amount() float64 {
	return m.amount
}

func (m *Money) Transfer(target *Trader) (*DealInfo, error) {
	deal, err := (*target).Accept(m)

	if err != nil {
		return nil, err
	}

	return (*deal).Info(), nil
}

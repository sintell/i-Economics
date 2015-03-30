package model

type Production struct {
	amount int
	price  float64
}

type GoodsAmount int

func (p *Production) Amount() int {
	return p.amount
}

func (p *Production) Transfer(trader *Trader) (*DealInfo, error) {
	deal, err := (*trader).Accept(p)

	if err != nil {
		return nil, err
	}

	return (*deal).Info(), nil
}

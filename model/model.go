package model

import (
	"github.com/sintell/i-Economics/utils"
	"log"
	"os"
)

func init() {
	logger = utils.NewLogger(utils.DEBUG, os.Stdout, log.Ldate|log.Ltime|log.Lmicroseconds)
}

var logger *utils.Logger

//  Contains model which depends on
//  Market struct
//  Company struct
//

type Trader interface {
	SellTo(*Trader, *Currency) (*Deal, error)
	BuyFrom(*Trader, *Currency) (*Deal, error)
	Accept(interface{}) (*Deal, error)
	Account() *Account
}

type Organization interface {
	RegistrationId() int32
	Name() string
	OrganizationForm() string
	InitialCapital() CurrencyAmount
}

type Currency interface {
	Amount() float64
	Transfer(*Trader) (*DealInfo, error)
}

type Goods interface {
	Amount() int
	InitialCost() CurrencyAmount
	Transfer(*Trader) (*DealInfo, error)
}

type Deal interface {
	Info() *DealInfo
}

type Account interface {
	Amount() float64
	Withdraw(float64) (bool, error)
	Increase(float64) (bool, error)
}

type ModelTime int64

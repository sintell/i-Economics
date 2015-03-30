package model

import (
	"errors"
	"strconv"
)

type TransferError struct {
	description   string
	operationType int
}

func (t *TransferError) Error() string {
	return t.description
}

func (t *TransferError) Type() int {
	return t.operationType
}

type DealInfo struct {
	Type         string
	AmountTraded int
	Subject      *Goods
	Price        *Currency
}

type Bank struct {
	accounts map[int32]*BankAccount
}

type BankAccount struct {
	amount        CurrencyAmount
	operationsLog []string
	owner         *Organization
}

func (b *Bank) Transfer(from *Account, to *Account, amount CurrencyAmount) (bool, error) {
	(*from).Withdraw(float64(amount))
	(*to).Increase(float64(amount))

	return true, nil
}

func (b *Bank) OpenAccount(org Organization) error {
	if _, exists := b.accounts[org.RegistrationId()]; exists {
		return errors.New("Account for that registration id already exists")
	}
	b.accounts[org.RegistrationId()] = &BankAccount{
		org.InitialCapital(),
		[]string{"CREATION", "INCOME:" + strconv.FormatFloat(float64(org.InitialCapital()), byte('f'), 2, 64)},
		&org,
	}
	return nil
}

func (b *BankAccount) Amount() CurrencyAmount {
	return b.amount
}

func (b *BankAccount) Withdraw(withdrawAmount CurrencyAmount) (bool, error) {
	if b.amount < withdrawAmount {
		return false, &TransferError{"Withdraw amount exeeds current funds awailable on account", 0}
	}

	b.amount = b.amount - withdrawAmount
	b.operationsLog = append(b.operationsLog, "WITHDRAW")
	return true, nil
}

func (b *BankAccount) Increase(income CurrencyAmount) (bool, error) {
	b.amount = b.amount + income
	b.operationsLog = append(b.operationsLog, "INCOME")
	return true, nil
}

func NewBank() *Bank {
	return &Bank{
		make(map[int32]*BankAccount),
	}
}

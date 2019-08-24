/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-20 15:45:48
 * @LastEditTime: 2019-08-21 10:08:41
 * @LastEditors: Please set LastEditors
 */
package day02

import "errors"

type coin int
type creditCard int
type BankCard int

const (
	CREDITCARD = "creditCard"
	BANKCARD   = "BankCard"
)

type Wallet struct {
	balance       coin       //零钱
	creditbalance creditCard //信用卡
	Bankbalance   BankCard   //银行卡

}

func (w *Wallet) Deposit(amount int, card string) {
	switch card {
	case CREDITCARD:
		w.creditbalance += creditCard(amount)
	case BANKCARD:
		w.Bankbalance += BankCard(amount)
	default:
		w.balance += coin(amount)
	}
}

var InsufficientFundsError = errors.New("超出了总金额")

func (w *Wallet) Withdraw(amount int, card string) error {
	switch card {
	case CREDITCARD:
		if creditCard(amount) > w.creditbalance {
			return InsufficientFundsError
		}
		w.creditbalance -= creditCard(amount)
	case BANKCARD:
		if BankCard(amount) > w.Bankbalance {
			return InsufficientFundsError
		}
		w.Bankbalance -= BankCard(amount)
	default:
		if coin(amount) > w.balance {
			return InsufficientFundsError
		}
		w.balance += coin(amount)
	}
	return nil
}

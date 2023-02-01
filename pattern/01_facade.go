package main

import (
	"fmt"
)

/*
	Фасад - структурный паттерн, который предоставляет простой интерфейс к сложной системе.
	Плюс его в упрощении интерфейса и сокрытии основной бизнес логики.
	Минус в том,что фасад может стать интерфейсом супер класса => привязка всей системы к нему,
	все будущие функции будут проходить через него.
*/

type WalletFacade struct { // Фасад позволяет клиенту работать с десятками компонентов, используя при этом простой интерфейс.
	account      *Account
	wallet       *Wallet
	securityCode *SecurityCode
	notification *Notification
	ledger       *Ledger
}

func newWalletFacade(accountID string, code int) *WalletFacade { // Создание аккаунта
	fmt.Println("Starting create account")
	walletFacacde := &WalletFacade{
		account:      newAccount(accountID), // Клиенту необходимо ввести реквизиты карты,
		securityCode: newSecurityCode(code), // код безопасности, стоимость оплаты и тип операции.
		wallet:       newWallet(),
		notification: &Notification{},
		ledger:       &Ledger{},
	}
	fmt.Println("Account created")
	return walletFacacde
}

// Фасад управляет дальнейшей коммуникацией между различными компонентами без контакта клиента со сложными внутренними механизмами.

func (w *WalletFacade) addMoneyToWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Starting add money to wallet")
	err := w.account.checkAccount(accountID)
	if err != nil {
		return err
	}
	err = w.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}
	w.wallet.creditBalance(amount)
	w.notification.sendWalletCreditNotification()
	w.ledger.makeEntry(accountID, "credit", amount)
	return nil
}

func (w *WalletFacade) deductMoneyFromWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Starting debit money from wallet")
	err := w.account.checkAccount(accountID)
	if err != nil {
		return err
	}

	err = w.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}
	err = w.wallet.debitBalance(amount)
	if err != nil {
		return err
	}
	w.notification.sendWalletDebitNotification()
	w.ledger.makeEntry(accountID, "credit", amount)
	return nil
}

// Логика подсистемы аккаунта
type Account struct {
	name string
}

func newAccount(accountName string) *Account {
	return &Account{
		name: accountName,
	}
}

func (a *Account) checkAccount(accountName string) error { // Проверка аккаунта
	if a.name != accountName {
		return fmt.Errorf("Account Name is incorrect")
	}
	fmt.Println("Account Verified")
	return nil
}

// Логика подсистемы PIN-кода
type SecurityCode struct {
	code int
}

func newSecurityCode(code int) *SecurityCode {
	return &SecurityCode{
		code: code,
	}
}

func (s *SecurityCode) checkCode(incomingCode int) error { // Проверка PIN-кода
	if s.code != incomingCode {
		return fmt.Errorf("Security Code is incorrect")
	}
	fmt.Println("SecurityCode Verified")
	return nil
}

// Логика подсистемы баланса

type Wallet struct {
	balance int
}

func newWallet() *Wallet {
	return &Wallet{
		balance: 0,
	}
}

func (w *Wallet) creditBalance(amount int) { // Проверка дебета
	w.balance += amount
	fmt.Println("Wallet balance added successfully")
	return
}

func (w *Wallet) debitBalance(amount int) error { // Проверка кредита
	if w.balance < amount {
		return fmt.Errorf("Balance is not sufficient")
	}
	fmt.Println("Wallet balance is Sufficient")
	w.balance = w.balance - amount
	return nil
}

// Логика подсистемы бухгалтерской книги

type Ledger struct {
}

func (s *Ledger) makeEntry(accountID, txnType string, amount int) { // Запись в бухгалтерской книге
	fmt.Printf("Make ledger entry for accountId %s with txnType %s for amount %d\n", accountID, txnType, amount)
	return
}

// Логика подсистемы уведомлений

type Notification struct {
}

func (n *Notification) sendWalletCreditNotification() {
	fmt.Println("Sending wallet credit notification")
}

func (n *Notification) sendWalletDebitNotification() {
	fmt.Println("Sending wallet debit notification")
}

func main() {
	fmt.Println()
	walletFacade := newWalletFacade("abc", 1234)
	fmt.Println()

	err := walletFacade.addMoneyToWallet("abc", 1234, 10)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	fmt.Println()
	err = walletFacade.deductMoneyFromWallet("abc", 1234, 5)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
}

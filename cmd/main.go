package main

import (
	"fmt"
	"github.com/Iftikhor99/wallet/pkg/wallet"
)


func main() {
	svc := &wallet.Service{}
	account , err := svc.RegisterAccount("+992000000001")
	if err != nil {
		fmt.Println(err)
		return
	} 

	err = svc.Deposit(account.ID, 100)
	if err != nil {
		switch err {
		case wallet.ErrAmountMustBePositive:
			fmt.Println("Сумма должна быть положительной")
		case wallet.ErrAccountNotFound:
			fmt.Println("Аккаунт пользователя не найден")		
		}		
		return
	}
	fmt.Println(account.Balance)

	newP, ee2 := svc.Pay(account.ID,20,"food")

	
	fmt.Println(account.Balance)
	fmt.Println(newP)
	fmt.Println(ee2)

	newP2, ee3 := svc.FindAccountById(1)
	fmt.Println(newP2)
	fmt.Println(ee3)

}
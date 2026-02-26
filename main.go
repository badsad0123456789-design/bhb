package main

import (
	"fmt"
	"okak/account"
	"okak/files"
	"okak/okak"
	"strings"

	"github.com/fatih/color"
)

var menu = map[string]func(*account.VaultWithDb){
	"1": createAccount,
	"2": findAccountByUrl,
	"3": findAccountByLogin,
	"4": deleteAccount,
}

var menuVariants = []string{
	"1. Создать аккаунт",
	"2. Найти аккаунт по URL",
	"3. Найти аккаунт по логину",
	"4. Удалить аккаунт",
	"5. Шығу",
	"Танда вариант",
}

func menuCounter() func() {
	i := 0
	return func() { // возвращаем функция она увеличивает 1
		i++
		fmt.Println(i)
	}
}

func main() {
	okak.PrintError(1)
	fmt.Println("Менеджер паролей ")
	counter := menuCounter()
	vault := account.NewVault(files.NewJsonDb("data.json"))
	//vault := account.NewVault(cloud.NewCloudDb("https://a.ru")) // как это связано с cloud

Menu:
	for {
		counter()
		variant := promptData(menuVariants...)
		menuFunc := menu[variant]
		if menuFunc == nil {
			break Menu
		}
		menuFunc(vault)
		// switch variant {
		// case "1":
		// 	createAccount(vault)
		// case "2":
		// 	findAccount(vault)
		// case "3":
		// 	deleteAccount(vault)
		// default:
		// 	break Menu
		// }
	}
}

func findAccountByUrl(vault *account.VaultWithDb) {
	url := promptData("Введите URL для поиска")
	accounts := vault.FindAccountsByUrl(url, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Url, str)
	})
	outPutResult(&accounts)
}

func findAccountByLogin(vault *account.VaultWithDb) {
	login := promptData("Введите login для поиска")
	accounts := vault.FindAccountsByUrl(login, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Login, str)
	})
	outPutResult(&accounts)
}

func outPutResult(accounts *[]account.Account) {
	if len(*accounts) == 0 {
		color.Red("Аккаунтов не найдено")
	}
	for _, acount := range *accounts {
		acount.OutputPass()
	}
}

func deleteAccount(vault *account.VaultWithDb) {
	url := promptData("Введите Url для уадление ")
	isDeleted := vault.DeleteAccount(url)
	if isDeleted {
		color.Green("Удалено")
	} else {
		okak.PrintError(" Не найдено")
	}
}

func createAccount(vault *account.VaultWithDb) {
	login := promptData("Введите логин")
	passowrd := promptData("введите пароль")
	url := promptData("введите url")
	MyAccount, err := account.NewAccount(login, passowrd, url)
	if err != nil {
		okak.PrintError("далбан неверный формат юрл или логин")
		return
	}
	vault.AddAccount(*MyAccount)
}

func promptData(promt ...string) string {
	for i, line := range promt {
		if i == len(promt)-1 {
			fmt.Printf("%v:", line)
		} else {
			fmt.Println(line)
		}
	}
	var res string
	fmt.Scanln(&res)
	return res
}

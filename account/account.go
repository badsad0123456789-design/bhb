package account

import (
	"errors"
	"math/rand/v2"
	"net/url"
	"time"

	"github.com/fatih/color"
)

var letterRunes = []rune("abdefghijklmnoprstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-*!")

type Account struct {
	Login    string    `json:"Login"`
	Passowrd string    `json:"Passowrd"`
	Url      string    `json:"Url"`
	CratedAt time.Time `json:"CreateAt"`
	UpdateAt time.Time `json:"UpdateAt"`
}

func (acc *Account) OutputPass() {
	color.Cyan("acc.Login")
}

func (acc *Account) GeneratePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.Passowrd = string(res)
}

func NewAccount(login, passowrd, urlString string) (*Account, error) {
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}

	if login == "" {
		return nil, errors.New("LOGIN_REQUIRED, логин не может быть пустым ")
	}

	newAcc := &Account{
		CratedAt: time.Now(),
		UpdateAt: time.Now(),
		Login:    login,
		Passowrd: passowrd,
		Url:      urlString,
	}
	if passowrd == "" {
		newAcc.GeneratePassword(12)
	}
	return newAcc, nil
}

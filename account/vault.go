package account

import (
	"encoding/json"
	"okak/encrypter"
	"okak/okak"
	"strings"
	"time"
)

type ByteReader interface {
	Read() ([]byte, error)
}

type ByteWriter interface {
	Write([]byte)
}

type Db interface {
	ByteReader
	ByteWriter
}
type Vault struct {
	Accounts []Account `json:"accounts"`
	UpdateAt time.Time `json:"UpdateAt"`
}

type VaultWithDb struct {
	Vault
	db  Db
	enc encrypter.Encrypter
}

func NewVault(db Db, enc encrypter.Encrypter) *VaultWithDb {
	file, err := db.Read()
	if err != nil {
		return &VaultWithDb{
			Vault: Vault{
				Accounts: []Account{},
				UpdateAt: time.Now(),
			},
			db:  db,
			enc: enc,
		}
	}
	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		okak.PrintError("Не удалось разобрать файл data.json")
		return &VaultWithDb{
			Vault: Vault{
				Accounts: []Account{},
				UpdateAt: time.Now(),
			},
			db:  db,
			enc: enc,
		}
	}
	return &VaultWithDb{
		Vault: vault,
		db:    db,
		enc:   enc,
	}
}

func (vault *VaultWithDb) FindAccountsByUrl(str string, checker func(Account, string) bool) []Account {
	var accounts []Account
	for _, account := range vault.Accounts {
		isMathced := checker(account, str)
		if isMathced {
			accounts = append(accounts, account)
			continue
		}
	}
	return accounts
}

func (valut *VaultWithDb) DeleteAccount(url string) bool {
	var accounts []Account
	isDeleted := false
	for _, account := range valut.Accounts {
		isMathced := strings.Contains(account.Url, url)
		if !isMathced {
			accounts = append(accounts, account)
		}
		isDeleted = true
	}
	valut.Accounts = accounts
	valut.save()
	return isDeleted
}

func (vault *VaultWithDb) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.save()
}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (vault *VaultWithDb) save() {
	vault.UpdateAt = time.Now()
	data, err := vault.Vault.ToBytes()
	if err != nil {
		okak.PrintError(err)
	}
	vault.db.Write(data)
}

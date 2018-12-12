package balance

import (
	"log"
	"time"
	"wallet/db"
)

type Account struct {
	UserID string `json:"id,omitempty"`
	CurrencyId string `json:"currencyid,omitempty"`
	Amount     float64 `json:"amount,omitempty"`
}

type AddressInfo struct {
	ID 	int `json:"id,omitempty"`
	CurrencyId string `json:"currencyid,omitempty"`
	Address  string   `json:"address,omitempty"`
	PublicKey string   `json:"pubkey,omitempty"`
	AddrType  int  `json:"addrtype,omitempty"`
	Amount     float64 `json:"amount,omitempty"`
	LockState  float64  `json:"lockstate,omitempty"`
	UserID string `json:"userid,omitempty"`
	CreateTime time.Time `json:"createtime,omitempty"`
	UpdateTime time.Time `json:"unpdatetime,omitempty"`
}

func GetBalance(userId string , currencyId string, userInfo chan Account) {
	db := db.InitDB(db.DefaultBitcoinDB)
	defer db.Close()
	rows := db.QueryRow("SELECT ID,currencyid,amount FROM address_account where userid =?", userId)
	addressInfo := &AddressInfo{}
	if err := rows.Scan(&addressInfo.ID, &addressInfo.CurrencyId, &addressInfo.Amount); err != nil {
		log.Panic("Read fail")
	}
	account := Account{
		UserID: userId,
		CurrencyId: currencyId,
		Amount: addressInfo.Amount,
	}
	userInfo <- account
}

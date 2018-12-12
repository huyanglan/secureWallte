package db

import (
	"fmt"
	"log"
	"testing"
	"time"
)

type AddressInfo struct {
	ID 	int `json:"id,omitempty"`
	CurrencyId string `json:"currencyid,omitempty"`
	Address  string   `json:"address,omitempty"`
	PublicKey string   `json:"pubkey,omitempty"`
	AddrType  int  `json:"addrtype,omitempty"`
	Amount     float64 `json:"amount,omitempty"`
	LockState  float64  `json:"lockstate,omitempty"`
	UserID string `json:"id,omitempty"`
	CreateTime time.Time `json:"createtime,omitempty"`
	UpdateTime time.Time `json:"unpdatetime,omitempty"`
}

func TestQuery(t *testing.T) {
	dbName   := "db_btc_wallet"
	db :=InitDB(dbName)
	defer db.Close()
	rows, err := db.Query("SELECT ID,CurrencyId FROM address_account")
	if err != nil{
		log.Panic("select error")
	}
	info := &AddressInfo{}
	for rows.Next() {
		err = rows.Scan(&info.ID,&info.CurrencyId)
		fmt.Println("row", info.ID, 1111, info.CurrencyId)
	}

}

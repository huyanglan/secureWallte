package db

import (
	"fmt"
	"github.com/cointux/server/pacautils/src/github.com/jinzhu/gorm"
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

func TestGrom(t *testing.T) {
	db, err := gorm.Open("mysql", "root:123456@tcp(www.pacasys.com:3306)/gotest?charset=utf8&parseTime=True")
	defer db.Close()
	if err != nil {
		log.Panic("Connect fail !")
	}
	log.Print("Connect Success !")
}

func Test(t *testing.T) {
	fmt.Print(time.Now().Unix())
}

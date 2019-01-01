package main


import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strconv"
)

func main() {

	//url := "https://chain.api.btc.com/v3/address/1EH1aZ7czpoc2QidinqiY11BgDHrnVMKwr"

	//url := "https://chain.api.btc.com/v3/address/1EH1aZ7czpoc2QidinqiY11BgDHrnVMKwr,15bPG8mNevnmcDBU5GPTfJxGW32VUm7fn1"

	url :="https://chain.api.btc.com/v3/address/13QLr1MB7pLxn9igEwxhiNNYGWs7ishP49"
	req, err1 := http.NewRequest("GET", url, nil)
    log.Print("err1 ::      " , err1,  req  )
	req.Header.Add("accept", "text/plain,text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	//req.Header.Add("accept-encoding", "gzip, deflate, br")
	//req.Header.Add("accept-language", "zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Add("cache-control", "max-age=0")
	//req.Header.Add("content-type", "application/json")
	req.Header.Add("content-type", "application/json; charset=UTF-8")
	res,err2 := http.DefaultClient.Do(req)
	log.Print("err2 ::      " ,err2,   req  )
	//defer res.Body.Close()
	body,err3 := ioutil.ReadAll(res.Body)

	log.Print("err3 ::      " ,   err3  )

	type data struct {
		Address       string  `json:"address"`
		Received      int `json:"received"`
		Sent       int `json:"sent"`
		Balance      int `json:"balance"`
		Txcount       int  `json:"tx_count"`
		Unconfirmedtxcount      int `json:"unconfirmed_tx_count"`
		Unconfirmedreceived       int  `json:"unconfirmed_received"`
		Unconfirmedsent      int `json:"unconfirmed_sent"`
		Unspenttxcount       int  `json:"unspent_tx_count"`
		Firsttx      string `json:"first_tx"`
		Lasttx      string `json:"last_tx"`


	}
	type Province struct {

		Errno       int  `json:"err_no"`
		Data        data `json:"data"`

	}
	
	provinces := &Province{}
	log.Println("---body---", []byte(body))

	err := json.Unmarshal([]byte(body), &provinces)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(provinces)
	fmt.Println(provinces.Errno)
	s := strconv.Itoa(provinces.Data.Balance)
	fmt.Println(provinces.Data.Balance)
	addressBalance, _ := strconv.ParseFloat(s, 64)
	fmt.Println(reflect.TypeOf(addressBalance/100000000))


	var a float64
	a = 0.0
	if addressBalance/100000000 != a {
		fmt.Println("地址余额不为0.00")
	}else {
		fmt.Println("地址余额不为>>>>>>")
	}

	fmt.Println( strconv.FormatFloat((addressBalance/100000000), 'f', -1, 64))
	v := 3.1415926535
	s1 := strconv.FormatFloat(v, 'f',-1, 64)
	fmt.Println(s1)
}

